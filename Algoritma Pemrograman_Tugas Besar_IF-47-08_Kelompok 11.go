// Package dan Import

package main

import (
	"fmt"
	"os"
	"time"
)

// Program Utama

func main() {
	menuLoginRegister()
}

// Inisialisasi Data

// 1. Data User

const NMAX int = 5

type structUser struct {
	nama, username, password, status string
}

type arrayUser [NMAX]structUser

type structDataUser struct {
	aUser arrayUser
	n     int
}

var dataUser structDataUser

// 2. Data Masyarakat

type structDataMain struct {
	nama, asal, gender, status string
	umur                       int
	tinggi                     float64
	tanggal                    string
	stringTemp                 string
	intTemp                    int
	floatTemp                  float64
}

type arrDataMain [NMAX]structDataMain

type structMainArray struct {
	aData  arrDataMain
	n      int
	sort   bool
	sortby string
	Order  string
}

var dataMain structMainArray

// 3. Data Lain

const maxBulan int = 24

type indkatorStruct struct {
	severe, stunted float64
}

type indikasiStuntingPerBulan [maxBulan]indkatorStruct

type indikasiStunting struct {
	A    indikasiStuntingPerBulan
	nama string
}

var WHOIndicator = indikasiStunting{
	A: indikasiStuntingPerBulan{
		{2, 2.4}, {2.7, 3.2}, {3.4, 3.9}, {4, 4.5}, {4.4, 5}, {4.8, 5.4}, {5.1, 5.7}, {5.3, 6}, {5.6, 6.3}, {5.8, 6.5}, {5.9, 6.7}, {6.1, 6.9},
		{6.3, 7}, {6.4, 7.2}, {6.6, 7.4}, {6.7, 7.6}, {6.9, 7.7}, {7, 7.9}, {7.2, 8.1}, {7.3, 8.2}, {7.5, 8.4}, {7.6, 8.6}, {7.8, 8.7}, {7.9, 8.9},
	},
	nama: "WHO Indicator",
}

var KemenKesIndicator = indikasiStunting{
	A: indikasiStuntingPerBulan{
		{1.6, 1.92}, {2.16, 2.56}, {2.72, 3.12}, {3.2, 3.6}, {3.52, 4}, {3.84, 4.32}, {4.08, 4.56}, {4.24, 4.8}, {4.48, 5.04}, {4.64, 5.2},
		{4.72, 5.36}, {4.88, 5.52}, {5.04, 5.6}, {5.12, 5.76}, {5.28, 5.92}, {5.36, 6.08}, {5.52, 6.16}, {5.6, 6.32}, {5.76, 6.48}, {5.84, 6.56},
		{6, 6.72}, {6.08, 6.88}, {6.24, 6.96}, {6.32, 7.12},
	},
	nama: "KemenKes Indicator",
}

var currentIndicator indikasiStunting = WHOIndicator

var tipeSorting string = "Selection"

// Fungsi menghilangkan print

func clearLines(n int) {
	for i := 0; i < n; i++ {
		// Move the cursor up one line
		fmt.Print("\033[A")
		// Clear the line
		fmt.Print("\033[K")
	}
}

// Fungsi Center Teks [Untuk Tampilan Menu]

const width int = 45

const border string = "============================================="

func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	paddingStr := ""

	for i := 0; i < padding; i++ {
		paddingStr += " "
	}

	return paddingStr + text
}

// Menu Keluar Aplikasi

func menuExit() {

	fmt.Println(border)
	fmt.Println(centerText("Menu Exit", width))
	fmt.Println(border)
	fmt.Println("1. Iya")
	fmt.Println("2. Cancle")
	fmt.Println()
	fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")
	fmt.Println(border)

	var in int

	fmt.Scan(&in)
	switch in {
	case 1:
		clearLines(100)
		os.Exit(0)
	case 2:
		clearLines(100)
		menuLoginRegister()
	default:
		clearLines(100)
		menuExit()
	}
}

// Menu Login dan Registrasi

func menuLoginRegister() {

	fmt.Println(border)
	fmt.Println(centerText("Aplikasi Pendataan Kondisi Stunting", width))
	fmt.Println(border)
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Exit")
	fmt.Println()
	fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")
	fmt.Println(border)

	var in int
	fmt.Scan(&in)

	switch in {
	case 1:
		clearLines(100)
		loginAkun()

	case 2:
		clearLines(100)
		daftarAkun()

	case 3:
		clearLines(100)
		menuExit()

	default:
		clearLines(100)
		menuLoginRegister()

	}
}

// Subprogram Register

func daftarAkun() {

	fmt.Println(border)
	fmt.Println(centerText("Menu Registrasi", width))
	fmt.Println(border)

	if dataUser.n == NMAX {

		fmt.Println(border)
		fmt.Println(centerText("Data User telah penuh", width))
		fmt.Println(border)
		fmt.Println("1. Kembali ke menu utama")
		fmt.Println("2. Keluar dari aplikasi")
		fmt.Println()
		fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")
		fmt.Println(border)

		var in int
		fmt.Scan(&in)

		switch in {
		case 1:
			clearLines(100)
			menuLoginRegister()

		case 2:
			clearLines(100)
			menuExit()

		default:
			clearLines(100)
			daftarAkun()

		}
	}

	isiNama()

	isiUsername()

	isiPassword()

	isiStatus()

	cekRegister()

}

