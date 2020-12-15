package s3format

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

const (
	s3path = "^(s3-|s3\\.)?(s3.*)\\.amazonaws\\.com"
	s3url  = "^s3://*"
	s3vh   = "(s3.*)\\.amazonaws\\.com$"
)

// Format is a
func Format(buckets chan string, options *Options) {

	var uniqueBuckets sync.Map
	var uniqueResults sync.Map

	for bucket := range buckets {
		if _, exists := uniqueBuckets.Load(bucket); exists {
			continue
		}

		target := strings.Replace(bucket, "http://", "", 1)
		target = strings.Replace(target, "https://", "", 1)
		target = strings.Replace(target, "s3://", "s3:////", 1)
		target = strings.Replace(target, "//", "", 1)

		var s3name string

		if path, _ := regexp.MatchString(s3path, target); path {
			target = strings.Replace(target, "s3.amazonaws.com/", "", 1)
			target = strings.Split(target, "/")[0]
			s3name = target
		}

		if vhost, _ := regexp.MatchString(s3vh, target); vhost {
			target = strings.Replace(target, ".s3.amazonaws.com", "", 1)
			target = strings.Split(target, "/")[0]
			s3name = target
		}

		if url, _ := regexp.MatchString(s3url, target); url {
			target = strings.Replace(target, "s3://", "", 1)
			target = strings.Split(target, "/")[0]
			s3name = target
		}

		var result string

		switch options.Format {
		case "path":
			result = "https://s3.amazonaws.com/" + s3name
		case "name":
			result = s3name
		case "url":
			result = "s3://" + s3name
		case "vhost":
			result = s3name + ".s3.amazonaws.com"
		default:
			result = bucket
		}

		if _, exists := uniqueResults.Load(result); exists {
			continue
		}

		fmt.Println(result)

		uniqueResults.Store(result, struct{}{})
		uniqueBuckets.Store(bucket, struct{}{})
	}
}
