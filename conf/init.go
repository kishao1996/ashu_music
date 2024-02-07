package conf

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	config = &Config{}
)

type Config struct {
	RootDir    string
	FfmpegPath string
}

func GetConfig() *Config {
	return config
}

func Init() {
	_, name, _, _ := runtime.Caller(0)
	paths := strings.Split(name, "/")
	rootDir := strings.Join(paths[:len(paths)-2], "/") + "/"
	config.RootDir = rootDir
	fmt.Println("ROOTDIR: ", config.RootDir)
}
