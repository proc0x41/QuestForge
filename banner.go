package main

import "fmt"

func banner() {
	println("   ____                  _   ______                   ")
	println("  / __ \\                | | |  ____|                  ")
	println(" | |  | |_   _  ___  ___| |_| |__ ___  _ __ __ _  ___ ")
	println(" | |  | | | | |/ _ \\/ __| __|  __/ _ \\| '__/ _` |/ _ \\")
	println(" | |__| | |_| |  __/\\__ \\ |_| | | (_) | | | (_| |  __/")
	println(`  \___\_\\__,_|\___||___/\__|_|  \___/|_|  \__, |\___|`)
	println("                                            __/ |     ")
	println("                                           |___/      ")

	println("Gerador de Lista de Questões - by proc0x40")
}

func dialog(text string) int {
	println(text)
	println("Digite uma opção: ")
	var response int
	fmt.Scan(&response)
	return response
}
