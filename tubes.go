package main

import "fmt"

const NMAX_BARANG int = 100
const NMAX_TRANSAKSI int = 200

type Harga int
type Stok int
type TanggalStr string

type Barang struct {
	id       int
	nama     string
	harga    Harga
	stok     Stok
	kategori string
}

type ItemTransaksi struct {
	idBarang   int
	namaBarang string
	jumlah     int
	subtotal   Harga
}

type Transaksi struct {
	id      int
	tanggal TanggalStr
	items   [10]ItemTransaksi
	jmlItem int
	total   Harga
}

type arrBarang [NMAX_BARANG]Barang
type arrTransaksi [NMAX_TRANSAKSI]Transaksi

var dataBarang arrBarang
var dataTransaksi arrTransaksi

func seqSearchBarang(T arrBarang, n int, idCari int) int {
	var found int = -1
	var j int = 0
	for j < n && found == -1 {
		if T[j].id == idCari {
			found = j
		}
		j = j + 1
	}
	return found
}

func binarySearchBarang(T arrBarang, n int, idCari int) int {
	var found int = -1
	var mid int
	var left int = 0
	var right int = n - 1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if idCari < T[mid].id {
			right = mid - 1
		} else if idCari > T[mid].id {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

func selectionSortBarang(T *arrBarang, n int, field int, asc bool) {
	var i, j, idx_eks int
	var t Barang
	i = 1
	for i <= n-1 {
		idx_eks = i - 1
		j = i
		for j < n {
			var lebihEks bool
			if field == 1 {
				if asc {
					lebihEks = T[j].id < T[idx_eks].id
				} else {
					lebihEks = T[j].id > T[idx_eks].id
				}
			} else {
				if asc {
					lebihEks = T[j].harga < T[idx_eks].harga
				} else {
					lebihEks = T[j].harga > T[idx_eks].harga
				}
			}
			if lebihEks {
				idx_eks = j
			}
			j = j + 1
		}
		t = T[idx_eks]
		T[idx_eks] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func kondisiGeserTransaksi(T arrTransaksi, temp Transaksi, j int, field int, asc bool) bool {
	var hasil bool = false
	if j > 0 {
		if field == 1 {
			if asc {
				hasil = temp.id < T[j-1].id
			} else {
				hasil = temp.id > T[j-1].id
			}
		} else {
			if asc {
				hasil = temp.total < T[j-1].total
			} else {
				hasil = temp.total > T[j-1].total
			}
		}
	}
	return hasil
}

func insertionSortTransaksi(T *arrTransaksi, n int, field int, asc bool) {
	var i, j int
	var temp Transaksi
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for kondisiGeserTransaksi(*T, temp, j, field, asc) {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func hitungTotal(items [10]ItemTransaksi, idx int, jmlItem int) Harga {
	if idx >= jmlItem {
		return 0
	}
	return items[idx].subtotal + hitungTotal(items, idx+1, jmlItem)
}

func tambahBarang(T *arrBarang, n *int) {
	if *n >= NMAX_BARANG {
		fmt.Println("Data barang sudah penuh!")
	} else {
		var b Barang
		fmt.Print("ID Barang  : ")
		fmt.Scan(&b.id)

		if seqSearchBarang(*T, *n, b.id) != -1 {
			fmt.Println("ID sudah digunakan, gunakan ID lain!")
		} else {
			fmt.Print("Nama Barang: ")
			fmt.Scan(&b.nama)
			fmt.Print("Harga      : ")
			var h int
			fmt.Scan(&h)
			fmt.Print("Stok       : ")
			var s int
			fmt.Scan(&s)
			fmt.Print("Kategori   : ")
			fmt.Scan(&b.kategori)
			if h <= 0 && s < 0 {
				fmt.Println("Harga dan stok harus lebih dari 0 !")
			} else if h <= 0 {
				fmt.Println("Harga harus lebih dari 0!")
			} else if s < 0 {
				fmt.Println("Stok harus lebih dari 0!")
			} else {
				b.harga = Harga(h)
				b.stok = Stok(s)
				T[*n] = b
				*n = *n + 1
				fmt.Println("Barang berhasil ditambahkan!")
			}
		}
	}
}

func editBarang(T *arrBarang, n int) {
	if n == 0 {
		fmt.Println("Belum ada data barang!")
	} else {
		var id int
		fmt.Print("Masukkan ID barang yang ingin diubah: ")
		fmt.Scan(&id)

		var idx int = seqSearchBarang(*T, n, id)

		if idx == -1 {
			fmt.Println("Barang dengan ID tersebut tidak ditemukan!")
		} else {
			fmt.Printf("Data saat ini: ID=%d | Nama=%s | Harga=%d | Stok=%d | Kategori=%s\n",
				T[idx].id, T[idx].nama, T[idx].harga, T[idx].stok, T[idx].kategori)
			fmt.Println("Masukkan data baru!")

			fmt.Print("Nama Barang: ")
			fmt.Scan(&T[idx].nama)
			fmt.Print("Harga      : ")
			var h int
			fmt.Scan(&h)
			fmt.Print("Stok       : ")
			var s int
			fmt.Scan(&s)
			fmt.Print("Kategori   : ")
			fmt.Scan(&T[idx].kategori)
			if h <= 0 && s < 0 {
				fmt.Println("Harga dan stok harus lebih dari 0 !")
			} else if h <= 0 {
				fmt.Println("Harga harus lebih dari 0!")
			} else if s < 0 {
				fmt.Println("Stok harus lebih dari 0!")
			} else {
				T[idx].harga = Harga(h)
				T[idx].stok = Stok(s)
				fmt.Println("Data barang berhasil diubah!")
			}
		}
	}
}

func hapusBarang(T *arrBarang, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data barang!")
	} else {
		var id int
		fmt.Print("Masukkan ID barang yang ingin dihapus: ")
		fmt.Scan(&id)

		var idx int = seqSearchBarang(*T, *n, id)

		if idx == -1 {
			fmt.Println("Barang dengan ID tersebut tidak ditemukan!")
		} else {
			fmt.Printf("Menghapus: %s (ID: %d)\n", T[idx].nama, T[idx].id)
			var i int = idx
			for i < *n-1 {
				T[i] = T[i+1]
				i = i + 1
			}
			T[*n-1] = Barang{}
			*n = *n - 1
			fmt.Println("Barang berhasil dihapus!")
		}
	}
}

func tampilBarang(T arrBarang, n int) {
	if n == 0 {
		fmt.Println("Belum ada data barang.")
	} else {
		fmt.Println("\nUrutkan berdasarkan:")
		fmt.Println("1. ID (Terkecil)")
		fmt.Println("2. ID (Terbesar)")
		fmt.Println("3. Harga (Termurah)")
		fmt.Println("4. Harga (Termahal)")
		fmt.Print("Pilih: ")
		var pSort int
		fmt.Scan(&pSort)

		if pSort == 1 {
			selectionSortBarang(&T, n, 1, true)
		} else if pSort == 2 {
			selectionSortBarang(&T, n, 1, false)
		} else if pSort == 3 {
			selectionSortBarang(&T, n, 2, true)
		} else if pSort == 4 {
			selectionSortBarang(&T, n, 2, false)
		} else {
			fmt.Println("Pilihan tidak valid, Pilih 1/2/3/4!.")
		}

		fmt.Println("\n------------------------------------------------------------")
		fmt.Printf("%-5s %-12s %-12s %-12s %-12s\n", "ID", "Nama", "Harga", "Stok", "Kategori")
		fmt.Println("------------------------------------------------------------")
		var i int = 0
		for i < n {
			b := T[i]
			fmt.Printf("%-5d %-11s %-9d %-8d %-9s\n",
				b.id, b.nama, b.harga, b.stok, b.kategori)
			i = i + 1
		}
		fmt.Println("------------------------------------------------------------")
	}
}

func cariBarang(T *arrBarang, n int) {
	if n == 0 {
		fmt.Println("Belum ada data barang!")
	} else {
		var id int
		fmt.Print("Masukkan ID barang yang dicari: ")
		fmt.Scan(&id)

		var salinan arrBarang = *T
		selectionSortBarang(&salinan, n, 1, true)
		var idx int = binarySearchBarang(salinan, n, id)

		if idx == -1 {
			fmt.Println("Barang tidak ditemukan.")
		} else {
			b := salinan[idx]
			fmt.Println("\n=== Detail Barang ===")
			fmt.Printf("ID       : %d\n", b.id)
			fmt.Printf("Nama     : %s\n", b.nama)
			fmt.Printf("Harga    : Rp%d\n", b.harga)
			fmt.Printf("Stok     : %d\n", b.stok)
			fmt.Printf("Kategori : %s\n", b.kategori)
		}
	}
}

func tambahTransaksi(TB *arrBarang, nB int, TT *arrTransaksi, nT *int) {
	if *nT >= NMAX_TRANSAKSI {
		fmt.Println("Data transaksi sudah penuh!")
	} else if nB == 0 {
		fmt.Println("Belum ada data barang!")
	} else {
		var tanggal string
		fmt.Print("Tanggal transaksi (YYYY-MM-DD): ")
		fmt.Scan(&tanggal)

		var items [10]ItemTransaksi
		var jmlItem int = 0
		var lanjut bool = true

		for jmlItem < 10 && lanjut {
			fmt.Printf("\n--- Item ke-%d ---\n", jmlItem+1)
			fmt.Print("ID Barang (0 = selesai): ")
			var id int
			fmt.Scan(&id)

			if id == 0 {
				lanjut = false
			} else {
				var idx int = seqSearchBarang(*TB, nB, id)

				if idx == -1 {
					fmt.Println("ID barang tidak ditemukan!")
				} else if TB[idx].stok <= 0 {
					fmt.Println("Stok barang habis!")
				} else {
					fmt.Printf("Barang: %s | Harga: Rp%d | Stok: %d\n",
						TB[idx].nama, TB[idx].harga, TB[idx].stok)
					fmt.Print("Jumlah beli: ")
					var jumlah int
					fmt.Scan(&jumlah)

					if jumlah <= 0 {
						fmt.Println("Jumlah tidak valid!")
					} else if int(TB[idx].stok) < jumlah {
						fmt.Printf("Stok tidak cukup! Stok tersedia: %d\n", TB[idx].stok)
					} else {
						var subtotal Harga = TB[idx].harga * Harga(jumlah)
						items[jmlItem] = ItemTransaksi{
							idBarang:   id,
							namaBarang: TB[idx].nama,
							jumlah:     jumlah,
							subtotal:   subtotal,
						}
						TB[idx].stok = TB[idx].stok - Stok(jumlah)
						jmlItem = jmlItem + 1
						fmt.Printf("Ditambahkan. Subtotal: Rp%d\n", subtotal)
					}
				}
			}
		}

		if jmlItem == 0 {
			fmt.Println("Tidak ada item, transaksi dibatalkan.")
		} else {
			var total Harga = hitungTotal(items, 0, jmlItem)

			TT[*nT] = Transaksi{
				id:      *nT + 1,
				tanggal: TanggalStr(tanggal),
				items:   items,
				jmlItem: jmlItem,
				total:   total,
			}
			*nT = *nT + 1

			fmt.Println("\n========== STRUK TRANSAKSI ==========")
			fmt.Printf("ID Transaksi : %d\n", *nT)
			fmt.Printf("Tanggal      : %s\n", tanggal)
			fmt.Println("--------------------------------------")
			var k int = 0
			for k < jmlItem {
				fmt.Printf("%-20s x%d  Rp%d\n",
					items[k].namaBarang, items[k].jumlah, items[k].subtotal)
				k = k + 1
			}
			fmt.Println("--------------------------------------")
			fmt.Printf("TOTAL        : Rp%d\n", total)
			fmt.Println("======================================")
		}
	}
}

func tampilTransaksi(T arrTransaksi, n int) {
	if n == 0 {
		fmt.Println("Belum ada data transaksi.")
	} else {
		fmt.Println("\nUrutkan berdasarkan:")
		fmt.Println("1. ID Transaksi (Ascending)")
		fmt.Println("2. ID Transaksi (Descending)")
		fmt.Println("3. Total (Ascending)")
		fmt.Println("4. Total (Descending)")
		fmt.Print("Pilih: ")
		var pSort int
		fmt.Scan(&pSort)

		if pSort == 1 {
			insertionSortTransaksi(&T, n, 1, true)
		} else if pSort == 2 {
			insertionSortTransaksi(&T, n, 1, false)
		} else if pSort == 3 {
			insertionSortTransaksi(&T, n, 2, true)
		} else if pSort == 4 {
			insertionSortTransaksi(&T, n, 2, false)
		} else {
			fmt.Println("Pilihan tidak valid, Pilih 1/2/3/4!.")
		}

		fmt.Println("\n--------------------------------------------------")
		fmt.Printf("%-5s %-12s %-6s %-12s\n", "ID", "Tanggal", "Item", "Total")
		fmt.Println("--------------------------------------------------")

		var totalSemua Harga
		var i int = 0
		for i < n {
			t := T[i]
			fmt.Printf("%-5d %-12s %-6d Rp%-12d\n",
				t.id, t.tanggal, t.jmlItem, t.total)
			totalSemua = totalSemua + t.total
			i = i + 1
		}
		fmt.Println("--------------------------------------------------")
		fmt.Printf("TOTAL SEMUA TRANSAKSI: Rp%d\n", totalSemua)
	}
}

func omzetHarian(T arrTransaksi, n int) {
	var tanggal string
	fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
	fmt.Scan(&tanggal)

	var total Harga
	var ada bool = false

	fmt.Printf("\n=== Transaksi pada %s ===\n", tanggal)
	fmt.Printf("%-5s %-6s %-12s\n", "ID", "Item", "Total")
	fmt.Println("------------------------")

	var i int = 0
	for i < n {
		if dataTransaksi[i].tanggal == TanggalStr(tanggal) {
			t := T[i]
			fmt.Printf("%-5d %-6d Rp%-12d\n", t.id, t.jmlItem, t.total)
			total = total + t.total
			ada = true
		}
		i = i + 1
	}

	if !ada {
		fmt.Println("Tidak ada transaksi pada tanggal tersebut.")
	} else {
		fmt.Println("------------------------")
		fmt.Printf("Omzet Harian: Rp%d\n", total)
	}
}

func LarisKurangLaku(TB arrBarang, nB int, TT arrTransaksi, nT int) {
	if nB == 0 || nT == 0 {
		fmt.Println("Belum ada data barang atau transaksi!")
		return
	}

	var totalTerjual [NMAX_BARANG]int

	var i, j, k int

	for i = 0; i < nT; i++ {
		for j = 0; j < TT[i].jmlItem; j++ {
			for k = 0; k < nB; k++ {
				if TT[i].items[j].idBarang == TB[k].id {
					totalTerjual[k] += TT[i].items[j].jumlah
				}
			}
		}
	}

	var idxMax, idxMin int = 0, 0

	for i = 1; i < nB; i++ {
		if totalTerjual[i] > totalTerjual[idxMax] {
			idxMax = i
		}

		if totalTerjual[i] < totalTerjual[idxMin] {
			idxMin = i
		}
	}

	fmt.Println("\n=== BARANG TERLARIS ===")
	fmt.Printf("Nama Barang : %s\n", TB[idxMax].nama)
	fmt.Printf("Total Terjual : %d\n", totalTerjual[idxMax])

	fmt.Println("\n=== BARANG KURANG LAKU ===")
	fmt.Printf("Nama Barang : %s\n", TB[idxMin].nama)
	fmt.Printf("Total Terjual : %d\n", totalTerjual[idxMin])
}

func awal() bool {
	var pilih int = -1
	for pilih != 1 && pilih != 2 {
		fmt.Println("==========================================")
		fmt.Println("=             MINIMART MEOW              =")
		fmt.Println("=         GERAI RETAIL SERBA ADA         =")
		fmt.Println("==========================================")
		fmt.Println("1. Masuk ke Sistem")
		fmt.Println("2. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)
		if pilih != 1 && pilih != 2 {
			fmt.Println("pilihan tidak tersedia!, pilih 1/2.")
		}
	}
	if pilih == 1 {
		return true
	} else {
		return false
	}
}

func menuBarang(nB *int) {
	var pilih int = -1
	for pilih != 0 {
		fmt.Println("\n--- MANAJEMEN BARANG ---")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Edit Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Tampil Semua Barang")
		fmt.Println("5. Cari Barang")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahBarang(&dataBarang, nB)
		} else if pilih == 2 {
			editBarang(&dataBarang, *nB)
		} else if pilih == 3 {
			hapusBarang(&dataBarang, nB)
		} else if pilih == 4 {
			tampilBarang(dataBarang, *nB)
		} else if pilih == 5 {
			cariBarang(&dataBarang, *nB)
		} else if pilih != 0 {
			fmt.Println("Pilihan tidak tersedia!, pilih 0/1/2/3/4/5")
		}
	}
}

func menuTransaksi(nB int, nT *int) {
	var pilih int = -1
	for pilih != 0 {
		fmt.Println("\n--- TRANSAKSI ---")
		fmt.Println("1. Catat Transaksi Baru")
		fmt.Println("2. Tampil Semua Transaksi")
		fmt.Println("3. Omzet Harian")
		fmt.Println("4. Barang Terlaris & Kurang Laku")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahTransaksi(&dataBarang, nB, &dataTransaksi, nT)
		} else if pilih == 2 {
			tampilTransaksi(dataTransaksi, *nT)
		} else if pilih == 3 {
			omzetHarian(dataTransaksi, *nT)
		} else if pilih == 4 {
			LarisKurangLaku(dataBarang, nB, dataTransaksi, *nT)
		} else if pilih != 0 {
			fmt.Println("Pilihan tidak tersedia!, pilih 0/1/2/3/4!")
		}
	}
}

func menu() {
	var nBarang int = 0
	var nTransaksi int = 0
	var pilih int = -1

	for pilih != 0 {
		fmt.Println("\n============================")
		fmt.Println("=    KASIR  MINIMART      =")
		fmt.Println("============================")
		fmt.Println("1. Manajemen Barang")
		fmt.Println("2. Transaksi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			menuBarang(&nBarang)
		} else if pilih == 2 {
			menuTransaksi(nBarang, &nTransaksi)
		} else if pilih == 0 {
			fmt.Println("\nTerima kasih telah menggunakan Meow Mart. Sampai jumpa!")
		} else {
			fmt.Println("Pilihan tidak tersedia!, pilih 0/1/2 ")
		}
	}
}

func main() {
	if awal() {
		menu()
	} else {
		fmt.Println("Program selesai.")
	}
}