func isiNama() {
	// IS: Dijalankan ketika ingin mengisi nama pada suatu data User
	// FS: nama pada data user terisi

	fmt.Print("Nama Lengkap	: ")
	fmt.Scan(&dataUser.aUser[dataUser.n].nama)

}

func isiUsername() {
	// IS: Dijalankan ketika ingin mengisi username pada suatu data User
	// FS: username pada data user terisi
	/*
		Keterangan: Username merupakan data yang unik pada aplikasi, sehingga bila terdapat duplikat
					Subprogram ini akan meminta user mengisi ulang username lain yang unik
	*/

	var usernameDuplikat bool
	usernameDuplikat = true

	for usernameDuplikat {
		fmt.Print("Username	: ")
		fmt.Scan(&dataUser.aUser[dataUser.n].username)

		usernameDuplikat = univSearch_String(dataUser.aUser[dataUser.n].username, "username", dataUser.n) != -1

		if usernameDuplikat {
			fmt.Println("Username sudah ada di database, silahkan isi ulang")
		}
	}

}

func isiPassword() {
	// IS: Dijalankan ketika ingin mengisi password pada suatu data User
	// FS: password pada data user terisi

	fmt.Print("Password	: ")
	fmt.Scan(&dataUser.aUser[dataUser.n].password)

}

func isiStatus() {
	// IS: Dijalankan ketika ingin mengisi status pada suatu data User
	// FS: status pada data user terisi
	/*
		Keterangan: Status user hanya bisa terisi antara Tenaga Kesehatan (TenKes) atau Admin (Adm)
	*/

	fmt.Println(border)
	fmt.Println(centerText("Tentukan Status", width))
	fmt.Println(border)
	fmt.Println("1. Tenaga Kesehatan")
	fmt.Println("2. Admin")
	fmt.Println()
	fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")
	fmt.Println(border)
	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		dataUser.aUser[dataUser.n].status = "Tenaga Kesesehatan"

	case 2:
		dataUser.aUser[dataUser.n].status = "Admin"

	default:
		clearLines(100)
		isiStatus()
	}

}

func cekRegister() {
	// IS: Terdapat data suatu user
	// FS: Mengubah data berdasarkan input yang diminta user
	/*
		Keterangan: User dapat mengubah nama, username, password, atau status
	*/

	fmt.Println(border)
	fmt.Println("Nama		:", dataUser.aUser[dataUser.n].nama)
	fmt.Println("Username	:", dataUser.aUser[dataUser.n].username)
	fmt.Println("Password	:", dataUser.aUser[dataUser.n].password)
	fmt.Println("Status		:", dataUser.aUser[dataUser.n].status)
	fmt.Println(border)
	fmt.Println("Apakah akun sudah sesuai? (Y/T)")
	fmt.Println()
	fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")
	fmt.Println(border)

	var inS string
	fmt.Scan(&inS)
	switch inS {
	case "Y":
		clearLines(100)
		dataUser.n++
		menuLoginRegister()
	case "T":
		clearLines(100)
		fmt.Println(border)
		fmt.Println(centerText("Menu Ubah Data User", width))
		fmt.Println(border)
		fmt.Println("1. Nama")
		fmt.Println("2. Username")
		fmt.Println("3. Password")
		fmt.Println("4. Status")
		fmt.Println()
		fmt.Println("Ketik opsi yang ingin diubah, lalu tekan Enter")
		fmt.Println(border)

		var in int
		fmt.Scan(&in)
		switch in {
		case 1:
			clearLines(100)
			isiNama()

		case 2:
			clearLines(100)
			isiUsername()

		case 3:
			clearLines(100)
			isiPassword()

		case 4:
			clearLines(100)
			isiStatus()

		default:
			clearLines(100)
			cekRegister()

		}

		cekRegister()

	default:
		cekRegister()

	}

}

type arrStr [NMAX]string

func univSearch_String(S, tab string, n int) int {
	var A arrStr
	var i int
	i = 0
	switch tab {
	case "username":
		for i < n {
			A[i] = dataUser.aUser[i].username
			i++
		}
	case "nama":
		for i < n {
			A[i] = dataUser.aUser[i].nama
			i++
		}
	case "password":
		for i < n {
			A[i] = dataUser.aUser[i].password
			i++
		}
	case "status":
		for i < n {
			A[i] = dataUser.aUser[i].status
			i++
		}
	}

	var idx int
	idx = -1
	i = 0

	for idx == -1 && i < n {
		if A[i] == S {
			idx = i
		}
		i++
	}

	return idx
}

// Subprogram Login

