package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

///////////////////////////////////
//هر تایپ محدود به انجام کار مشخص خودش باید باشه
var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount,
		text)
	j.entries = append(j.entries, entry)
	return entryCount
}

///////////////////////////////////////////

// ///مثلا این متود داره نقض میکنه مسئولیت تایپ ژورنال رو و داره مسئولیت سیو کردن روی فایل رو هم بهش اضافه میکنه

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

// Persistence /////// کار صحیح این هست که برای ماندگار کردن ژورنال از یه تایپ دیگه استفاده بشه
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("test entry")
	j.AddEntry("another test entry")
	fmt.Println(strings.Join(j.entries, "\n"))

	//
	p := Persistence{"\n"}
	p.saveToFile(&j, "journal.txt")
}
