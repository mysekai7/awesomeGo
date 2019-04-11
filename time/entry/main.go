package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	at := GetNow()
	log.Printf("at: %+v", at)
	return

	//获取本地location
	toBeCharge := "2015-01-01 00:23:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)                                                 //打印输出时间戳 1420041600

	//时间戳转日期
	dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)

	fmt.Println(time.Local)

}

// AuditedTime Struct containing created and updated times
type AuditedTime struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateTime mark the start
func (t *AuditedTime) CreateTime() time.Time {
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}
	return t.CreatedAt
}

// UpdateTime mark as updated. Will set started time if not set
func (t *AuditedTime) UpdateTime() time.Time {
	t.CreateTime()
	t.UpdatedAt = time.Now()
	return t.UpdatedAt
}

// GetNow return auditedTime for creation
func GetNow() AuditedTime {
	now := AuditedTime{}
	now.UpdateTime()
	return now
}
