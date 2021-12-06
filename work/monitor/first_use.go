package main

import (
	"fmt"
	"github.com/blurooo/go-monitor"
	"time"
)

// 注册得到一个上报客户端用于http服务质量监控
var httpReportClient = monitor.Register(monitor.ReportClientConfig {
	Name: "http服务监控",
	StatisticalCycle: 100,  // 每100ms统计一次服务质量
})

func main() {
	t := time.NewTicker(100 * time.Millisecond)
	for curTime := range t.C {
		// 每10ms向http监控客户端上报一条http服务数据，耗时0-100ms，状态为200
		fmt.Println("curTime = ",curTime)
		httpReportClient.Report("GET - /app/api/users", uint32(curTime.Nanosecond() % 100), 200)
	}
}