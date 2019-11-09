package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Mahasiswa struct {
	nama   string
	nim    string
	matkul [10]Matakuliah
}

type arrMahasiswa struct {
	data [1000]Mahasiswa
	num  int
}

type Matakuliah struct {
	nama  string
	nilai [3]int
	sks   int
}

var (
	mhs arrMahasiswa
)

func main() {
	var (
		pil, sure, paused string
	)
	clearScreen()

	// jumlah mahasiswa
	mhs.num = 0
	fmt.Println("Selamat Datang Di Aplikasi Nilai Mahasiswa")
	menu()
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pil)
	clearScreen()
	for pil != "0" {
		if pil == "1" {
			// input identitas mahasiswa
			loadData(&mhs)
		} else if pil == "2" {

		} else if pil == "3" {

		} else if pil == "4" {

		} else if pil == "5" {

		} else if pil == "6" {

		} else {
			fmt.Println("Inputan tidak valid")
		}
		fmt.Println("press any button")
		fmt.Scanln(&paused)
		clearScreen()

		menu()
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pil)
		clearScreen()

		if pil == "0" {
			fmt.Print("Anda yakin untuk keluar?[Y/N]: ")
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

//Fungsi untuk input data mahasiswa
func loadData(mhs *arrMahasiswa) {
	var n int

	fmt.Print("Jumlah mahasiswa : ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Nama	: ")
		fmt.Scanln(&mhs.data[i].nama)
		fmt.Print("NIM	: ")
		fmt.Scanln(&mhs.data[i].nim)
		fmt.Println("saved")
		fmt.Println()
		mhs.num++
	}

}

//Menampilkan data mahasiswa
func showData(jumlahMhs int, dataMhs arrMahasiswa) {
	for i := 0; i < jumlahMhs; i++ {
		fmt.Printf("Mahasiswa ke %d : \n", i+1)
		fmt.Println("Nama	:", dataMhs.data[i].nama)
		fmt.Println("Nim	:", dataMhs.data[i].nim)
	}
}