func loginAkun() {
	var username, password string

	fmt.Println(border)
	fmt.Println(centerText("Login Akun", width))
	fmt.Println(border)
	fmt.Println("Isi Username dan Password:")
	fmt.Print("Username	: ")
	fmt.Scan(&username)
	fmt.Print("Password	: ")
	fmt.Scan(&password)

	var idx int
	idx = univSearch_String(username, "username", dataUser.n)

	if idx != -1 && dataUser.aUser[idx].password == password {
		if dataUser.aUser[idx].status == "Admin" {
			menuUserAdmin()
		} else {
			menuData()
		}

	} else {
		fmt.Println(border)
		fmt.Println("Terdapat kesalahan pada pengisian username atau password, apakah ingin mengisi ulang? (Y/T)")
		fmt.Println(border)

		var in string
		fmt.Scan(&in)
		switch in {
		case "Y":
			clearLines(100)
			loginAkun()
		case "T":
			clearLines(100)
			menuLoginRegister()
		}
	}
}

// Menu Admin [Mengubah Pengaturan - pengaturan]

func menuUserAdmin() {
	clearLines(100)
	fmt.Println(border)
	fmt.Println(centerText("Menu Perantara Admin", width))
	fmt.Println(border)
	fmt.Println("1. Menu Data")
	fmt.Println("2. Menu Admin")
	fmt.Println("3. Exit")
	fmt.Println()
	fmt.Println("Pilih Menu yang akan diakses")
	fmt.Println(border)
	var in int
	fmt.Scan(&in)

	switch in {
	case 1:
		clearLines(100)
		menuData()
	case 2:
		clearLines(100)
		menuAdmin()
	case 3:
		clearLines(100)
		menuExit()
	default:
		clearLines(100)
		menuUserAdmin()
	}
}

func menuAdmin() {
	fmt.Println(border)
	fmt.Println(centerText("Menu Admin", width))
	fmt.Println(border)
	fmt.Println("1. Ubah metode sorting")
	fmt.Println("2. Ubah akun tenaga kesehatan")
	fmt.Println("3. Ubah indikasi stunting")
	fmt.Println("4. Kembali ke menu sebelumnya")
	fmt.Println("5. Exit")
	fmt.Println()
	fmt.Println("Pilih pengaturan apa yang ingin di ubah:")
	fmt.Println(border)

	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		clearLines(100)
		ubahSorting()
	case 2:
		clearLines(100)
		ubahAkunUser()
	case 3:
		clearLines(100)
		ubahIndikasiStunting()
	case 4:
		clearLines(100)
		menuUserAdmin()
	case 5:
		clearLines(100)
		menuExit()
	default:
		clearLines(100)
		menuAdmin()
	}
}

// Pengaturan Ubah Algoritma Sorting

func ubahSorting() {
	fmt.Println("=== Metode sorting yang digunakan sekarang adalah:", tipeSorting, "===")
	fmt.Println("1. Sorting Selection")
	fmt.Println("2. Sorting Insertion")
	fmt.Println("3. Kembali ke menu sebelumnya")
	fmt.Println()
	fmt.Println("Ubah metode sorting menjadi:")
	fmt.Println(border)

	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		clearLines(100)
		tipeSorting = "Selection"
		ubahSorting()
	case 2:
		clearLines(100)
		tipeSorting = "Insertion"
		ubahSorting()
	case 3:
		clearLines(100)
		menuAdmin()
	default:
		clearLines(100)
		ubahSorting()
	}
}

// Pengaturan Mengubah User

func ubahAkunUser() {
	fmt.Println(border)
	fmt.Println(centerText("Menu Ubah Akun User", width))
	fmt.Println(border)
	fmt.Println("1. Tampilkan list user")
	fmt.Println("2. Kembali ke menu sebelumnya")
	fmt.Println()
	fmt.Println("Pilih opsi:")
	fmt.Println(border)

	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		clearLines(100)
		listUser()
	case 2:
		clearLines(100)
		menuAdmin()
	default:
		clearLines(100)
		ubahAkunUser()
	}
}

func listUser() {
	var i int
	i = 0

	fmt.Println(border)
	fmt.Println("Nama	Username	Password	Status")
	for i < dataUser.n {
		fmt.Printf("%-10s %-10s %-10s %-10s\n",
			dataUser.aUser[i].nama,
			dataUser.aUser[i].username,
			dataUser.aUser[i].password,
			dataUser.aUser[i].status)
		i++
	}
	fmt.Println(border)

	fmt.Println("1. Ubah Data")
	fmt.Println("2. Hilangkan Data")
	fmt.Println("3. Kembali ke menu sebelumnya")
	fmt.Println()
	fmt.Println("Pilih opsi yang diinginkan")
	fmt.Println(border)
	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		clearLines(100)
		adminUbahData()
	case 2:
		clearLines(100)
		adminHilangData()
	case 3:
		clearLines(100)
		menuAdmin()
	default:
		clearLines(100)
		listUser()
	}
}

