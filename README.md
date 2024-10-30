# Tugas ke-dua Bootcamp Golang
Anda diminta untuk membuat sebuah program untuk manajemen inventaris di sebuah toko elektronik. Program ini akan menyimpan informasi tentang produk elektronik seperti nama produk, kategori, harga, dan stok yang tersedia. Pemilik toko ingin dapat melihat data semua produk, menambahkan produk baru, dan menghitung total nilai inventaris berdasarkan harga dan stok produk.

## Detail Tugas :
#### 1. Tipe Data & Variabel
Buat variabel yang menyimpan kategori produk elektronik seperti "Laptop",
"Smartphone", dan "Tablet" menggunakan tipe data yang sesuai (contoh: string).

#### 2. Struct
Buatlah struct bernama Product yang memiliki properti:
- Name: Nama produk (tipe data string).
- Category: Kategori produk (tipe data string).
- Price: Harga produk (tipe data float64).
- Stock: Jumlah stok produk (tipe data int).

#### 3. Array
Buatlah array atau slice yang menyimpan beberapa produk elektronik yang
dimiliki oleh toko. Contoh data:go  [{"Laptop A", "Laptop", 10000000, 5}, {"Smartphone B", "Smartphone", 5000000, 10}, {"Tablet C", "Tablet", 3000000, 7}]

#### 4. Function
- Buatlah fungsi AddProduct yang menerima parameter berupa nama produk, kategori, harga, dan stok, kemudian menambahkan produk tersebut ke dalam array produk yang sudah ada.
- Buatlah fungsi DisplayProducts yang menampilkan daftar semua produk yang tersedia di toko, termasuk nama, kategori, harga, dan stok.

#### 5. Looping
Buatlah fungsi DisplayProducts yang menampilkan daftar semua produk yang tersedia di toko, termasuk nama, kategori, harga, dan stok.

#### 6. Tantangan Ekstra
Tambahkan validasi untuk memasukan bahwa stok tidak boleh negatif dan harga tidak boleh 0 atau negatif.

## Spesifikasi Program:
Program harus menyediakan menu berikut:
1. Tampilkan Semua Produk: Menampilkan daftar semua produk beserta harga dan
stoknya.
2. Tambah Produk Baru: Menginput data produk baru (nama, kategori, harga, dan stok).
3. Hitung Total Nilai Inventaris: Menghitung total nilai dari seluruh produk (harga x stok).
4. Keluar: Mengakhiri program.

## Contoh Output:
Selamat datang di Program Manajemen Inventaris Toko Elektronik ---------------------------------------------
##### Pilih opsi:
1. Tampilkan Semua Produk
2. Tambah Produk Baru
3. Hitung Total Nilai Inventaris
4. Keluar

Pilihan Anda: 1

##### Daftar Produk:
1. Laptop A - Kategori: Laptop - Harga: 10000000 - Stok: 5
2. Smartphone B - Kategori: Smartphone - Harga: 5000000 - Stok: 10 3. Tablet C - Kategori: Tablet - Harga: 3000000 - Stok: 7

##### Pilih opsi:
1. Tampilkan Semua Produk
2. Tambah Produk Baru
3. Hitung Total Nilai Inventaris 4. Keluar

##### Pilihan Anda: 3
Total Nilai Inventaris: Rp. 120,000,000
