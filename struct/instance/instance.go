package instance

type Person struct {
	Name string
}

// ファクトリー関数
func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}

func main() {
	// newで作成(あまり使わない)
	p1 := new(Person)

	// var変数宣言で作成
	var p2 Person

	// 複合リテラルで初期化
	p3 := &Person{
		Name: "石田",
	}

	var p4 *Person
}
