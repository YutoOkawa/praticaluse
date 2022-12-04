package combine

import (
	"log"
	"strings"
)

func main() {
	src := []string{"Back", "To", "The", "Future", "Part", "Ⅲ"}
	var title string

	for i, word := range src {
		if i != 0 {
			title += " "
		}
		title += word
	}
	log.Println(title)

	var builder strings.Builder
	builder.Grow(100)
	for i, word := range src {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	log.Println(builder.String())
}
