package Sky

type HorizontalCS struct {
	azm *Azimuth
	alt *Altitude
}

func NewHorizontalCS(azm *Azimuth, alt *Altitude) *HorizontalCS {
	return &HorizontalCS{azm: azm, alt: alt}
}

func (h HorizontalCS) Azm() *Azimuth {
	return h.azm
}

func (h HorizontalCS) Alt() *Altitude {
	return h.alt
}
