package initialize

import (
	"danmu/services/config"
	"os"
)

func Config() {
	config.Global = &config.GlobalConfigType{}
	config.ReadFile(config.Global)
	config.Secret = os.Getenv("Secret")
}