func adminUbahData() {
	fmt.Println(border)
	fmt.Println("Ketiklah username yang ingin diubah")
	fmt.Println(border)
	var in string
	var idx int
	fmt.Scan(&in)

	idx = univSearch_String(in, "username", dataUser.n)

	if idx != -1 {
		fmt.Println(border)
		fmt.Println(centerText("Menu Ubah User", width))
		fmt.Println(border)
		fmt.Println("1. Nama")
		fmt.Println("2. Username")
		fmt.Println("3. Password")
		fmt.Println("4. Status")
		fmt.Println()
		fmt.Println("Pilih opsi yang diinginkan")
		fmt.Println(border)

		var in2 int
		fmt.Scan(&in2)
		switch in2 {
		case 1:
			fmt.Println(border)
			fmt.Println("Isi nama yang baru:")
			fmt.Scan(&dataUser.aUser[idx].nama)
			fmt.Println(border)
			fmt.Println("Data telah diubah")
		case 2:
			var username string
			var usernameUnik bool
			usernameUnik = false
			for !usernameUnik {
				fmt.Println(border)
				fmt.Println("Isi username yang baru:")
				fmt.Scan(&username)
				usernameUnik = univSearch_String(username, "username", dataUser.n) == -1
				if !usernameUnik {
					fmt.Println(border)
					fmt.Println("Username sudah ada, ganti dengan username yang unik")
				} else {
					dataUser.aUser[idx].username = username
				}
			}
			fmt.Println(border)
			fmt.Println("Data telah diubah")
		case 3:
			fmt.Println(border)
			fmt.Println("Isi password yang baru:")
			fmt.Scan(&dataUser.aUser[idx].password)
			fmt.Println(border)
			fmt.Println("Data telah diubah")
		case 4:
			var in int
			fmt.Println(border)
			fmt.Println("Isi status yang baru:")
			fmt.Println(border)
			fmt.Println("1. Tenaga Kesehatan")
			fmt.Println("2. Admin")
			fmt.Println()
			fmt.Println("Pilih Status")
			fmt.Println(border)
			fmt.Scan(&in)
			for in != 1 && in != 2 {
				fmt.Println(border)
				fmt.Println("Isi status yang baru:")
				fmt.Scan(&in)
			}
			if in == 1 {
				dataUser.aUser[idx].status = "Tenaga Kesesehatan"
			} else {
				dataUser.aUser[idx].status = "Admin"
			}
			fmt.Println(border)
			fmt.Println("Data telah diubah")
		default:
			adminUbahData()
		}

		listUser()

	} else {
		fmt.Println(border)
		fmt.Println(centerText("Menu Data Tidak Ditemukan", width))
		fmt.Println(border)
		fmt.Println("1. Ubah username yang lain")
		fmt.Println("2. kembali kemenu sebelumnya")
		fmt.Println()
		fmt.Println("Data tidak dapat ditemukan, Pilih opsi yang diinginkan")
		fmt.Println(border)
		var in int
		fmt.Scan(&in)
		switch in {
		case 1:
			clearLines(100)
			adminUbahData()
		case 2:
			clearLines(100)
			listUser()
		default:
			adminUbahData()
		}
	}

}

func adminHilangData() {
	fmt.Println(border)
	fmt.Println("Pilih username yang ingin dihilangkan")
	var in string
	var idx int
	fmt.Scan(&in)

	idx = univSearch_String(in, "username", dataUser.n)

	if idx == -1 {
		fmt.Println(border)
		fmt.Println("Username tidak ditemukan")
		fmt.Println("1. Cari ulang username")
		fmt.Println("2. Kembali ke menu sebelumnya")
		fmt.Println()
		fmt.Println("Pilih opsi yang diinginkan")
		fmt.Println(border)

		var in2 int
		fmt.Scan(&in2)
		switch in2 {
		case 1:
			clearLines(100)
			adminHilangData()
		case 2:
			clearLines(100)
			listUser()
		default:
			adminHilangData()
		}
	} else {
		fmt.Println(border)
		fmt.Println("Apakah anda yakin ingin menghilangkan data ini?")
		fmt.Println(dataUser.aUser[idx].nama, dataUser.aUser[idx].username, dataUser.aUser[idx].password, dataUser.aUser[idx].status)
		fmt.Println("Y/T")
		fmt.Println(border)
		var in3 string
		fmt.Scan(&in3)
		switch in3 {
		case "Y":
			clearLines(100)
			geserKiriDataUser(idx)
			fmt.Println("Data telah dihilangkan")
			listUser()
		case "T":
			clearLines(100)
			adminHilangData()
		default:
			clearLines(100)
			adminHilangData()
		}
	}
}

func geserKiriDataUser(idx int) {
	var i int
	i = idx
	for i < dataUser.n {
		dataUser.aUser[i] = dataUser.aUser[i+1]
		i++
	}
	dataUser.n--
}

// Menu Mengubah Indikasi Stunting [Berdasarkan WHO, KEMENKES, Custom]

