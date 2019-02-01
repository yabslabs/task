package configuration

// ServiceConfig is the basic config of each service
// generall implementation in services is reading it from an environment based config.toml file
// deprecated
type ServiceConfig struct {
	UserPassword
	GRPCPort    string
	GatewayPort string
}

type UserPassword struct {
	Username string
	Password string
}

func (u UserPassword) GetUsername() string {
	return u.Username
}

func (u UserPassword) GetPassword() string {
	return u.Password
}
