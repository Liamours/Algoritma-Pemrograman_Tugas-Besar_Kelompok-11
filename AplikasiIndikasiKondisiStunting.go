package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	menuLoginRegister()
}

const NMAX int = 2028

type structUser struct {
	nama, username, password, status string
}

type arrayUser [NMAX]structUser

type structDataUser struct {
	aUser arrayUser
	n     int
}

var dataUser structDataUser

// Subprogram yang akan dijalankan ketika ingin menutup aplikasi

func menuExit() {

	fmt.Println("Apakah anda yakin akan meninggalkan aplikasi?")
	fmt.Println("1. Iya")
	fmt.Println("2. Cancle")
	fmt.Println()
	fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")

	var in int

	fmt.Scan(&in)
	switch in {
	case 1:
		os.Exit(0)
	case 2:
		menuLoginRegister()
	default:
		menuExit()
	}
}

// Menu untuk Login atau Registrasi

func menuLoginRegister() {

	fmt.Println("Aplikasi Pendataan Kondisi Stunting")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Exit")
	fmt.Println()
	fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")

	var in int
	fmt.Scan(&in)

	switch in {
	case 1:
		menuLogin()

	case 2:
		menuRegister()

	case 3:
		menuExit()

	default:
		menuLoginRegister()

	}
}

// Subprogram yang akan dijalankan saat ingin melakukan Registrasi

func menuRegister() {

	fmt.Println("Menu Register")
	fmt.Println("1. Daftar Akun Baru")
	fmt.Println("2. Kembali")
	fmt.Println()
	fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")

	var in int
	fmt.Scan(&in)

	switch in {
	case 1:
		daftarAkun()

	case 2:
		menuLoginRegister()

	default:
		menuRegister()
	}

}

