package main

import (
	"fmt"
	"hacktiv/modules"
	"os"
)

func main() {
	data := []string{
		modules.Biodata{
			Nama:      "Adi",
			Alamat:    "Jakarta",
			Pekerjaan: "Pelajar",
			Alasan:    "Mau Belajar Coding",
		},
		modules.Biodata{
			Nama:      "Doni",
			Alamat:    "Bandung",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Mau Belajar Golang",
		},
	}

	// args := os.Args

	fmt.Println(data[1], os.Args[1])
}
