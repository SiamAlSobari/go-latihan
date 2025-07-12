package main

import "fmt"

func main(){
	fmt.Println("Siapa nama kamu")
	var name string
	fmt.Scan(&name)
	fmt.Println("berapa umur kamu")
	var umur int
	fmt.Scan(&umur)

	var penentu string
	if umur >= 18 {
		penentu = "sudah dewasa"
	} else if umur >= 13 {
		penentu = "sudah remaja"
	} else {
		penentu = "belum dewasa"
	}

	fmt.Println("Halo", name, "umur kamu", umur, "tahun, kamu", penentu)
}