func daftarAkun() {

	if dataUser.n == NMAX {
		fmt.Println("Data User telah penuh")
		fmt.Println("1. Kembali ke menu utama")
		fmt.Println("2. Keluar dari aplikasi")
		fmt.Println()
		fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")

		var in int
		fmt.Scan(&in)

		switch in {
		case 1:
			menuLoginRegister()

		case 2:
			menuExit()

		default:
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

	fmt.Print("Isi Nama Lengkap: ")
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
		fmt.Print("Username: ")
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

	fmt.Print("Password: ")
	fmt.Scan(&dataUser.aUser[dataUser.n].password)

}

func isiStatus() {
	// IS: Dijalankan ketika ingin mengisi status pada suatu data User
	// FS: status pada data user terisi
	/*
		Keterangan: Status user hanya bisa terisi antara Tenaga Kesehatan (TenKes) atau Admin (Adm)
	*/

	fmt.Println("Tentukan Status:")
	fmt.Println("1. Tenaga Kesehatan")
	fmt.Println("2. Admin")
	fmt.Println()
	fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")
	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		dataUser.aUser[dataUser.n].status = "TenKes"

	case 2:
		dataUser.aUser[dataUser.n].status = "Adm"

	default:
		isiStatus()
	}

}

func cekRegister() {
	// IS: Terdapat data suatu user
	// FS: Mengubah data berdasarkan input yang diminta user
	/*
		Keterangan: User dapat mengubah nama, username, password, atau status
	*/

	fmt.Println("Nama		:", dataUser.aUser[dataUser.n].nama)
	fmt.Println("Username	:", dataUser.aUser[dataUser.n].username)
	fmt.Println("Password	:", dataUser.aUser[dataUser.n].password)
	fmt.Println("Status		:", dataUser.aUser[dataUser.n].status)
	fmt.Println("Apakah akun sudah sesuai? (Y/T)")
	fmt.Println()
	fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")

	var inS string
	fmt.Scan(&inS)
	switch inS {
	case "Y":
		dataUser.n++
		menuLoginRegister()
	case "T":
		fmt.Println("Data apa yang ingin diubah?")
		fmt.Println("1. Nama")
		fmt.Println("2. Username")
		fmt.Println("3. Password")
		fmt.Println("4. Status")
		fmt.Println()
		fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")

		var in int
		fmt.Scan(&in)
		switch in {
		case 1:
			isiNama()

		case 2:
			isiUsername()

		case 3:
			isiPassword()

		case 4:
			isiStatus()

		default:
			cekRegister()

		}

		cekRegister()

	default:
		cekRegister()

	}

}

// Function untuk melakukan searching pada array string

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

// Subprogram yang akan dijalankan saat ingin melakukan Login

func menuLogin() {

	fmt.Println("Menu Login")
	fmt.Println("1. Login Akun")
	fmt.Println("2. Kembali Ke Menu Utama")
	fmt.Println()
	fmt.Println("Ketik opsi yang diinginkan, lalu tekan Enter")

	var in int
	fmt.Scan(&in)

	switch in {
	case 1:
		loginAkun()

	case 2:
		menuLoginRegister()

	default:
		menuLogin()

	}

}

func loginAkun() {
	var username, password string

	fmt.Println("Login Akun")
	fmt.Println("Isi Username dan Password:")
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	var idx int
	idx = univSearch_String(username, "username", dataUser.n)

	if idx != -1 && dataUser.aUser[idx].password == password {
		if dataUser.aUser[idx].status == "Adm" {
			menuUserAdmin()
		} else {
			menuData()
		}

	} else {
		fmt.Println("Terdapat kesalahan pada pengisian username atau password, apakah ingin mengisi ulang? (Y/T)")

		var in string
		fmt.Scan(&in)
		switch in {
		case "Y":
			loginAkun()
		case "T":
			menuLoginRegister()
		}
	}
}

func menuUserAdmin() {
	fmt.Println("Pilih menu:")
	fmt.Println("1. Menu Data")
	fmt.Println("2. Menu Admin")
	fmt.Println("3. Exit")
	var in int
	fmt.Scan(&in)

	switch in {
	case 1:
		menuData()
	case 2:
		menuAdmin()
	case 3:
		menuExit()
	default:
		menuUserAdmin()
	}
}

func menuAdmin() {
	fmt.Println("Pilih pengaturan apa yang ingin di ubah:")
	fmt.Println("1. Ubah metode sorting")
	fmt.Println("2. Ubah akun tenaga kesehatan")
	fmt.Println("3. Ubah indikasi stunting")
	fmt.Println("4. Kembali ke menu sebelumnya")
	fmt.Println("5. Exit")

	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		ubahSorting()
	case 2:
		ubahAkunUser()
	case 3:
		ubahIndikasiStunting()
	case 4:
		menuUserAdmin()
	case 5:
		menuExit()
	default:
		menuAdmin()
	}
}

var tipeSorting string = "Selection"

func ubahSorting() {
	fmt.Println("Metode sorting yang digunakan sekarang adalah:", tipeSorting)
	fmt.Println("Ubah metode sorting menjadi:")
	fmt.Println("1. Sorting Selection")
	fmt.Println("2. Sorting Insertion")
	fmt.Println("3. Kembali ke menu sebelumnya")

	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		tipeSorting = "Selection"
		ubahSorting()
	case 2:
		tipeSorting = "Insertion"
		ubahSorting()
	case 3:
		menuAdmin()
	default:
		ubahSorting()
	}
}

func ubahAkunUser() {
	fmt.Println("Pilih opsi:")
	fmt.Println("1. Tampilkan list user")
	fmt.Println("2. Kembali ke menu sebelumnya")

	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		listUser()
	case 2:
		menuAdmin()
	default:
		ubahAkunUser()
	}
}