func ubahIndikasiStunting() {
	fmt.Println(border)
	fmt.Println("Sekarang Aplikasi ini menggunakan indikasi", currentIndicator.nama)
	fmt.Println(border)
	fmt.Println("1. World Health Organization (WHO)")
	fmt.Println("2. Kementrian Kesehatan (KemenKes)")
	fmt.Println("3. Custom")
	fmt.Println("4. Kembali ke menu sebelumnya")
	fmt.Println()
	fmt.Println("Pilih opsi yang diinginkan")
	fmt.Println(border)

	var in int
	fmt.Scan(&in)

	switch in {
	case 1:
		currentIndicator = WHOIndicator
	case 2:
		currentIndicator = KemenKesIndicator
	case 3:
		fmt.Println(border)
		fmt.Println("Isi ulang nilai yang akan menjadi indikasi kondisi stunting:")
		var i int
		i = 0
		for i < maxBulan {
			fmt.Print("Pada Bulan ke-", i, " Indikasi Kondisi Stunted:")
			fmt.Scan(&currentIndicator.A[i].stunted)
			i = i + 1
		}
		i = 0
		for i < maxBulan {
			fmt.Print("Pada Bulan ke-", i, " Indikasi Kondisi Severe:")
			fmt.Scan(&currentIndicator.A[i].severe)
			i = i + 1
		}
		currentIndicator.nama = "Custom"

	case 4:
		clearLines(100)
		menuAdmin()
	default:
		clearLines(100)
		ubahIndikasiStunting()
	}
	ubahIndikasiStunting()
}

// Menu Data Masyarakat

func menuData() {
	fmt.Println(border)
	fmt.Println(centerText("Menu Data", width))
	fmt.Println(border)
	fmt.Println("1. Lihat Data")
	fmt.Println("2. Input Data")
	fmt.Println("3. Exit")
	fmt.Println()
	fmt.Println("Pilih Menu yang ingin di akses")
	fmt.Println(border)

	var in int
	fmt.Scan(&in)

	switch in {
	case 1:
		clearLines(100)
		menuLihatData()
	case 2:
		clearLines(100)
		opsiDaftarDataBaru()
	case 3:
		clearLines(100)
		menuExit()
	default:
		clearLines(100)
		menuData()
	}

	menuData()
}

// Menu Lihat Data Masyarakat

func menuLihatData() {
	var i int
	i = 0
	fmt.Println(border)
	fmt.Printf("%-6s %-8s %-7s %-7s %-5s %-7s %-10s\n", "Nama", "Asal", "Gender", "Status", "Umur", "Tinggi", "Tanggal")
	fmt.Println("-----------------------------------------------------------")
	for i < dataMain.n {
		fmt.Printf("%-6s %-8s %-7s %-7s %-5d %-7.1f %-10s\n", dataMain.aData[i].nama, dataMain.aData[i].asal, dataMain.aData[i].gender, dataMain.aData[i].status, dataMain.aData[i].umur, dataMain.aData[i].tinggi, dataMain.aData[i].tanggal)
		i = i + 1
	}

	fmt.Println(border)
	fmt.Println("Pilih Opsi:")
	fmt.Println("1. Cari Data")
	fmt.Println("2. Urutkan Data")
	fmt.Println("3. Kembali Ke Menu Sebelumnya")
	fmt.Println("4. Keluar dari Aplikasi")
	fmt.Println()
	fmt.Println("Pilih opsi yang diinginkan")
	fmt.Println(border)

	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		clearLines(100)
		cariDataMasyarakat()
	case 2:
		clearLines(100)
		menuUrutDataMasyarakat()
	case 3:
		clearLines(100)
		menuData()
	case 4:
		clearLines(100)
		menuExit()
	}
}

func cariDataMasyarakat() {
	fmt.Println(border)
	fmt.Print("Isi Nama yang akan dicari: ")
	var nama string
	fmt.Scan(&nama)
	var idx int
	if dataMain.sort {
		idx = univSearch_DM_Binary(nama)
	} else {
		idx = univSearch_DM_Sequential(nama)
	}

	if idx != -1 {
		dataMainKetemu(idx)
	} else {
		fmt.Println(border)
		fmt.Println("Nama tidak ditemukan")
		ulangSearchDataMain()
	}
}

func dataMainKetemu(i int) {
	fmt.Println(border)
	fmt.Println("Data Ditemukan:")
	fmt.Println("Nama  Asal     Gender  Status  Umur  Tinggi  Tanggal")
	fmt.Printf("%-6s %-8s %-7s %-7s %-5d %-7.1f %-10s\n", dataMain.aData[i].nama, dataMain.aData[i].asal, dataMain.aData[i].gender, dataMain.aData[i].status, dataMain.aData[i].umur, dataMain.aData[i].tinggi, dataMain.aData[i].tanggal)
	fmt.Println(border)
	fmt.Println("1. Cari Data Lain")
	fmt.Println("2. Kembali ke Menu Data")
	fmt.Println("3. Tambah Data Baru")

	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		clearLines(100)
		ulangSearchDataMain()
	case 2:
		clearLines(100)
		menuData()
	case 3:
		clearLines(100)
		opsiDaftarDataBaru()
	default:
		clearLines(100)
		dataMainKetemu(i)
	}
}

func ulangSearchDataMain() {
	fmt.Println()
	fmt.Println("Apakah Ingin mengulangi pencarian? (Y/T)")
	fmt.Println(border)
	var in string
	fmt.Scan(&in)
	switch in {
	case "Y":
		clearLines(100)
		cariDataMasyarakat()
	case "T":
		clearLines(100)
		menuLihatData()
	default:
		clearLines(100)
		ulangSearchDataMain()
	}
}

