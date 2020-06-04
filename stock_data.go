package main

import (
    "github.com/xiaotian/stock/pkg/collector"
    "github.com/xiaotian/stock/pkg/persistent"
    "github.com/xiaotian/stock/pkg/s-logger"
    "github.com/xiaotian/stock/pkg/tool"
    "time"
)

var logger = s_logger.New()
var companyFile = persistent.CompanyFilePreserver{}
var dataFile = persistent.DataFilePreserver{}

func main() {
    logger.Infow("项目启动")

    for {
        now := time.Now()
        duration := nextRunDurationZh(now)
        timer := time.NewTimer(duration)
        logger.Infof("项目定时器设置成功,定时器信息: %v", timer)
        logger.Infow("项目定时器设置成功.", "nextRunTime", tool.DateTime(now.Add(duration)))
        <-timer.C
    logger.Infow("收集公司信息触发执行")
    companyInfos := collector.CollectCompanyInfo()
    if err := companyFile.Save(companyInfos); err != nil {
        logger.Errorw("保存数据失败", "error", err)
    }

    dataList := collector.CollectData(companyInfos)
    if err := dataFile.Save(dataList); err != nil {
        logger.Errorw("保存数据失败", "error", err)
    }
    }
}

func nextRunDurationZh(now time.Time) time.Duration {
    return nextRunDuration(now, 15, 30, 0)
}

func nextRunDuration(now time.Time, hour int, minute int, second int) time.Duration {
    workDay := now
    if now.Hour() < 15 {
        if isWeekEndDay(now) {
            workDay = nextWorkDay(now)
        }
        date := time.Date(workDay.Year(), workDay.Month(), workDay.Day(), hour, minute, second, 0, time.Local)
        return date.Sub(now)
    }

    workDay = nextWorkDay(now)
    date := time.Date(workDay.Year(), workDay.Month(), workDay.Day(), hour, minute, second, 0, time.Local)
    return date.Sub(now)
}

func nextWorkDay(t time.Time) time.Time {
    oneDayDuration := time.Hour * 24
    if t.Weekday() == time.Friday {
        return t.Add(oneDayDuration * 3)
    }
    if t.Weekday() == time.Saturday {
        return t.Add(oneDayDuration * 2)
    }
    return t.Add(oneDayDuration)
}

func isWeekEndDay(t time.Time) bool {
    return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}
