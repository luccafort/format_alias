package main

import (
	"fmt"
	"time"
)

// Datetime 対応表に格納する構造体
type Datetime struct {
	src  string // time.Format()側のフォーマット
	dest string // PHPなどで使用されている%[A-Za-z]なフォーマット
}

// DatetimeZeroFormat Y-m-d H:i:s形式
const DatetimeZeroFormat = "2006-01-02 15:04:05"

// DateZeroFormat Y-m-d形式
const DateZeroFormat = "2006-01-02"

// DatetimeFormat 0つきでない
const DatetimeFormat = "2006-1-2 3:4:5"

// DateFormat 0つきでない
const DateFormat = "2006-1-2"

// DefaultFormat Goのデフォルトな日付フォーマット
const DefaultFormat = DatetimeFormat

// WeekdayList 日本語の曜日リスト
const WeekdayList = [...]string{"日", "月", "火", "水", "木", "金", "土"}

func main() {

	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)

	//fmt.Printf("%q\n", now.Format(DatetimeFormat))
	//fmt.Printf("%q\n", now.Format(DateFormat))
	fmt.Printf("%q\n", Format(now, "Y-m-d H:i:s"))
	fmt.Printf("%q\n", Format(now, "Y-m-d"))
}

// Format Time.Format()のAlias関数
func Format(datetime time.Time, format string) string {
	// TODO:正規表現で実装したかったけどちょっとやり方が今わからないので完全一致で実装
	formatList := [][]string{
		{"Y-m-d H:i:s", DatetimeZeroFormat},
		{"Y-n-j h:i:s", DatetimeFormat},
		{"Y-m-d", DateZeroFormat},
		{"Y-n-j", DateFormat},
	}

	var result string
	for _, array := range formatList {
		if format == array[0] {
			result = array[1]
		}
	}

	if result == "" {
		result = DefaultFormat
	}
	return datetime.Format(result)
}

// WeekdayName 日本語の曜日を返す
func WeekdayName(datetime time.Time) string {
	return WeekdayList[datetime.Weekday()]
}

// isLeapYear() 閏年を判定
func isLeapYear(datetime time.Time) bool {
	year := datetime.Year()
	switch year {
	case year%400 == 0:
		fallthrough
	case year%4 == 0:
		// 閏年
		return true
	case year%100 == 0:
		fallthrough
	default:
		// 閏年じゃない
		return false
	}
}
