package helpers

var person []Person

// Slice data berdasarkan inisialisasi data dari Struct Person
func Init() {
	person = []Person{
		{ No: 1, Name: "Syahra", Address: "Jakarta", Job: "Frontend Developer", ReasonGoClass: "want to be fullstack developer"},
		{ No: 2, Name: "Nadia", Address: "Surabaya", Job: "Frontend Developer", ReasonGoClass: "want to be fullstack developer"},
		{ No: 3, Name: "Syela", Address: "Yogyakarta", Job: "Frontend Developer", ReasonGoClass: "want to be fullstack developer"},
	}
}