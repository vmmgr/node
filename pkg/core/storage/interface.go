package storage

type Base struct {
	Controller string
	UUID       string
	IsDebug    bool
}

type SFTPAuth struct {
	IP   string
	User string
	Pass string
}