func listUser() {
	var i int
	i = 0
	fmt.Println("Nama, Username, Password, Status")
	for i < dataUser.n {
		fmt.Println(dataUser.aUser[i].nama, dataUser.aUser[i].username, dataUser.aUser[i].password, dataUser.aUser[i].status)
		i++
	}

	fmt.Println("1. Ubah Data")
	fmt.Println("2. Hilangkan Data")
	fmt.Println("3. Kembali ke menu sebelumnya")
	var in int
	fmt.Scan(&in)
	switch in {
	case 1:
		adminUbahData()
	case 2:
		adminHilangData()
	case 3:
		ubahAkunUser()
	default:
		listUser()
	}
}

func adminUbahData() {
	fmt.Println("Ketiklah username yang ingin diubah")
	var in string
	var idx int
	fmt.Scan(&in)

	idx = univSearch_String(in, "username", dataUser.n)

	if idx != -1 {
		fmt.Println("Bagian apa yang indin diubah:")
		fmt.Println("1. Nama")
		fmt.Println("2. Username")
		fmt.Println("3. Password")
		fmt.Println("4. Status")

		var in2 int
		fmt.Scan(&in2)
		switch in2 {
		case 1:
			fmt.Println("Isi nama yang baru:")
			fmt.Scan(&dataUser.aUser[idx].nama)
			fmt.Println("Data telah diubah")
		case 2:
			var username string
			var usernameUnik bool
			usernameUnik = false
			for !usernameUnik {
				fmt.Println("Isi username yang baru:")
				fmt.Scan(&username)
				usernameUnik = univSearch_String(username, "username", dataUser.n) == -1
				if !usernameUnik {
					fmt.Println("Username sudah ada, ganti dengan username yang unik")
				} else {
					dataUser.aUser[idx].username = username
				}
			}
			fmt.Println("Data telah diubah")
		case 3:
			fmt.Println("Isi password yang baru:")
			fmt.Scan(&dataUser.aUser[idx].password)
			fmt.Println("Data telah diubah")
		case 4:
			var in int
			fmt.Println("Isi status yang baru:")
			fmt.Println("1. Tenaga Kesehatan")
			fmt.Println("2. Admin")
			fmt.Scan(&in)
			for in != 1 || in != 2 {
				fmt.Println("Isi status yang baru:")
				fmt.Scan(&in)
			}
			if in == 1 {
				dataUser.aUser[idx].status = "TenKes"
			} else {
				dataUser.aUser[idx].status = "Adm"
			}
			fmt.Println("Data telah diubah")
		default:
			adminUbahData()
		}

		listUser()

	} else {
		fmt.Println("Data tidak ditemukan")
		fmt.Println("1. Ubah username yang lain")
		fmt.Println("2. kembali kemenu sebelumnya")
		var in int
		fmt.Scan(&in)
		switch in {
		case 1:
			adminUbahData()
		case 2:
			listUser()
		default:
			adminUbahData()
		}
	}

}

