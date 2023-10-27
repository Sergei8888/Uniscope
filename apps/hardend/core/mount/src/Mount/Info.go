package mount

type Info interface {
	String() string
	GetInfoByID(infoID int) string
}
