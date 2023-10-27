package constants

type ASI_FLIP_STATUS uint8

const (
	ASI_FLIP_NONE  = iota // no flip
	ASI_FLIP_HORIZ        // horizontal image flip
	ASI_FLIP_VERT         // vertical image flip
	ASI_FLIP_BOTH         // horizontal + vertical image flip
)
