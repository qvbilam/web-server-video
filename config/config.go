package config

type ServerConfig struct {
	Name              string            `mapstructure:"name" json:"name"`
	Host              string            `mapstructure:"host" json:"host"`
	Port              string            `mapstructure:"port" json:"port"`
	Tags              []string          `mapstructure:"tags" json:"tags"`
	UserServerConfig  UserServerConfig  `mapstructure:"user-server" json:"user-server"`
	VideoServerClient VideoServerClient `mapstructure:"video-server" json:"video-server"`
}

type UserServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type VideoServerClient struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
