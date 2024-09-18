package main

import "time"

func main() {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()
	for i := 0; i < 4; i++ {
		<-ticker.C
		hoge()
		println("tick")
	}
}

func hoge() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	<-ticker.C
	println("hoge")
}
