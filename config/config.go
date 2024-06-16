package config

type Config struct {
	Server         Server
	TemplatesPaths []string
}

type Server struct {
	Host string
	Port string
}
