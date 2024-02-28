package main

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"Project-1-Golang/helpers"
	"os"
	"strconv"
)

func main() {
	// Mendapatkan nomor absen dari argumen terminal
	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Argumen tidak valid:", os.Args[1])
		return
	}
	// helpers.getData()
	// Mengambil data teman
	teman, err := helpers.Data()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Menampilkan data teman berdasarkan nomor absen
	if absen > 0 && absen <= len(teman) {
		fmt.Printf("Data Teman ke-%d\n", absen)
		fmt.Println("Nama:", teman[absen-1].Nama)
		fmt.Println("Alamat:", teman[absen-1].Alamat)
		fmt.Println("Pekerjaan:", teman[absen-1].Pekerjaan)
		fmt.Println("Alasan:", teman[absen-1].Alasan)
	} else {
		fmt.Println("Nomor absen tidak valid:", absen)
	}
}


