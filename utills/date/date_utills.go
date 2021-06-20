package date

import "time"
const(
	apiTimeFormat="2021-06-13T15:04:05Z"
)
func GetTime() time.Time {
	return time.Now().UTC()
}

func GetTimeString() string{
	return GetTime().Format(apiTimeFormat)
}