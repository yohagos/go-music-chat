package util

import (
	"fmt"
	"os"
)

var (
	musicDir string
)

func CreateMusicFolder() {
	path, _ := os.UserHomeDir()
	
	musicDir = fmt.Sprintf("%s/music", path)
	if _, err := os.Stat(musicDir); os.IsNotExist(err) {
		os.Mkdir(musicDir, 0755)
	}
}

func CreateUserFolder(user string) string {
	path := fmt.Sprintf("%s/%s", musicDir, user)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		return path
	}
	return path
}

