package util

import "time"

// 获取时间戳(单位:秒 s)
func TimeStamp() int64 {
	return time.Now().Unix()
}

// 获取时间戳（单位:毫秒 ms）
func TimeStampMs() int64 {
	return time.Now().UnixNano() / 1e6
}

// 获取时间戳(单位：微秒 us)
func TimeStampUs() int64 {
	return time.Now().UnixNano() / 1e3
}

// 获取时间戳(单位：纳秒 ns)
func TimeStampNs() int64 {
	return time.Now().UnixNano()
}
