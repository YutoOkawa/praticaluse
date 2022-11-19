package control

import (
	"flag"
	"log"
)

var (
	FlagStr = flag.String("string", "default", "文字フラグ")
	FlagInt = flag.Int("int", -1, "数値フラグ")
)

func main() {
	flag.Parse()
	log.Println(*FlagStr)
	log.Println(*FlagInt)
	log.Println(flag.Args())
}
