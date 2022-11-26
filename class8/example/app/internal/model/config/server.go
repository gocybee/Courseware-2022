package config

import (
	"time"
)

type Server struct {
	Mode         string `mapstructure:"mode" yaml:"mode"`
	Port         string `mapstructure:"port" yaml:"port"`
	ReadTimeout  string `mapstructure:"readTimeout" yaml:"readTimeout"`
	WriteTimeout string `mapstructure:"writeTimeout" yaml:"writeTimeout"`
}

func (s *Server) Addr() string {
	return ":" + s.Port
}

func (s *Server) GetReadTimeout() time.Duration {
	t, _ := time.ParseDuration(s.ReadTimeout)
	return t
}

func (s *Server) GetWriteTimeout() time.Duration {
	t, _ := time.ParseDuration(s.WriteTimeout)
	return t
}
