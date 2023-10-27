<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';

const props = defineProps<{
    stream: MediaStream;
    dc: RTCDataChannel;
}>();

const video = ref<HTMLVideoElement | null>(null);

onMounted(() => {
    watch(
        () => props.stream,
        (stream) => {
            if (stream) setStream(stream);
        },
        {
            immediate: true,
        }
    );
});

function setStream(stream: MediaStream) {
    if (!video.value) {
        throw new Error('Video element not found');
    }

    console.log('Received stream: ', stream);
    video.value.srcObject = stream;
}

function sendSlew(direction: string) {
    props.dc.send(JSON.stringify({ type: 'slew', direction }));
}

function sendPhoto() {
    props.dc.send(JSON.stringify({ type: 'take_photo' }));
}
</script>

<template>
    <div class="control-panel">
        <video
            ref="video"
            class="control-panel__video"
            playsinline
            autoplay
            muted
        ></video>
        <div class="control-panel__controls">
            <button
                aria-label="Up"
                class="control-panel__slew-btn control-panel__slew-btn--up"
                @click="sendSlew('up')"
            ></button>
            <button
                aria-label="Down"
                class="control-panel__slew-btn control-panel__slew-btn--down"
                @click="sendSlew('down')"
            ></button>
            <button
                aria-label="Left"
                class="control-panel__slew-btn control-panel__slew-btn--left"
                @click="sendSlew('left')"
            ></button>
            <button
                aria-label="Right"
                class="control-panel__slew-btn control-panel__slew-btn--right"
                @click="sendSlew('right')"
            ></button>
        </div>
        <button class="control-panel__photo-btn" @click="sendPhoto"></button>
    </div>
</template>

<style scoped lang="scss">
.control-panel {
    position: relative;
    display: flex;
    justify-content: center;
    margin: 0 auto;
    width: 800px;
    aspect-ratio: 4 / 3;

    &__video {
        width: 100%;
        height: 100%;
    }

    &__controls {
        position: absolute;
        bottom: 10px;
        right: 10px;
        width: 100px;
        height: 100px;
        background-color: rgba(0, 0, 0, 0.5);
    }

    &__slew-btn {
        width: 25px;
        height: 25px;
        background: url('/img/arrow.svg') no-repeat center center;
        background-size: contain;

        &:active {
            opacity: 0.8;
        }

        &--up {
            position: absolute;
            top: 0;
            left: 50%;
            transform: rotate(-90deg) translateY(-50%);
        }

        &--down {
            position: absolute;
            bottom: 0;
            left: 50%;
            transform: rotate(90deg) translateY(50%);
        }

        &--left {
            position: absolute;
            top: 50%;
            left: 0;
            transform: rotate(180deg) translateY(50%);
        }

        &--right {
            position: absolute;
            top: 50%;
            right: 0;
            transform: translateY(-50%);
        }
    }

    &__photo-btn {
        position: absolute;
        bottom: 15px;
        left: 15px;
        width: 25px;
        height: 25px;
        background: url('/img/do_photo.svg') no-repeat center center;
        background-size: contain;
    }
}
</style>
