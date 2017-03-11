package utils

import "time"

func GetTimeNow() int64 {
	return time.Now().Unix()
}
