package configobj

type Config interface {
	GetConfigPath() string
	GetConfigData(key ...string) (map[string]interface{}, error)
}
