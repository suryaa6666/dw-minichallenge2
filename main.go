package main

import (
	"fmt"

	"github.com/yudapc/go-rupiah"
)

func hitungVoucher(kodeVoucher string, totalHarga float64) {
	var harga, totalPotongan, totalPembayaran float64
	var diskon string

	if kodeVoucher == "DumbWMerchBerkah" && totalHarga >= 50000 {
		potongan := totalHarga * 25 / 100
		if potongan > 20000 {
			potongan = 20000
		}

		harga = totalHarga
		diskon = "25%"
		totalPotongan = potongan
		totalPembayaran = totalHarga - potongan
	}

	if kodeVoucher == "DumbMerchMurahBanget" && totalHarga >= 70000 {
		potongan := totalHarga * 50 / 100
		if potongan > 45000 {
			potongan = 45000
		}

		harga = totalHarga
		diskon = "50%"
		totalPotongan = potongan
		totalPembayaran = totalHarga - potongan
	}

	hargaToRupiah := rupiah.FormatRupiah(harga)
	totalPotonganToRupiah := rupiah.FormatRupiah(totalPotongan)
	totalPembayaranToRupiah := rupiah.FormatRupiah(totalPembayaran)

	fmt.Println("- Total Belanja : " + hargaToRupiah + "\n" + "- Diskon : " + diskon + "\n" + "- Potongan : " + totalPotonganToRupiah + "\n" + "- Total Pembayaran : " + totalPembayaranToRupiah)
}

func main() {
	hitungVoucher("DumbMerchMurahBanget", 120000)
}
