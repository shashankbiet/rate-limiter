package config

// Viper uses the mapstructure package under the hood for unmarshaling values, so we use the mapstructure tags to specify the name of each config field.
type Configuration struct {
	Environment string           `mapstructure:"ENVIRONMENT"`
	AppName     string           `mapstructure:"APP_NAME"`
	HttpServer  HttpServerConfig `mapstructure:"HTTP_SERVER"`
	GrpcServer  GrpcServerConfig `mapstructure:"GRPC_SERVER"`
}

type HttpServerConfig struct {
	Port               int `mapstructure:"PORT"`
	ReadTimeoutMs      int `mapstructure:"READ_TIMEOUT_MS"`
	WriteTimeoutMs     int `mapstructure:"WRITE_TIMEOUT_MS"`
	IdleTimeoutMs      int `mapstructure:"IDLE_TIMEOUT_MS"`
	KeepAliveTimeoutMs int `mapstructure:"KEEP_ALIVE_TIMEOUT_MS"`
}

type GrpcServerConfig struct {
	Port int `mapstructure:"PORT"`
}
