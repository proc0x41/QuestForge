package main

import (
	"bufio"
	"os"
)

func wait(msg string) {
	println(msg)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func Gui() {
	banner()
	difficulty := dialog(`
Escolha a dificuldade da questão:
1 - Fácil
2 - Média
3 - Difícil
`)
	print(difficulty)
	wait("")
}

func main() {
	Gui()
}
