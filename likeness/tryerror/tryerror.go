package tryerror

import (
	"bufio"
	"os"
)

func tryerror() {
	f, err := os.Open("important.txt")
	if err != nil {
		// エラーハンドリング
	}

	r := bufio.NewReader(f)
	l, err := r.ReadString('\n')
	if err != nil {
		// エラーハンドリング
	}
	// ここでエラーは発生していない
}
