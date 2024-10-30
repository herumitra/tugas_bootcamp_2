package main

import (
	"fmt"
)

// Definisikan struct untuk Produk
type Product struct {
	Name     string
	Category string
	Price    float64
	Stock    int
}

// Daftar kategori produk
var categories = []string{"Laptop", "Smartphone", "Tablet"}

// Slice untuk menyimpan daftar produk
var products []Product

// Fungsi untuk menambah produk baru ke dalam slice produk
func AddProduct(name, category string, price float64, stock int) {
	if price <= 0 {
		fmt.Println("Harga tidak boleh 0 atau negatif.")
		return
	}
	if stock < 0 {
		fmt.Println("Stok tidak boleh negatif.")
		return
	}

	// Menambah produk ke daftar produk
	product := Product{
		Name:     name,
		Category: category,
		Price:    price,
		Stock:    stock,
	}
	products = append(products, product)
	fmt.Println("Produk berhasil ditambahkan.")
}

// Fungsi untuk menampilkan semua produk yang ada di toko
func DisplayProducts() {
	if len(products) == 0 {
		fmt.Println("Tidak ada produk di inventaris.")
		return
	}

	fmt.Println("Daftar Produk:")
	for i, product := range products {
		fmt.Printf("%d. %s - Kategori: %s - Harga: Rp.%.2f - Stok: %d\n", i+1, product.Name, product.Category, product.Price, product.Stock)
	}
}

// Fungsi untuk menghitung total nilai inventaris
func CalculateTotalInventoryValue() float64 {
	totalValue := 0.0
	for _, product := range products {
		totalValue += product.Price * float64(product.Stock)
	}
	return totalValue
}

// Fungsi untuk menampilkan menu dan menangani pilihan pengguna
func main() {
	// Tambahkan produk awal untuk contoh
	AddProduct("Laptop MacBook", "Laptop", 25000000, 5)
	AddProduct("Smartphone Iphone", "Smartphone", 7500000, 10)
	AddProduct("Tablet Ipad", "Tablet", 15000000, 7)

	for {
		fmt.Println("\nSelamat datang di Program Manajemen Inventaris Toko Elektronik")
		fmt.Println("-------------------------------------------------------------")
		fmt.Println("1. Tampilkan Semua Produk")
		fmt.Println("2. Tambah Produk Baru")
		fmt.Println("3. Hitung Total Nilai Inventaris")
		fmt.Println("4. Keluar")
		fmt.Print("Pilihan Anda: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			DisplayProducts()
		case 2:
			var name, category string
			var price float64
			var stock int

			fmt.Print("Nama Produk: ")
			fmt.Scanln(&name)
			fmt.Print("Kategori Produk (Laptop/Smartphone/Tablet): ")
			fmt.Scanln(&category)
			fmt.Print("Harga Produk: ")
			fmt.Scanln(&price)
			fmt.Print("Stok Produk: ")
			fmt.Scanln(&stock)

			AddProduct(name, category, price, stock)
		case 3:
			totalValue := CalculateTotalInventoryValue()
			fmt.Printf("Total Nilai Inventaris: Rp.%.2f\n", totalValue)
		case 4:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
