package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"allan/config"
)

var EP string
var AK string
var SK string

func init() {
	AK = "LTAI4FoGLg8Q8kQJEwTF58L6"
	SK = "dx6g63mDak07ioZQo9ZXhDzsNEO34B"
	EP = config.GetOssAddr()
}

func UploadToOss(filename string, path string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service error: %s", err)
		return false
	}

	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("Getting bucket error: %s", err)
		return false
	}

	err = bucket.UploadFile(filename, path, 500*1024, oss.Routines(3))
	if err != nil {
		log.Printf("Uploading object error: %s", err)
		return false
	}

	return true
}
