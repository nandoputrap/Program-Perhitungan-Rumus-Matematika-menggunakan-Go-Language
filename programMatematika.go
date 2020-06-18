// Program matematika sederhana
// Created by: ndo_kun
// Date: June 15 2020

package main

import (
	"fmt"
)

func main() {
	var pilihBangun, pilihBangunDatar, pilihBangunRuang int
	var p, l, t, s, a, r, luas, volume float64
	phi := 3.14
	fmt.Println("\n==============================================")
	fmt.Println("Selamat Datang di Program Matematika Sederhana")
	fmt.Println("==============================================\n")
	fmt.Println("MENU:")
	fmt.Println("1. Bangun Datar")
	fmt.Println("2. Bangun Ruang")
	fmt.Println("3. Exit\n")
	fmt.Print("Pilih menu nomor berapa? ")

	// scan untuk input angka, scanf untuk string
	fmt.Scan(&pilihBangun)
	// fmt.Scanf("%d", strconv.Atoi(pilihBangun))
	switch pilihBangun {
	case 1:
		fmt.Println("Luas Bangun Datar")
		fmt.Println("1. Persegi")
		fmt.Println("2. Persegi Panjang")
		fmt.Println("3. Segitiga")
		fmt.Println("4. Lingkaran")
		fmt.Print("Pilih Bangun Datar: ")
		fmt.Scan(&pilihBangunDatar)
		switch pilihBangunDatar {
		case 1:
			fmt.Println("Luas Persegi")
			fmt.Print("Masukkan sisi: ")
			fmt.Scan(&s)
			luas = s * s
			fmt.Println(luas)
		case 2:
			fmt.Println("Luas Persegi Panjang")
			fmt.Print("Masukkan panjang: ")
			fmt.Scan(&p)
			fmt.Print("Masukkan lebar: ")
			fmt.Scan(&l)
			luas = p * l
			fmt.Println(luas)
		case 3:
			fmt.Println("Luas Segitiga")
			fmt.Print("Masukkan alas: ")
			fmt.Scan(&a)
			fmt.Print("Masukkan tinggi: ")
			fmt.Scan(&t)
			luas = a * t / 2
			fmt.Println(luas)
		case 4:
			fmt.Println("Luas lingkaran")
			fmt.Print("Masukkan jari-jari: ")
			fmt.Scan(&r)
			luas = phi * (r * r)
			fmt.Println(luas)
		default:
			fmt.Println("Bangun ruang tidak ada")
		}
	case 2:
		fmt.Println("Volume Bangun Ruang")
		fmt.Println("1. Kubus")
		fmt.Println("2. Balok")
		fmt.Println("3. Kerucut")
		fmt.Println("4. Bola")
		fmt.Print("Pilih Bangun Ruang: ")
		fmt.Scan(&pilihBangunRuang)
		switch pilihBangunRuang {
		case 1:
			fmt.Println("Volume Kubus")
			fmt.Print("Masukkan sisi: ")
			fmt.Scan(&s)
			volume = s * s * s
			fmt.Println(volume)
		case 2:
			fmt.Println("Volume Balok")
			fmt.Print("Masukkan panjang: ")
			fmt.Scan(&p)
			fmt.Print("Masukkan lebar: ")
			fmt.Scan(&l)
			fmt.Print("Masukkan tinggi: ")
			fmt.Scan(&t)
			volume = p * l * t
			fmt.Println(volume)
		case 3:
			fmt.Println("Volume Kerucut")
			fmt.Print("Masukkan jari-jari: ")
			fmt.Scan(&r)
			fmt.Print("Masukkan tinggi: ")
			fmt.Scan(&t)
			// di golang jika ingin membagi float harus menggunakan angka dibelakang koma. Jika tidak maka hasilnya akan dianggap integer
			volume = (1.0 / 3.0) * phi * (r * r) * t
			fmt.Println(volume)
		case 4:
			fmt.Println("Volume Bola")
			fmt.Print("Masukkan jari-jari: ")
			fmt.Scan(&r)
			volume = (4.0 / 3.0) * phi * (r * r * r)
			fmt.Println(volume)
		default:
			fmt.Println("Bangun ruang tidak ada")
		}
	default:
		fmt.Println("Terima Kasih...")
	}
}
