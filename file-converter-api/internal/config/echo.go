package config

type ServerConf struct {
	Port string
}

func NewServerConf() *ServerConf {
	return &ServerConf{
		Port: ":3000",
	}
}
