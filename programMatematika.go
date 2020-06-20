// Program matematika sederhana
// Created by: ndo

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
		fmt.Println("\nLuas Bangun Datar\n")
		fmt.Println("1. Persegi")
		fmt.Println("2. Persegi Panjang")
		fmt.Println("3. Segitiga")
		fmt.Println("4. Lingkaran \n")
		fmt.Print("Pilih Bangun Datar: ")
		fmt.Scan(&pilihBangunDatar)
		switch pilihBangunDatar {
		case 1:
			fmt.Println("\nLuas Persegi")
			fmt.Print("Masukkan sisi (cm): ")
			fmt.Scan(&s)
			luas = s * s
			s := fmt.Sprint(s)
			luas := fmt.Sprint(luas)
			fmt.Println("\nLuas persegi dengan sisi " + s + " adalah " + luas + " cm2")
		case 2:
			fmt.Println("\nLuas Persegi Panjang")
			fmt.Print("Masukkan panjang (cm): ")
			fmt.Scan(&p)
			fmt.Print("Masukkan lebar (cm): ")
			fmt.Scan(&l)
			luas = p * l
			p := fmt.Sprint(p)
			l := fmt.Sprint(l)
			luas := fmt.Sprint(luas)
			fmt.Println("\nLuas persegi panjang dengan panjang " + p + " cm dan lebar " + l + " cm adalah " + luas + " cm2")
		case 3:
			fmt.Println("\nLuas Segi Tiga")
			fmt.Print("Masukkan alas (cm): ")
			fmt.Scan(&a)
			fmt.Print("Masukkan tinggi (cm): ")
			fmt.Scan(&t)
			luas = a * t / 2
			a := fmt.Sprint(a)
			t := fmt.Sprint(t)
			luas := fmt.Sprint(luas)
			fmt.Println("\nLuas segi tiga dengan alas " + a + " cm dan tinggi " + t + " cm adalah " + luas + " cm2")
		case 4:
			fmt.Println("\nLuas lingkaran")
			fmt.Print("Masukkan jari-jari (cm): ")
			fmt.Scan(&r)
			luas = phi * (r * r)
			r := fmt.Sprint(r)
			luas := fmt.Sprint(luas)
			fmt.Println("\nLuas segi tiga dengan jari-jari " + r + " cm adalah " + luas + " cm2")
		default:
			fmt.Println("\nBangun ruang tidak ada")
		}
	case 2:
		fmt.Println("\nVolume Bangun Ruang")
		fmt.Println("1. Kubus")
		fmt.Println("2. Balok")
		fmt.Println("3. Kerucut")
		fmt.Println("4. Bola")
		fmt.Print("\nPilih Bangun Ruang: ")
		fmt.Scan(&pilihBangunRuang)
		switch pilihBangunRuang {
		case 1:
			fmt.Println("\nVolume Kubus")
			fmt.Print("Masukkan sisi (cm): ")
			fmt.Scan(&s)
			volume = s * s * s
			s := fmt.Sprint(s)
			volume := fmt.Sprint(volume)
			fmt.Println("\nVolume kubus dengan sisi " + s + " adalah " + volume + " cm3")
		case 2:
			fmt.Println("\nVolume Balok")
			fmt.Print("Masukkan panjang (cm): ")
			fmt.Scan(&p)
			fmt.Print("Masukkan lebar (cm): ")
			fmt.Scan(&l)
			fmt.Print("Masukkan tinggi (cm): ")
			fmt.Scan(&t)
			volume = p * l * t
			p := fmt.Sprint(p)
			l := fmt.Sprint(l)
			t := fmt.Sprint(t)
			volume := fmt.Sprint(volume)
			fmt.Println("\nVolume balok dengan panjang " + p + " cm, lebar " + l + " cm, dan tinggi " + t + " cm adalah " + volume + " cm3")
		case 3:
			fmt.Println("\nVolume Kerucut")
			fmt.Print("Masukkan jari-jari (cm): ")
			fmt.Scan(&r)
			fmt.Print("Masukkan tinggi: (cm) ")
			fmt.Scan(&t)
			// di golang jika ingin membagi float harus menggunakan angka dibelakang koma. Jika tidak maka hasilnya akan dianggap integer
			volume = (1.0 / 3.0) * phi * (r * r) * t
			r := fmt.Sprint(r)
			t := fmt.Sprint(t)
			volume := fmt.Sprint(volume)
			fmt.Println("\nVolume kerucut dengan jari-jari " + r + " cm dan tinggi " + t + " cm adalah " + volume + " cm3")
		case 4:
			fmt.Println("\nVolume Bola")
			fmt.Print("Masukkan jari-jari: ")
			fmt.Scan(&r)
			volume = (4.0 / 3.0) * phi * (r * r * r)
			r := fmt.Sprint(r)
			volume := fmt.Sprint(volume)
			fmt.Println("\nVolume bola dengan jari-jari " + r + " cm adalah " + volume + " cm3")
		default:
			fmt.Println("\nBangun ruang tidak ada")
		}
	default:
		fmt.Println("\nMenu yang dipilih tidak ada...")
	}
}
