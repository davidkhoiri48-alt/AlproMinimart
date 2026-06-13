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