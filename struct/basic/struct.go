package basic

import (
	"fmt"
	"time"
)

type Author struct {
	FirstName string
	LastName  string
}

type Book struct {
	Title     string
	Author    Author
	Publisher string
	RelesedAt time.Time
	ISBN      string
}

func main() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	book := Book{
		Title: "Real World HTTP",
		Author: Author{
			FirstName: "よしき",
			LastName:  "渋川",
		},
		Publisher: "オライリー・ジャパン",
		ISBN:      "48888888",
		RelesedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book.Title)

	internalBook := struct {
		Title      string
		Author     string
		Publisher  string
		ISBN       string
		ReleasedAt time.Time
	}{
		Title:      "Real World HTTP",
		Author:     "渋川よしき",
		Publisher:  "オライリー・ジャパン",
		ISBN:       "5939395393",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(internalBook.Title)
}
