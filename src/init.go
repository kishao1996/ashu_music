package src

import (
	"ashu_music/conf"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	MusicDir = ""
	Musics   = []*Music{}
)

type Music struct {
	Name string
	Path string
}

func Init() {
	rootDir := conf.GetConfig().RootDir
	os.Setenv("FYNE_FONT", rootDir+"/src/theme/font/微软雅黑.ttf")
	MusicDir = filepath.Join(rootDir, "music")
	files, _ := ioutil.ReadDir(MusicDir)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".mp3") {
			music := &Music{
				Name: file.Name()[:len(file.Name())-4],
				Path: filepath.Join(MusicDir, file.Name()),
			}
			Musics = append(Musics, music)
		}
	}
}
