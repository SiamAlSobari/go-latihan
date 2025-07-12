package main

import "fmt"

func main(){
	var nama string
	var umur int
	var jenis_kelamin string

	fmt.Println("Masukan nama kamu : ")
	fmt.Scan(&nama)
	fmt.Println("Masukan umur kamu : ")
	fmt.Scan(&umur)
	fmt.Println("Masukan jenis kelamin kamu P/L : ")
	fmt.Scan(&jenis_kelamin)

	var umur_penentu string
	var jenis_kelamin_penentu string

	switch  {
	case umur >= 18:
		umur_penentu = "sudah dewasa"
	 case umur >= 13:
		umur_penentu = "sudah beranjak remaja"
	 case umur >= 7:
		umur_penentu = "sudah beranjak remaja"
	 default:
		umur_penentu = "belum masih anak-anak"
	}

	switch jenis_kelamin {
		case "P", "p":
			jenis_kelamin_penentu = "perempuan"
		case "L", "l":
			jenis_kelamin_penentu = "laki-laki"
		
	}

	fmt.Println("Halo", nama, "umur kamu", umur, "tahun, kamu", umur_penentu, "dan jenis kelamin kamu", jenis_kelamin_penentu)
}