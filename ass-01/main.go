package main

import (
	"fmt"
	"os"
)

// create struct student
type Student struct {
	Id, Nama, Pekerjaan, Alamat, Alasan string
}

func main() {
	var lists = []Student{
		{Id: "1", Nama: "Joko", Pekerjaan: "Karyawan", Alamat: "Jl. Satu", Alasan: "menambah skill"},
		{Id: "2", Nama: "Janu", Pekerjaan: "Karyawan", Alamat: "Jl. Dua", Alasan: "tuntutan pekerjaan"},
		{Id: "3", Nama: "Indah", Pekerjaan: "Mahasiswa", Alamat: "Jl. Tiga", Alasan: "menambah porto"},
		{Id: "4", Nama: "Puspa", Pekerjaan: "Mahasiswa", Alamat: "Jl. Empat", Alasan: "menambah skill"},
	}

	idStudent := os.Args[1] //mengambil input data setelah .go
	ShowStudent(lists, idStudent)
}

func ShowStudent(list []Student, id string) {
	for i := range list {
		//mencari data berdasarkan id
		if list[i].Id == id {
			fmt.Printf("Nama\t\t: %s \nPekerjaan\t: %s\nAlamat\t\t: %s\nAlasan\t\t: %s\n", list[i].Nama, list[i].Pekerjaan, list[i].Alamat, list[i].Alasan)
		}
	}
}
