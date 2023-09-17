package tool

import (
	"bytes"
	"context"

)

func COSUpload(file []byte, cosPath string) (path string, err error) {
	_, err = cosClient.Object.Put(
		context.Background(), cosPath, bytes.NewReader(file), nil,
	)
	if err != nil {
		return "", err
	}
	return Conf.COS.COSBucketAddr + "/" + cosPath, nil
}
