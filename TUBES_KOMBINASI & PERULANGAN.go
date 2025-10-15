package main

import (
	"fmt"   // Untuk input & output
	"strings" // Untuk mengonversi string ke uppercase
)

func main() {
	var jumlah_siswa int
	fmt.Println("Masukkan jumlah siswa:")
	fmt.Scanln(&jumlah_siswa) // Input jumlah siswa

	for i := 1; i <= jumlah_siswa; i++ {
		var nilai int
		fmt.Printf("Masukkan nilai siswa ke-%d: ", i)
		fmt.Scanln(&nilai) // Input nilai sebagai integer

		fmt.Println("Pilih jurusan: MIPA, IPS, atau BAHASA")
		var jurusan string
		fmt.Scanln(&jurusan) // Input jurusan
		jurusan = strings.ToUpper(jurusan) // Konversi ke uppercase untuk toleransi case

		if jurusan == "MIPA" || jurusan == "IPS" || jurusan == "BAHASA" {
			if nilai >= 87 {
				kelas := "A"
				fmt.Printf("Anda berada di jurusan %s dan berada di %s\n", jurusan, kelas)
			} else if nilai >= 75 && nilai < 87 {
				kelas := "B"
				fmt.Printf("Anda berada di jurusan %s dan berada di %s\n", jurusan, kelas)
			} else if nilai >= 60 && nilai < 75 {
				kelas := "C"
				fmt.Printf("Anda berada di jurusan %s dan berada di %s\n", jurusan, kelas)
			} else {
				fmt.Println("Maaf anda belum diterima")
			}
		} else {
			fmt.Println("Jurusan tidak tersedia dalam pilihan")
		}
	}
}

