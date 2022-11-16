package usevar

var a = 1             // イコールの右側の1が定数(defaultがint)
var b int32 = 1 + 100 // 定数の中で演算子を使った計算も可能(明示的にかける)

// 定数に名前がつけられる
// コンパイル時に値が決まらないものは使えない
const (
	c     = 9223372036854775807 + 1 // uint64を超える数値も定義可能
	d     = "hello world"
	f int = iota + 10 // const宣言でのみ`iota`が使える
// g int32 = 4294967295 + 1 // 型の範囲を超えるためエラー
// h []int = []int{1, 2, 3} // 配列やスライスもエラー
//
//	i map[string]int = map[string]int{
//		"Tokyo":    10,
//		"Kanagawa": 11,
//		"Chiba":    12,
//	} // mapもエラー
//
// j = function() // 関数の返り値もエラー
)

type errDatabase int

func (e errDatabase) Error() string {
	return "Database Error"
}

// errors.New()ではコンパイルエラーになるので、定数定義ではこのように定義型を使う
// 利用者は==演算子を使ってエラーの種別を比較できる
const (
	ErrDatabase errDatabase = 0
)
