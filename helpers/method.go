package helpers

import "fmt"

// Cari data berdasarkan nama
func DisplayClassmateByName(Name string) {
	for _, p := range person {
		if p.Name == Name {
			fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan memilih kelas Golang: %s\n", p.Name, p.Address, p.Job, p.ReasonGoClass)
			return
		}
	}
	fmt.Println("Data dengan nama tersebut tidak tersedia.")
}

// Cari data berdasarkan nomor
func DisplayClassmateByNo(No int) {
	for _, p := range person {
		if p.No == No {
			fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan memilih kelas Golang: %s\n", p.Name, p.Address, p.Job, p.ReasonGoClass)
			return
		}
	}
	fmt.Println("Data dengan nomor tersebut tidak tersedia.")
}