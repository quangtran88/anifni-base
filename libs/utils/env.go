package baseUtils

import "os"

type EnvManager struct {
}

var envManager *EnvManager

func GetEnvManager() *EnvManager {
	if envManager == nil {
		envManager = &EnvManager{}
	}
	return envManager
}

func (m EnvManager) GetEnv(key string) string {
	return os.Getenv(key)
}

func (m EnvManager) GetEnvDefault(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return defaultValue
	}
}
