package expedition

import "time"

// getNowTime is a function to get the current time.
func getNowTime() string {
	t := time.Now()
	const layout = "2006-01-02 15:04:05.757"
	return t.Format(layout)
}
