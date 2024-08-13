package main

import (
	"ClassmateCLI/helpers"
	"fmt"
	"os"
	"strconv"
)

func main() {
    helpers.Init()

		// Cek apakah ada argumen
    if len(os.Args) < 2 {
        fmt.Println("Gunakan: go run main.go <nama> atau go run main.go <nomor>")
        return
    }

		// Ambil argumen
    arg := os.Args[1]

    // Coba konversi argumen ke integer untuk pencarian berdasarkan nomor
    if no, err := strconv.Atoi(arg); err == nil {
        helpers.DisplayClassmateByNo(no)
    } else {
        // Jika konversi gagal, anggap argumen sebagai nama
        helpers.DisplayClassmateByName(arg)
    }
}