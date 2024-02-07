package download

import (
	"ashu_music/conf"
	"ashu_music/src"
	"fmt"
	"testing"

	imageio "github.com/openatx/go-imageio"
)

var (
	tiangouBvs = []string{
		"BV1hY411J7j5",
		"BV12B4y1Q7om",
		"BV1cg411v7Rr",
		"BV1cP411K7bn",
		"BV1Bv4y1m7wv",
		"BV1RM4y1q7dQ",
	}
	chengzhongcunBvs = []string{
		"BV1Q54y1u7yy",
		"BV1Xo4y1W73N",
		"BV1vs4y1U7db",
		"BV1yM4y1C7Zo",
		"BV1nh411j7Lv",
		"BV1pL411a7sp",
	}
)

func TestDownload(t *testing.T) {
	conf.Init()
	src.Init()
	Init()
	// for _, bv := range chengzhongcunBvs {
	// 	err := Download(bv)
	// 	fmt.Println(err)
	// }
	err := Download("BV13b4y1P7dy")
	fmt.Println(err)
}

func Test_getBvInfo(t *testing.T) {
	Init()
	info, err := getBvInfo("BV1wY411N756")
	fmt.Println(info)
	fmt.Println(err)
}

func Test_read(t *testing.T) {
	exe, err := imageio.GetFFmpegExe()
	fmt.Println(exe, err)
	// path := fmt.Sprintf("%s:%s/%s", os.Getenv("PATH"), "/Users/bytedance/workspace/ashu_music/src/download", exe)
	// os.Setenv("PATH", path)
	// fmt.Println(path)
	// fmt.Println(imageio.CheckIfFFmpegInPATH())
}