func univSearch_DM_Binary(n string) int {
	var l, m, r int
	l = 0
	r = dataMain.n - 1
	m = (l + r) / 2
	for n != dataMain.aData[m].nama && l <= r {
		if dataMain.aData[m].nama < n {
			l = m + 1
		} else {
			r = m - 1
		}
		m = (l + r) / 2
	}
	return m
}

func univSearch_DM_Sequential(n string) int {
	var i, idx int
	idx = -1
	i = 0
	for i < dataMain.n {
		if n == dataMain.aData[i].nama {
			idx = i
		}
		i = i + 1
	}
	return idx
}

func menuUrutDataMasyarakat() {

	fmt.Println(border)
	fmt.Println(centerText("Menu Pengurutan", width))
	fmt.Println(border)
	fmt.Println("1. Urutkan berdasarkan Nama")
	fmt.Println("2. Urutkan berdasarkan Asal")
	fmt.Println("3. Urutkan berdasarkan Gender")
	fmt.Println("4. Urutkan berdasarkan Status")
	fmt.Println("5. Urutkan berdasarkan Umur")
	fmt.Println("6. Urutkan berdasarkan Tinggi")
	fmt.Println("7. Kembali ke menu sebelumnya")
	fmt.Println("8. Keluar")
	fmt.Println()
	fmt.Println("Pilih opsi yang diinginkan")
	fmt.Println(border)

	var in int
	fmt.Scan(&in)

	switch in {
	case 7:
		clearLines(100)
		menuLihatData()
	case 8:
		clearLines(100)
		menuExit()
	}

	fmt.Println(border)
	fmt.Println("Diurutkan Secara?")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Println()
	fmt.Println("Pilih opsi yang diinginkan")
	fmt.Println(border)

	var in2 int
	fmt.Scan(&in2)

	var sortOrder string
	if in2 == 1 {
		sortOrder = "A"
	} else {
		sortOrder = "D"
	}

	var ref string

	switch in {
	case 1:
		ref = "nama"
	case 2:
		ref = "asal"
	case 3:
		ref = "gender"
	case 4:
		ref = "status"
	case 5:
		ref = "umur"
	case 6:
		ref = "tinggi"
	default:
		menuUrutDataMasyarakat()
	}

	prosesPilihAlgoritmaPengurutan(ref, sortOrder)

	if sortOrder == "A" {
		reverseArray()
	}

	menuTelahDiurutkan(sortOrder, ref)
}

func menuTelahDiurutkan(Order, by string) {
	var ord string

	if Order == "A" {
		ord = "Ascending"
	} else {
		ord = "Descending"
	}

	fmt.Println(border)
	fmt.Printf("Data telah diurutkan berdasarkan %s dengan Order %s, silahkan pilih opsi:\n", by, ord)

	dataMain.sort = true
	dataMain.sortby = by
	dataMain.Order = Order

	menuLihatData()
}

func reverseArray() {
	var i, j int
	var temp structDataMain
	i = 0
	for i < dataMain.n/2 {
		j = dataMain.n - 1 - i
		temp = dataMain.aData[i]
		dataMain.aData[i] = dataMain.aData[j]
		dataMain.aData[j] = temp
		i = i + 1
	}
}

func prosesPilihAlgoritmaPengurutan(by, Order string) {
	if tipeSorting == "Selection" {
		if by == "nama" || by == "asal" || by == "gender" || by == "status" {
			sortSelectionString(by, Order)
		} else if by == "umur" {
			sortSelectionInt(Order)
		} else {
			sortSelectionFloat(Order)
		}
	} else {
		if by == "nama" || by == "asal" || by == "gender" || by == "status" {
			sortInsertionString(by, Order)
		} else if by == "umur" {
			sortInsertionInt(Order)
		} else {
			sortInsertionFloat(Order)
		}
	}
}

func sortSelectionString(by, Order string) {
	var i int
	i = 0
	if by == "nama" {
		for i < dataMain.n {
			dataMain.aData[i].stringTemp = dataMain.aData[i].nama
			i = i + 1
		}
	} else if by == "asal" {
		for i < dataMain.n {
			dataMain.aData[i].stringTemp = dataMain.aData[i].asal
			i = i + 1
		}
	} else if by == "gender" {
		for i < dataMain.n {
			dataMain.aData[i].stringTemp = dataMain.aData[i].gender
			i = i + 1
		}
	} else {
		for i < dataMain.n {
			dataMain.aData[i].stringTemp = dataMain.aData[i].status
			i = i + 1
		}
	}

	var temp structDataMain
	var j, imax int

	i = 0
	for i < dataMain.n-1 {
		j = i + 1
		imax = i
		for j < dataMain.n {
			if dataMain.aData[j].stringTemp > dataMain.aData[imax].stringTemp {
				imax = j
			}
			j = j + 1
		}
		temp = dataMain.aData[i]
		dataMain.aData[i] = dataMain.aData[imax]
		dataMain.aData[imax] = temp
		i = i + 1
	}
}

