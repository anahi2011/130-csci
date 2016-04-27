package main

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
)

func handler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	ctx := appengine.NewContext(res)
	if bucket == "" {
		var err error
		if bucket, err = file.DefaultBucketName(ctx); err != nil {
			log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
			return
		}
	}
	//[END get_default_bucket]
	//Make Client
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
		return
	}
	defer client.Close()


	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(res, "Demo GCS Application running from Version: %v\n", appengine.VersionID(ctx))
	fmt.Fprintf(res, "Using bucket name: %v\n\n", bucket)

	d := &demo{
		w:      res,
		ctx:    ctx,
		client: client,
		bucket: client.Bucket(bucket),
	}

	n := "demo-testfile-go"

	d.createFile(n)
	d.listBucket()
}

func retrieve(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
		return
	}
	d := &demo{
		w:      res,
		ctx:    ctx,
		client: client,
		bucket: client.Bucket(bucket),
	}
	n := "demo-testfile-go"

	d.readFile(n)
	d.listBucket()
}

func index(res http.ResponseWriter, req *http.Request) {}

func init(){
	http.HandleFunc("/", index)
	http.HandleFunc("/handler", handler)
	http.HandleFunc("/retrieve", retrieve)
}