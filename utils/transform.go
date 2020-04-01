package utils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func GetTimestamp(ts *timestamp.Timestamp) *time.Time {
	if ts != nil {
		res, _ := ptypes.Timestamp(ts)
		return &res
	}
	return nil
}
