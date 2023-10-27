export class RtcSessionDescriptionDto {
    sdp?: string;
}

export class RtcIceCandidateDto implements RTCIceCandidateInit {
    candidate?: string;

    sdpMid?: string | null;

    sdpMLineIndex?: number | null;

    usernameFragment?: string | null;
}
