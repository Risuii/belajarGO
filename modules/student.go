package modules

import "fmt"

type Siswa struct {
	Nama         string
	Nilai        int
	JenisKelamin string
}

func (s Siswa) NilaiSiswa() {
	switch {
	case s.Nilai >= 80:
		fmt.Println("Nama:", s.Nama, "Jenis Kelamin:", s.JenisKelamin, "Hasil:", s.Nilai, "A")
	case s.Nilai < 80 && s.Nilai >= 60:
		fmt.Println("Nama:", s.Nama, "Jenis Kelamin:", s.JenisKelamin, "Hasil", s.Nilai, "B")
	case s.Nilai < 60:
		fmt.Println("Nama:", s.Nama, "Jenis Kelamin:", s.JenisKelamin, "Hasil", s.Nilai, "C")
	default:
		fmt.Println("Tidak lulus")
	}
}