func adminHilangData() {
	fmt.Println("Pilih username yang ingin dihilangkan")
	var in string
	var idx int
	fmt.Scan(&in)

	idx = univSearch_String(in, "username", dataUser.n)

	if idx == -1 {
		fmt.Println("Username tidak ditemukan")
		fmt.Println("1. Cari ulang username")
		fmt.Println("2. Kembali ke menu sebelumnya")

		var in2 int
		fmt.Scan(&in2)
		switch in2 {
		case 1:
			adminHilangData()
		case 2:
			listUser()
		default:
			adminHilangData()
		}
	} else {
		fmt.Println("Apakah anda yakin ingin menghilangkan data ini?")
		fmt.Println(dataUser.aUser[idx].nama, dataUser.aUser[idx].username, dataUser.aUser[idx].password, dataUser.aUser[idx].status)
		fmt.Printf("Y/T")
		var in3 string
		fmt.Scan(&in3)
		switch in3 {
		case "Y":
			geserKiriDataUser(idx)
			fmt.Println("Data telah dihilangkan")
			listUser()
		case "T":
			adminHilangData()
		default:
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

type indikasiStunting struct {
	severe, stunted, normal int
	nama                    string
}

var WHOIndicator indikasiStunting = indikasiStunting{1, 2, 3, "WHO"}
var currentIndicator indikasiStunting = indikasiStunting{1, 2, 3, "WHO"}
var KemenKesIndicator indikasiStunting = indikasiStunting{1, 2, 3, "KemenKes"}

func ubahIndikasiStunting() {
	fmt.Println("Sekarang Aplikasi ini menggunakan indikasi", currentIndicator.nama)
	fmt.Println("-3SD:", currentIndicator.severe)
	fmt.Println("-3SD sd <-2SD:", currentIndicator.stunted)
	fmt.Println("-2SD sd <3SD:", currentIndicator.normal)
	fmt.Println()
	fmt.Println("Aplikasi ini akan menentukan kondisi stunting (Normal, Stunted, Severe) berdasarkan:")
	fmt.Println("1. World Health Organization (WHO)")
	fmt.Println("2. Kementrian Kesehatan (KemenKes)")
	fmt.Println("3. Custom")
	fmt.Println("4. Kembali ke menu sebelumnya")

	var in int
	fmt.Scan(&in)

	switch in {
	case 1:
		currentIndicator = WHOIndicator
	case 2:
		currentIndicator = KemenKesIndicator
	case 3:
		fmt.Println("Isi ulang nilai yang akan menjadi indikasi kondisi stunting:")
		fmt.Print("Indikasi Kondisi Normal:")
		fmt.Scan(&currentIndicator.normal)
		fmt.Print("Indikasi Kondisi Stunted:")
		fmt.Scan(&currentIndicator.stunted)
		fmt.Print("Indikasi Kondisi Severe:")
		fmt.Scan(&currentIndicator.severe)
	case 4:
		menuAdmin()
	default:
		ubahIndikasiStunting()
	}
	ubahIndikasiStunting()
}

func menuData() {
	fmt.Println("1. Lihat Data")
	fmt.Println("2. Input Data")
	fmt.Println("3. Exit")

	var in int
	fmt.Scan(&in)

	switch in {
	case 1:
		menuLihatData()
	case 2:
		menuInputData()
	case 3:
		menuExit()
	default:
		menuData()
	}

	menuData()
}

func menuLihatData() {

}

func menuInputData() {
	fmt.Println("1. Daftar Data Baru")
	fmt.Println("2. Kembali ke menu sebelumnya")

	var in int
	fmt.Scan(&in)

	switch in {
	case 1:
		opsiDaftarDataBaru()
	case 2:
		menuData()
	default:
		menuInputData()
	}
	menuInputData()
}

type structDataMain struct {
	nama, asal, gender, status string
	tinggi                     float64
	tanggal                    string
}

type arrDataMain [NMAX]structDataMain

type structMainArray struct {
	aData arrDataMain
	n     int
	sort  bool
}

var dataMain structMainArray

func opsiDaftarDataBaru() {
	fmt.Println("Isi seluruh spesifikasi anak berikut:")
	isiNamaData()
	isiUmurData()
	isiAsalData()
	isiGenderData()
	isiTinggiData()
	dataMain.aData[dataMain.n].tanggal = time.Now().Format("02-01-2006")

	dataMain.n++

	fmt.Println("Data baru telah ditambahkan")

	tampilDataBaru()
	
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

func isiNamaData() {
	var nama string
	fmt.Print("Nama: ")
	fmt.Scan(&nama)
}

func isiUmurData() {
	var umur int
	fmt.Print("Umur: ")
	fmt.Scan(&umur)
}

func isiAsalData() {
	var asal string
	fmt.Print("Asal Daerah: ")
	fmt.Scan(&asal)
}

func isiGenderData() {
	var gender string
	fmt.Print("Gender: ")
	fmt.Scan(&gender)
}

func isiTinggiData() {
	var tinggi float64
	fmt.Print("Tinggi Badan: ")
	fmt.Scan(&tinggi)
}

func insertionSort() {

}

func selectionSort() {

}
