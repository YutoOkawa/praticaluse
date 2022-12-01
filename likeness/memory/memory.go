package memory

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, fname := range files {
		f, err := os.Open(fname)
		if err != nil {
			return err
		}
		// この書き方だとClose()は全部のループを抜けるまで実行されない
		// deferを使わずにファイルをつかっtあ後に自分でf.Close()を呼ぶ
		defer f.Close()
		data, _ := io.ReadAll(f)
		result = append(result, data)
	}
}

func deferReturnSample(fname string) (err error) {
	var f *os.File
	f, err = os.Create(fname)
	if err != nil {
		return fmt.Errorf("ファイルオープンのエラー %w", err)
	}
	defer func() {
		// Close()のエラーを拾って名前付き返り値に代入
		// すでにerrに別のものが入る可能性があるときは
		// さらに要注意
		err = f.Close()
	}()
	io.WriteString(f, "deferのエラーを拾うサンプル")
	return
}
