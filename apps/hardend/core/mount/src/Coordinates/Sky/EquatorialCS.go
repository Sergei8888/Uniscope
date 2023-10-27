package Sky

type EquatorialCS struct {
	ra  *RightAscension
	dec *Declination
}

func NewEquatorialCS(ra *RightAscension, dec *Declination) *EquatorialCS {
	return &EquatorialCS{ra: ra, dec: dec}
}

func (e EquatorialCS) Ra() *RightAscension {
	return e.ra
}

func (e EquatorialCS) Dec() *Declination {
	return e.dec
}
