package download

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	cli                  = http.Client{}
	dir                  = "/Users/bytedance/workspace/ashu_music/music"
	UserAgentHeaderKey   = "User-Agent"
	UserAgentHeaderValue = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3970.5 Safari/537.36"
)

func Init() {
	_, err := os.Stat(dir)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

type bvInfo struct {
	Id          string
	Name        string
	DownloadUrl string
}

func getBvInfo(bvId string) (*bvInfo, error) {
	// request
	request, err := http.NewRequest("GET",
		"https://www.bilibili.com/video/"+bvId,
		nil,
	)
	if err != nil {
		return nil, err
	}
	// add headers
	request.Header.Set(UserAgentHeaderKey, UserAgentHeaderValue)
	request.Header.Set("Referer", "https://www.bilibili.com/")
	// do req
	response, err := cli.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// parse response
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	content := doc.Find("script").Eq(2).Text()
	raw := struct {
		Data struct {
			Dash struct {
				Audio []struct {
					BaseUrl string `json:"baseUrl"`
				} `json:"audio"`
			} `json:"dash"`
		} `json:"data"`
	}{}
	err = json.Unmarshal([]byte(content[20:]), &raw)
	if err != nil {
		return nil, err
	}
	bvInfo := &bvInfo{
		Id:          bvId,
		DownloadUrl: raw.Data.Dash.Audio[0].BaseUrl,
		Name:        strings.Split(doc.Find("div.video-container-v1 .left-container").First().Text(), "\n")[0],
	}
	return bvInfo, nil
}

func getAudioPart(bvInfo *bvInfo, begin int, end int) ([]byte, bool, error) {
	request, err := http.NewRequest("GET", bvInfo.DownloadUrl, nil)
	// add headers
	request.Header.Set(UserAgentHeaderKey, UserAgentHeaderValue)
	request.Header.Set("Referer", "https://www.bilibili.com/video/"+bvInfo.Id)
	request.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", begin, end))
	response, err := cli.Do(request)
	if err != nil {
		return nil, false, err
	}
	// is end
	isEnd := false
	if response.StatusCode == 416 {
		request, err := http.NewRequest("GET", bvInfo.DownloadUrl, nil)
		request.Header.Set("Range", fmt.Sprintf("bytes=%d-", end+1))
		response, err = cli.Do(request)
		if err != nil {
			return nil, false, err
		}
		isEnd = true
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, false, err
	}
	return bytes, isEnd, nil
}

func Download(bvId string) error {
	bvInfo, err := getBvInfo(bvId)
	if err != nil {
		return err
	}
	filePath := filepath.Join(dir, bvInfo.Name+".m4a")
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	partSize := 1024 * 512
	begin, end := 0, partSize-1
	for {
		bytes, isEnd, err := getAudioPart(bvInfo, begin, end)
		if err != nil {
			return err
		}
		write.Write(bytes)
		if isEnd {
			break
		}
		begin += partSize
		end += partSize
	}
	write.Flush()
	return nil
}
