package files

import (
	"fmt"
	"path/filepath"
)

func AvatarPath(uid uint32) (string, error) {
	return filepath.Abs(fmt.Sprintf("./files/avatar/%d.avatar", uid))
}

func DefaultAvatarPath() (string, error) {
	return filepath.Abs("./statics/avatar.png")
}

func ConfigFilePath() (string, error) {
	return filepath.Abs("./config.json")
}
