package main

import (
	"fmt"
	"strings"
)

func main() {
	// Variable berisi string, masukan kalimat yang ingin dicoba
	kalimat := "selamat malam warga Tangerang"

	// Menghapus spasi dan memecah kalimat menjadi huruf
	huruf := strings.ReplaceAll(kalimat, " ", "")
	
	// Membuat map untuk menghitung kemunculan huruf
	count := make(map[rune]int)

	// Looping untuk menghitung kemunculan setiap huruf dan menampilkan setiap huruf
	for _, h := range huruf {
		count[h]++
		fmt.Println(string(h))
	}

	// Menampilkan hasil perhitungan
	fmt.Print("map: [")
	first := true
	for k, v := range count {
		if !first {
			fmt.Print(" ")
		}
		fmt.Printf("%c:%d", k, v)
		first = false
	}
	fmt.Println("]")
}