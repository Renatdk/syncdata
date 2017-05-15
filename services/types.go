package services

type Connection struct {
	Server   string `json: server`
	Database string `json: database`
	Port     int    `json: port`
	User     string `json: user`
	Password string `json: password`
	Debag    bool   `json: debag`
}
