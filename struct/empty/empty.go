package main

import "fmt"

func main() {
	wait := make(chan struct{})
	go func() {
		fmt.Println("送信")
		wait <- struct{}{}
	}()

	fmt.Println("受信待ち")
	<-wait
	fmt.Println("受信完了")
}
