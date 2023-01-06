package method

type Struct struct {
	v int
}

func (s Struct) PrintStatus() {
	log.Println("Struct:", s.v)
}

// インスタンスのレシーバ
// 値を変更してもインスタンスのフィールドは変更されない
func (s Struct) SetValue(v int) {
	s.v = v
}

// ポインタのレシーバ
// 値を変更できる
func (s *Struct) SetValue(v int) {
	s.v = v
}

type StructWithPointer struct {
	v *int
}

func (a StructWithPointer) Modify() {
	(*a.v) = 10
}

func String[T any](s T) {
	return fmt.Sprintf("%v", s)
}

func main() {
	s := StructWithPointer{}
	i := 1
	s.v = &i
	s.Modify()
	fmt.Println(*s.v)
}
