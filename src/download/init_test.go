package download

import (
	"fmt"
	"testing"
)

func TestDownload(t *testing.T) {
	Init()
	err := Download("BV1fi4y1j7qw")
	fmt.Println(err)
}

func Test_getBvInfo(t *testing.T) {
	Init()
	info, err := getBvInfo("BV1GK411M7Zb")
	fmt.Println(info)
	fmt.Println(err)
}
