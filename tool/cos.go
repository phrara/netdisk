package tool

import (
	"bytes"
	"context"
	"io"
	"math"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func COSUpload(file []byte, cosPath string) error {
	_, err := cosClient.Object.Put(
		context.Background(), cosPath, bytes.NewReader(file), nil,
	)
	if err != nil {
		return err
	}
	return nil
}

// 分片上传
func COSMultipartUpload(file []byte, cosPath string) error {
	// 分块初始化
	v, _, err := cosClient.Object.InitiateMultipartUpload(
		context.Background(), cosPath, nil,
	)
    if err != nil {
        return err
    }

	// 上传分片
	chunkSize := Conf.COS.ChunkSize * 1024 * 1024
	chunkNum := math.Ceil(float64(len(file)) / float64(chunkSize))
	opt := &cos.CompleteMultipartUploadOptions{}

	for i := 1; i <= int(chunkNum); i++ {
		if i != int(chunkNum) {
			if resp, err := cosClient.Object.UploadPart(
				context.Background(), cosPath, v.UploadID, i, bytes.NewReader(file[(i-1)*chunkSize:i*chunkSize]), nil,
			); err != nil {
				return err
			} else {
				opt.Parts = append(opt.Parts, cos.Object{
					PartNumber: i,
					ETag: resp.Header.Get("ETag"), 
				})
			}

		} else {
			if resp, err := cosClient.Object.UploadPart(
				context.Background(), cosPath, v.UploadID, i, bytes.NewReader(file[(i-1)*chunkSize:]), nil,
			); err != nil {
				return err
			} else {
				opt.Parts = append(opt.Parts, cos.Object{
					PartNumber: i,
					ETag: resp.Header.Get("ETag"),
				})
			}
		}
	}
    
	// 完成分片上传
	_, _, err = cosClient.Object.CompleteMultipartUpload(
		context.Background(), cosPath, v.UploadID, opt,
	)
	if err != nil {
		return err
	}

	return nil

}


func COSDownload(cosPath string) (file []byte, err error)  {
	resp, err := cosClient.Object.Get(context.Background(), cosPath, nil)
	if err != nil {
		return nil, err
	}
	bs, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	return bs, nil
}