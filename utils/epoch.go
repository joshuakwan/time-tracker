package utils

import "time"

func currentEpochTime() int64 {
	return time.Now().Unix()
}

func humanReadableToEpoch(date time.Time) int64 {
	return date.Unix()
}

func epochToHumanReadable(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}
