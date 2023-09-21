package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	s := CreateCaptcha(6)
	t.Log(s)
}

func CreateCaptcha(num int) string {
	str := "1"
	for i := 0; i < num; i++ {
		str += strconv.Itoa(0)
	}
	str10 := str
	int10, err := strconv.ParseInt(str10, 10, 32)
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		j := int32(int10)
		return fmt.Sprintf("%0" + strconv.Itoa(num) + "v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(j))
	}
}