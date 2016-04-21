package main

import (
	"fmt"
	"time"
)

// Datetie 対応表に格納する構造体
type Datetime struct {
	src  string // time.Format()側のフォーマット
	dest string // PHPなどで使用されている%[A-Za-z]なフォーマット
}

const DatetimeFormat = "2006-01-02 15:04:05"

const DateFormat = "2006-01-02"

const DefaultFormat = "2006-1-2 3:4:5"

func main() {

	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)

	fmt.Printf("%q\n", now.Format(DatetimeFormat))
	fmt.Printf("%q\n", now.Format(DateFormat))
	fmt.Printf("%q\n", Format(now, "%Y-%m-%d %H:%i:%s"))
	fmt.Printf("%q\n", Format(now, "%Y-%m-%d"))
}

// Time.Format()のAlias関数
func Format(datetime time.Time, format string) string {
	// TODO:本来はformatを""%[a-zA-z]"な形で正規表現でパース
	// TODO:対応表(constで定義)を用意しておき対応した形でFormatなされていればOK
	if format == "%Y-%m-%d %H:%i:%s" {
		return datetime.Format(DatetimeFormat)
	}
	return datetime.Format(DefaultFormat)
}

// TODO:曜日を返す関数(日本語/英語両対応)
// TODO:指定月の末日を返す
// TODO:指定月の頭日を返す
// TODO:フォーマット部分のみ置換して返す
