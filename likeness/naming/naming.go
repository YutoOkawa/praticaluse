// package名にはutil/common/apiなどの名前は使わない
// このパッケージがどのようなものなのかを明らかにならない名前はつけない
// 複数の用途に使われるのであれば、encoding/xmlのように末端のパッケージ名を分ける
// http.HTTPServerのように、パッケージ名と重複があるのは冗長
package naming

import "errors"

// var max_length int // Goではこう書かない
var maxLength int // Goらしい変数名のつけ方
var MaxLength int // パッケージ外から参照できる(public変数)

var URL string // 頭字語は全て大文字/小文字
// var Url String // このようにしない

var MarshalError error              // error型にはErrorとつける
var ErrTooLong = errors.New("Test") // エラーの変数にはErrとつける

var req string          // requestよりも短いreqが好まれる(ローカル変数であれば)
var smallBufferSize int // グローバル変数にはわかりやすい名前をつける

// 一つのメソッド飲みを持つインターフェースには`er`をつける
type Reader interface {
	read()
}
