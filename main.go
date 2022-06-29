package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var bucket string
var prefix = ""
var progress = true

func main() {
	flag.StringVar(&bucket, "bucket", "", "Bucket to read data from")
	flag.StringVar(&prefix, "prefix", "", "Path to use as prefix")
	flag.BoolVar(&progress, "progress", true, "Show progress in stdout")

	flag.Parse()

	start := time.Now()

	progressInfof("Setting up session\n")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := s3.New(sess)

	var totalDataBytes int64 = 0
	var totalFetchBytes int64 = 0
	var totalFetches int64 = 0
	buff := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloader(sess)

	progressInfof("Fetching object list\n")
	svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}, func(resp *s3.ListObjectsOutput, last bool) bool {
		for _, key := range resp.Contents {
			progressInfof("Streaming %s\n", *key.Key)

			numFetchBytes, err := downloader.Download(buff, &s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(*key.Key),
			})
			if err != nil {
				panic(err)
			}

			totalFetchBytes += numFetchBytes
			totalFetches += 1

			var reader io.Reader
			if strings.HasSuffix(*key.Key, ".gz") {
				reader, _ = gzip.NewReader(bytes.NewReader(buff.Bytes()))
			} else {
				reader = bytes.NewReader(buff.Bytes())
			}
			numBytes, _ := io.Copy(os.Stdout, reader)
			totalDataBytes += numBytes
		}
		return !last
	})

	progressInfof(
		"\nFinished streaming\n  Data: %s \n  Fetched %s in %d files \n  Duration: %s\n",
		ByteCount(totalDataBytes),
		ByteCount(totalFetchBytes),
		totalFetches,
		time.Since(start),
	)
}

func progressInfof(format string, args ...interface{}) {
	if progress {
		fmt.Fprintf(os.Stderr, format, args...)
	}
}

func ByteCount(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
