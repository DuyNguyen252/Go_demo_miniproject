package main

import (
	"fmt"
)

func main() {
	const tentruong = "THCS VO TRUONG TOAN"
	var name string
	var diem int

	fmt.Printf("--- Chao mung den voi %s ---\n", tentruong)

	fmt.Print("Nhap ten sinh vien: ")
	fmt.Scanln(&name)

	fmt.Print("Ban muon nhap diem cho bao nhieu mon? ")
	fmt.Scanln(&diem)

	var totalScore float64 = 0

	for i := 1; i <= diem; i++ {
		var score float64
		fmt.Printf("Nhap diem mon hoc %d: ", i)
		fmt.Scanln(&score)

		if score < 0 || score > 10 {
			fmt.Println("Diem khong dung - Vui long nhap lai")
			i--
			continue
		}

		totalScore += score
	}

	average := totalScore / float64(diem)

	fmt.Printf("\n--- Ket qua cho sinh vien: %s ---\n", name)
	fmt.Printf("Tong diem: %.2f\n", totalScore)
	fmt.Printf("Diem trung binh: %.2f\n", average)

	fmt.Print("Xep loai: ")
	switch {
	case average >= 8.0:
		fmt.Println("Gioi")
	case average >= 6.5:
		fmt.Println("Kha")
	case average >= 5.0:
		fmt.Println("Trung binh")
	default:
		fmt.Println("Yeu")
	}
}
