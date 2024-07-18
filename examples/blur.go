package main

import "github.com/ioi-hub/imgo"

func main() {
	imgo.Load("gopher.png").
		Blur(5).
		Save("out.png")
}
