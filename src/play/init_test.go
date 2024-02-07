package play

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func TestPlay(t *testing.T) {
	err := Play("/Users/bytedance/workspace/ashu_music/music/abc.mp3")
	fmt.Println(err)
}

func TestPlay1(t *testing.T) {
	p := os.Getenv("PATH") + ":" + "/Users/bytedance/workspace/ashu_music/src/download/"
	fmt.Println(p)
	os.Setenv("PATH", p)
	path := "/Users/bytedance/workspace/ashu_music/music/【4K修复】邓丽欣 - 电灯胆 MV.m4a"
	_resultVideoPath := filepath.Join("/Users/bytedance/workspace/ashu_music/music", fmt.Sprintf("%s.%s", "abc", "mp3"))
	err := ffmpeg.Input(path).
		Output(_resultVideoPath, ffmpeg.KwArgs{"acodec": "libmp3lame"}).
		OverWriteOutput().ErrorToStdOut().Run()
	fmt.Println(err)
}

// func Test_checkFfmpeg(t *testing.T) {
// 	conf.Init()
// 	src.Init()
// 	err := checkFfmpeg()
// 	fmt.Println(err)
// }
