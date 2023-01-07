package embed

type Book struct {
	Title string
	ISBN  string
}

func (b Book) GetAmazonURL() string {
	return "https://amazon.co.jp/dp/" + b.ISBN
}

type OreillyBook struct {
	Book
	// Book book ともできるがその場合はbookを経由しなければならない
	ISBN13 string
}

func (o OreillyBook) GetOreillyURL() string {
	return "https://www.oreilly.co.jp/books/" + o.ISBN13 + "/"
}

func main() {
	ob := OreillyBook{
		ISBN13: "8489498249824984982",
		Book: Book{
			Title: "Read World HTTP",
		},
	}

	// Bookのメソッドが利用可能
	fmt.Println(ob.GetAmazonURL())
	// OreillyBookのメソッドが利用可能
	fmt.Println(ob.GetOreillyURL())
}
