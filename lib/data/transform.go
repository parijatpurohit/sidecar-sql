package data

import (
	"regexp"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func GetTimestamp(ts *timestamp.Timestamp) *time.Time {
	if ts != nil {
		res, _ := ptypes.Timestamp(ts)
		return &res
	}
	return nil
}

func GetTimePtr(t time.Time) *time.Time {
	return &t
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
