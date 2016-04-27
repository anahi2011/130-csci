package main

import(
	"io"
	"fmt"
	"bytes"
	"io/ioutil"
	"strings"
	"net/http"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
)

var bucket = "bucket1.appspot.com"
type demo struct {
	bucket *storage.BucketHandle
	client *storage.Client
	w   http.ResponseWriter
	ctx context.Context
	cleanUp []string
	placeholder string
	failed bool
}

func (d *demo) createFile(fileName string) {
	fmt.Fprintf(d.w, "Creating file /%v/%v\n", bucket, fileName)
	wc := d.bucket.Object(fileName).NewWriter(d.ctx)
	wc.ContentType = "text/plain"
	wc.Metadata = map[string]string{
		"x-goog-meta-foo": "foo",
		"x-goog-meta-bar": "bar",
	}
	d.cleanUp = append(d.cleanUp, fileName)
	if _, err := wc.Write([]byte("abcde\n")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", bucket, fileName, err)
		return
	}
	if _, err := wc.Write([]byte(strings.Repeat("f", 1024*4) + "\n")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", bucket, fileName, err)
		return
	}
	if err := wc.Close(); err != nil {
		log.Errorf(d.ctx, "createFile: unable to close bucket %q, file %q: %v", bucket, fileName, err)
		return
	}
}


func (d *demo) readFile(fileName string) {
	io.WriteString(d.w, "\nAbbreviated file content (first line and last 1K):\n")
	rc, _ := d.bucket.Object(fileName).NewReader(d.ctx)

	defer rc.Close()
	slurp, _ := ioutil.ReadAll(rc)
	fmt.Fprintf(d.w, "%s\n", bytes.SplitN(slurp, []byte("\n"), 2)[0])
	if len(slurp) > 1024 { fmt.Fprintf(d.w, "...%s\n", slurp[len(slurp)-1024:]) } else { fmt.Fprintf(d.w, "%s\n", slurp) }
}
func (d *demo) listBucket() {
	io.WriteString(d.w, "\nListbucket result:\n")
	query := &storage.Query{Delimiter: "/"}
	objs, _ := d.bucket.List(d.ctx, query)
	for _, obj := range objs.Results {
		fmt.Fprintf(d.w,"\n" + obj.Name)
	}
	for _, i := range objs.Prefixes{
		fmt.Fprintf(d.w,"\n" + i)
	}
}
