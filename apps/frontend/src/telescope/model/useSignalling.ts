import { io } from 'socket.io-client';

import { ElNotification } from 'element-plus';
import { useRouter } from 'vue-router';

import { ref } from 'vue';

import { useTokenStore } from '@/shared/model/token.store';

export function useSignalling(telescopeId: number) {
    const router = useRouter();
    const isReady = ref(false);

    const socket = io(import.meta.env.VITE_TELESCOPE_CONNECTOR_URL, {
        path: import.meta.env.VITE_TELESCOPE_CONNECTOR_PATH,
        auth: {
            access_token: useTokenStore().accessToken,
            telescopeId: telescopeId,
        },
    });

    socket.on('exception', async (err) => {
        for (const errMessage of err) {
            ElNotification.error({
                title: 'Сокет сервер порвал соединение',
                message: errMessage,
                duration: 3000,
            });
        }
        await router.push({ name: 'telescope-catalog' });
    });

    let makingOffer = false;
    let ignoreOffer = false;
    const polite = false;

    const pc = new RTCPeerConnection({
        iceServers: [
            {
                urls: 'stun:stun.l.google.com:19302',
            },
        ],
    });

    const dc = pc.createDataChannel('commands');

    pc.addEventListener('connectionstatechange', (event) => {
        if (pc.connectionState === 'connected') {
            console.log('Connected');
            isReady.value = true;
        }

        if (pc.iceConnectionState === 'failed') {
            console.log('Ice failed, restarting ice');
            pc.restartIce();
        }
    });

    pc.addEventListener('negotiationneeded', async () => {
        console.log('Negotiation needed, making an offer');

        try {
            makingOffer = true;
            await pc.setLocalDescription();
            socket.emit('message', {
                description: pc.localDescription,
            });
        } catch (err) {
            console.error(err);
        } finally {
            makingOffer = false;
        }

        console.log('Offer sent: ', pc.localDescription);
    });

    pc.addEventListener('icecandidate', (event) => {
        console.log('Sending ice candidate');
        if (event.candidate) {
            socket.emit('message', {
                ice_candidate: event.candidate,
            });
        }
    });

    // Signalling server
    socket.on('exception', (err) => {
        console.error(err);
    });

    socket.on('message', async (message) => {
        if (message.description) {
            await handleDescription(message.description);
        } else if (message.ice_candidate) {
            await handleIceCandidate(message.ice_candidate);
        }
    });

    socket.on('disconnect', async () => {
        ElNotification.error({
            title: 'Сокет сервер порвал соединение',
            duration: 3000,
        });
        await router.push({ name: 'telescope-catalog' });
    });

    async function handleDescription(description: RTCSessionDescription) {
        console.log('Received description: ', description);
        const offerCollision =
            description.type === 'offer' &&
            (makingOffer || pc.signalingState !== 'stable');

        ignoreOffer = !polite && offerCollision;

        if (ignoreOffer) {
            console.log('Ignoring offer');
            return;
        }

        await pc.setRemoteDescription(description);
        if (description.type === 'offer') {
            await pc.setLocalDescription();
            console.log('Answer sent: ', pc.localDescription);
            socket.emit('message', {
                description: pc.localDescription,
            });
        }
    }

    async function handleIceCandidate(candidate: RTCIceCandidate) {
        console.log('Received ice candidate');
        try {
            await pc.addIceCandidate(candidate);
        } catch (err) {
            if (!ignoreOffer) {
                throw err;
            }
        }
    }

    return { isReady, dc, pc };
}
