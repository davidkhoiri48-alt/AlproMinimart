package main

import "fmt"

const MAX = 100

type Barang struct {
	id    int
	nama  string
	harga int
}

type Transaksi struct {
	idBarang int
	jumlah   int
	total    int
}

var dataBarang [MAX]Barang
var dataTransaksi [MAX]Transaksi

var jumlahBarang int = 0
var jumlahTransaksi int = 0

func awal() bool {
	var pilih int = -1

	for pilih != 1 && pilih != 2 {
		fmt.Println("====================================")
		fmt.Println("=== SELAMAT DATANG DI DAVID MART ===")
		fmt.Println("=== 	IT-49-03   ===")
		fmt.Println("=== 103032500040 & 1030325000 ===")
		fmt.Println("====================================")
		fmt.Println("1. Lanjut")
		fmt.Println("2. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih != 1 && pilih != 2 {
			fmt.Println()
			fmt.Println("Tolong input angka 1 dan 2 saja!")
			fmt.Println()
		}
	}

	if pilih == 1 {
		return true
	} else {
		fmt.Println("Program selesai.")
		return false
	}
}

func tambahBarang() {
	if jumlahBarang >= MAX {
		fmt.Println("Data barang sudah penuh!")
	} else {
		var b Barang
		fmt.Print("ID: ")
		fmt.Scan(&b.id)
		fmt.Print("Nama: ")
		fmt.Scan(&b.nama)
		fmt.Print("Harga: ")
		fmt.Scan(&b.harga)

		dataBarang[jumlahBarang] = b
		jumlahBarang++
		fmt.Println("Barang berhasil ditambahkan")
	}
}

func sequentialSearch(id int) int {
	i := 0
	idx := -1

	for i < jumlahBarang {
		if dataBarang[i].id == id {
			idx = i
		}
		i++
	}
	return idx
}

func binarySearch(id int) int {
	low := 0
	high := jumlahBarang - 1
	idx := -1

	for low <= high {
		mid := (low + high) / 2

		if dataBarang[mid].id == id {
			idx = mid
			low = high + 1
		} else if dataBarang[mid].id < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}

func selectionSort() {
	i := 0
	for i < jumlahBarang-1 {
		min := i
		j := i + 1

		for j < jumlahBarang {
			if dataBarang[j].id < dataBarang[min].id {
				min = j
			}
			j++
		}

		temp := dataBarang[i]
		dataBarang[i] = dataBarang[min]
		dataBarang[min] = temp

		i++
	}
	fmt.Println("Data berhasil diurutkan ascending")
}

func insertionSortDesc() {
	i := 1
	for i < jumlahBarang {
		key := dataBarang[i]
		j := i - 1

		for j >= 0 && dataBarang[j].id < key.id {
			dataBarang[j+1] = dataBarang[j]
			j--
		}
		dataBarang[j+1] = key

		i++
	}
	fmt.Println("Data berhasil diurutkan descending")
}

func tambahTransaksi() {
	if jumlahTransaksi >= MAX {
		fmt.Println("Data transaksi sudah penuh!")
	} else {
		var id, jumlah int
		fmt.Print("ID Barang: ")
		fmt.Scan(&id)

		idx := sequentialSearch(id)

		if idx == -1 {
			fmt.Println("ID barang tidak ditemukan! Silakan input ID yang benar.")
		} else {
			fmt.Print("Jumlah: ")
			fmt.Scan(&jumlah)

			if jumlah <= 0 {
				fmt.Println("Jumlah tidak valid!")
			} else {
				total := jumlah * dataBarang[idx].harga

				dataTransaksi[jumlahTransaksi] = Transaksi{
					idBarang: id,
					jumlah:   jumlah,
					total:    total,
				}

				jumlahTransaksi++
				fmt.Println("Transaksi berhasil!")
			}
		}
	}
}

func tampilBarang() {
	if jumlahBarang == 0 {
		fmt.Println("Belum ada data barang")
	} else {
		i := 0
		for i < jumlahBarang {
			fmt.Println(dataBarang[i])
			i++
		}
	}
}

func tampilTransaksi() {
	if jumlahTransaksi == 0 {
		fmt.Println("Belum ada transaksi")
	} else {
		totalOmzet := 0
		i := 0

		for i < jumlahTransaksi {
			fmt.Println(dataTransaksi[i])
			totalOmzet += dataTransaksi[i].total
			i++
		}

		fmt.Println("Omzet Harian:", totalOmzet)
	}
}

func menu() {
	var pilih int = -1

	for pilih != 0 {
		fmt.Println("\n=== KASIR MINIMART ===")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tampil Barang")
		fmt.Println("3. Sort Asc (Selection)")
		fmt.Println("4. Sort Desc (Insertion)")
		fmt.Println("5. Transaksi")
		fmt.Println("6. Tampil Transaksi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih < 0 || pilih > 6 {
			fmt.Println("Pilihan tidak tersedia!")
		} else if pilih == 1 {
			tambahBarang()
		} else if pilih == 2 {
			tampilBarang()
		} else if pilih == 3 {
			selectionSort()
		} else if pilih == 4 {
			insertionSortDesc()
		} else if pilih == 5 {
			tambahTransaksi()
		} else if pilih == 6 {
			tampilTransaksi()
		}
	}
}

func main() {
	lanjut := awal()

	if lanjut == true {
		menu()
	}
}
