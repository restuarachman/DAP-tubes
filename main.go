package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Namalengkap struct {
	lengkap [6]string
}
type Mahasiswa struct {
	nama   Namalengkap
	nim    string
	matkul [10]Matakuliah
}
type Matakuliah struct {
	nama  string
	nilai int
	sks   int
}

func main() {
	var pil, sure string
	clearScreen()

	fmt.Println("Selamat Datang Di Aplikasi Nilai Mahasiswa")
	menu()
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pil)
	clearScreen()
	for pil != "0" {
		if pil == "1" {

		} else if pil == "2" {

		} else if pil == "3" {

		} else if pil == "4" {

		} else if pil == "5" {

		} else if pil == "6" {

		} else {
			fmt.Println("Inputan tidak valid")
		}
		fmt.Scanln()
		clearScreen()

		menu()
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pil)
		clearScreen()
		if pil == "0" {
			fmt.Println("Anda yakin untuk keluar?[Y/N]: ")
			fmt.Scanln(&sure)
			if sure == "N" || sure == "n" {
				menu()
				fmt.Print("Pilihan : ")
				fmt.Scanln(&pil)
				clearScreen()
			}
		}
	}
}

func clearScreen() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

func menu() {
	fmt.Println("-------------Menu-------------")
	fmt.Println("1. Input Data Mahasiswa")
	fmt.Println("2. Edit Data Mahasiwa")
	fmt.Println("3. Sorting Data Mahasiswa")
	fmt.Println("4. Cari Data Mahasiswa")
	fmt.Println("5. Lihat Semua Data Mahasiswa")
	fmt.Println("6. Hapus Data Mahasiswa")
	fmt.Println("0. Keluar")
}
