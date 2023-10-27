package Sky

type SkyCoordinates interface {
	ToString() string
	ToHex() ([]byte, error)
	ToHexPrecise() ([]byte, error)
	//ConvToArcSeconds() int
}
