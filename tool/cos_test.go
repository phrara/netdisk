package tool

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func TestCOS1(t *testing.T) {
  
	u, _ := url.Parse("https://netdisk-1304379399.cos.ap-beijing.myqcloud.com")
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            // 通过环境变量获取密钥
            SecretID: "AKID3b4KXfhYCdtAJi4LWbeXEj196yDKGy0p",  
            SecretKey: "lKgpmW3kLkwvkMk05zgdS4TfcZSARUce",  
        },
    })


    key := "test/test1.jpg"


    _, _, err := client.Object.Upload(
        context.Background(), key, "C:\\Users\\lenovo\\Pictures\\Saved Pictures\\4.jpg", nil,
    )
    if err != nil {
        panic(err)
    }

}

func TestCOS2(t *testing.T) {
  
	u, _ := url.Parse("https://netdisk-1304379399.cos.ap-beijing.myqcloud.com")
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            // 通过环境变量获取密钥
            SecretID: "AKID3b4KXfhYCdtAJi4LWbeXEj196yDKGy0p",  
            SecretKey: "lKgpmW3kLkwvkMk05zgdS4TfcZSARUce",  
        },
    })


    key := "test/test2.jpg"

	b2, _ := os.ReadFile("C:\\Users\\lenovo\\Pictures\\Saved Pictures\\vsc.jpg")

    _, err := client.Object.Put(
        context.Background(), key, bytes.NewReader(b2), nil,
    )
    if err != nil {
        panic(err)
    }

}