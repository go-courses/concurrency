package main

import "fmt"

type (
	Board interface {
		Print()
	}
	BlueBoard struct {
		color string
	}
	DarkBoard struct {
		size string
	}
)

func (b *BlueBoard) Print() {
	fmt.Printf("blue board with color %s\n", b.color)
}
func (b *DarkBoard) Print() {
	fmt.Printf("dark board with size %s\n", b.size)
}
func main() {
	blue := &BlueBoard{color: "red"}
	dark := &DarkBoard{size: "big"}
	printBoard(blue)
	printBoard(dark)

}

func printBoard(b Board) {
	b.Print()
}