func sortSelectionInt(Order string) {
	var i int
	i = 0

	for i < dataMain.n {
		dataMain.aData[i].intTemp = dataMain.aData[i].umur
		i = i + 1
	}

	var temp structDataMain
	var j, imax int

	i = 0
	for i < dataMain.n-1 {
		j = i + 1
		imax = i
		for j < dataMain.n {
			if dataMain.aData[j].intTemp > dataMain.aData[imax].intTemp {
				imax = j
			}
			j = j + 1
		}
		temp = dataMain.aData[i]
		dataMain.aData[i] = dataMain.aData[imax]
		dataMain.aData[imax] = temp
		i = i + 1
	}
}

func sortSelectionFloat(Order string) {
	var i int
	i = 0

	for i < dataMain.n {
		dataMain.aData[i].floatTemp = dataMain.aData[i].tinggi
		i = i + 1
	}

	var temp structDataMain
	var j, imax int

	i = 0
	for i < dataMain.n-1 {
		j = i + 1
		imax = i
		for j < dataMain.n {
			if dataMain.aData[j].floatTemp > dataMain.aData[imax].floatTemp {
				imax = j
			}
			j = j + 1
		}
		temp = dataMain.aData[i]
		dataMain.aData[i] = dataMain.aData[imax]
		dataMain.aData[imax] = temp
		i = i + 1
	}
}

func sortInsertionString(by, Order string) {
	var i int
	i = 0
	if by == "nama" {
		for i < dataMain.n {
			dataMain.aData[i].stringTemp = dataMain.aData[i].nama
			i = i + 1
		}
	} else if by == "asal" {
		for i < dataMain.n {
			dataMain.aData[i].stringTemp = dataMain.aData[i].asal
			i = i + 1
		}
	} else if by == "gender" {
		for i < dataMain.n {
			dataMain.aData[i].stringTemp = dataMain.aData[i].gender
			i = i + 1
		}
	} else {
		for i < dataMain.n {
			dataMain.aData[i].stringTemp = dataMain.aData[i].status
			i = i + 1
		}
	}

	var j int
	var temp string
	var tempDat structDataMain

	i = 1

	for i < dataMain.n {
		tempDat = dataMain.aData[i]
		temp = dataMain.aData[i].stringTemp
		j = i - 1
		for j >= 0 && dataMain.aData[j].stringTemp > temp {
			dataMain.aData[j+1] = dataMain.aData[j]
			j = j - 1
		}
		dataMain.aData[j+1] = tempDat
		i = i + 1
	}
}

func sortInsertionInt(Order string) {
	var i int
	i = 0

	for i < dataMain.n {
		dataMain.aData[i].intTemp = dataMain.aData[i].umur
		i = i + 1
	}

	var j int
	var temp int
	var tempDat structDataMain

	i = 1

	for i < dataMain.n {
		tempDat = dataMain.aData[i]
		temp = dataMain.aData[i].intTemp
		j = i - 1
		for j >= 0 && dataMain.aData[j].intTemp > temp {
			dataMain.aData[j+1] = dataMain.aData[j]
			j = j - 1
		}
		dataMain.aData[j+1] = tempDat
		i = i + 1
	}
}

func sortInsertionFloat(Order string) {
	var i int
	i = 0

	for i < dataMain.n {
		dataMain.aData[i].floatTemp = dataMain.aData[i].tinggi
		i = i + 1
	}

	var j int
	var temp float64
	var tempDat structDataMain

	i = 1

	for i < dataMain.n {
		tempDat = dataMain.aData[i]
		temp = dataMain.aData[i].floatTemp
		j = i - 1
		for j >= 0 && dataMain.aData[j].floatTemp > temp {
			dataMain.aData[j+1] = dataMain.aData[j]
			j = j - 1
		}
		dataMain.aData[j+1] = tempDat
		i = i + 1
	}
}

// Menu Input Data Masyarakat

func opsiDaftarDataBaru() {
	fmt.Println(border)
	fmt.Println("Isi seluruh spesifikasi anak berikut:")
	isiNamaData()
	isiUmurData()
	isiAsalData()
	isiGenderData()
	isiTinggiData()
	dataMain.aData[dataMain.n].tanggal = time.Now().Format("02-01-2006")
	dataMain.aData[dataMain.n].status = pengukuranIndikasiStunting(dataMain.aData[dataMain.n].umur, dataMain.aData[dataMain.n].tinggi)
	dataMain.sort = false

	fmt.Println()
	fmt.Println(border)
	fmt.Println("Data baru telah ditambahkan")

	dataBaruPeubah()

	menuData()
}

func dataBaruPeubah() {
	tampilDataBaru()
	opsiUbahDataBaru()
}

func isiNamaData() {
	var nama string
	fmt.Print("Nama Lengkap	: ")
	fmt.Scan(&nama)
	var idx int
	idx = searchNamaMasyarakat(nama)
	duplikatNamaDataMasyatakat(idx, nama)
	dataMain.aData[dataMain.n].nama = nama
}

func searchNamaMasyarakat(in string) int {
	var i, idx int
	idx = -1
	i = 0
	for i < dataMain.n {
		if dataMain.aData[i].nama == in {
			idx = i
		}
		i = i + 1
	}
	return idx
}

