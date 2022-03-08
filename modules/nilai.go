package modules

type Student struct {
	Nama                            string
	NilaiIps, NilaiIpa, NilaiBahasa int
}

func (s *Student) Ips() {
	s.Nama = "Faris"
	s.NilaiIps = 90
}

func (s *Student) Ipa() {
	s.Nama = "Fahmi"
	s.NilaiIpa = 95
}

func (s *Student) Bahasa() {
	s.Nama = "Faris Fahmi"
	s.NilaiBahasa = 100
}
