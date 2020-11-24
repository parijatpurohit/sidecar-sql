package data

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

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