func duplikatNamaDataMasyatakat(idx int, nama string) {
	if idx != -1 {
		fmt.Println(border)
		fmt.Println("Nama", nama, "sudah ada dalam data, tolong berikan nama lain, apa ingin mengisi ulang nama? (Y/T)")
		var in string
		fmt.Scan(&in)
		switch in {
		case "Y":
			clearLines(100)
			isiNamaData()
		case "T":
			clearLines(100)
			menuData()
		default:
			clearLines(100)
			duplikatNamaDataMasyatakat(idx, nama)
		}
	}
}

func isiUmurData() {
	var umur int
	fmt.Print("Umur (Bulan)	: ")
	fmt.Scan(&umur)
	umurTidakValid(umur)
	dataMain.aData[dataMain.n].umur = umur
}

func umurTidakValid(umur int) {
	if umur <= 0 {
		fmt.Println(border)
		fmt.Println("Umur tidak valid, apakah ingin mengisi ulang umur? (Y/T)")
		var in string
		fmt.Scan(&in)
		switch in {
		case "Y":
			clearLines(100)
			isiUmurData()
		case "T":
			clearLines(100)
			menuData()
		default:
			clearLines(100)
			umurTidakValid(umur)
		}
	}
}

func isiAsalData() {
	var asal string
	fmt.Print("Asal Daerah	: ")
	fmt.Scan(&asal)
	dataMain.aData[dataMain.n].asal = asal
}

func isiGenderData() {
	var gender string
	fmt.Print("Gender (L/P)	: ")
	fmt.Scan(&gender)
	if gender == "L" || gender == "P" {
		dataMain.aData[dataMain.n].gender = gender
	} else {
		isiGenderData()
	}
}

func isiTinggiData() {
	var tinggi float64
	fmt.Print("Tinggi (cm)	: ")
	fmt.Scan(&tinggi)
	tinggiTidakValid(tinggi)
	dataMain.aData[dataMain.n].tinggi = tinggi
}

func tinggiTidakValid(tinggi float64) {
	if tinggi <= 0 {
		fmt.Println(border)
		fmt.Println("tinggi", tinggi, "tidak valid, apakah ingin mengisi ulang tinggi? (Y/T)")
		var in string
		fmt.Scan(&in)
		switch in {
		case "Y":
			clearLines(100)
			isiTinggiData()
		case "T":
			clearLines(100)
			opsiDaftarDataBaru()
		default:
			clearLines(100)
			tinggiTidakValid(tinggi)
		}
	}
}

func tampilDataBaru() {
	fmt.Println(border)
	fmt.Println("Data yang akan diinput:")
	fmt.Printf("%-6s %-8s %-7s %-7s %-5s %-7s %-10s\n", "Nama", "Asal", "Gender", "Status", "Umur", "Tinggi", "Tanggal")
	fmt.Printf("%-6s %-8s %-7s %-7s %-5d %-7.1f %-10s\n",
		dataMain.aData[dataMain.n].nama,
		dataMain.aData[dataMain.n].asal,
		dataMain.aData[dataMain.n].gender,
		dataMain.aData[dataMain.n].status,
		dataMain.aData[dataMain.n].umur,
		dataMain.aData[dataMain.n].tinggi,
		dataMain.aData[dataMain.n].tanggal,
	)
}

func opsiUbahDataBaru() {
	fmt.Println(border)
	fmt.Println("Apakah ada bagian yang ingin diubah?")
	fmt.Println(border)
	fmt.Println("1. Nama")
	fmt.Println("2. Asal")
	fmt.Println("3. Gender")
	fmt.Println("4. Tinggi")
	fmt.Println("5. Tidak ada")
	fmt.Println()
	fmt.Println("Pilih opsi yang diinginkan")
	fmt.Println(border)
	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		clearLines(100)
		isiNamaData()
	case 2:
		clearLines(100)
		isiAsalData()
	case 3:
		clearLines(100)
		isiGenderData()
	case 4:
		clearLines(100)
		isiTinggiData()
	case 5:
		clearLines(100)
		dataMain.n++
		menuData()
	default:
		clearLines(100)
		opsiUbahDataBaru()
	}
	clearLines(100)
	dataBaruPeubah()
}

func dataSearch(Str, Type string, n int) int {
	var i int
	var A [NMAX]string
	i = 0
	switch Type {
	case "nama":
		for i < n {
			A[i] = dataMain.aData[i].nama
			i++
		}
	case "asal":
		for i < n {
			A[i] = dataMain.aData[i].asal
			i++
		}
	case "gender":
		for i < n {
			A[i] = dataMain.aData[i].gender
			i++
		}
	case "status":
		for i < n {
			A[i] = dataMain.aData[i].status
			i++
		}
	}
	i = 0
	return i
}

func pengukuranIndikasiStunting(umur int, tinggi float64) string {
	if tinggi > currentIndicator.A[umur].stunted {
		return "Normal"
	} else if tinggi > currentIndicator.A[umur].severe {
		return "Stunted"
	} else if tinggi > 0 {
		return "Severe"
	} else {
		return "Tidak Valid"
	}
}
