package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Student struct {
	Id      int
	Name    string
	Address string
	Job     string
	Reason  string
}

func main() {
	students := []Student{
		{1, "Pram", "Surabaya", "Backend Engineer", "Mau belajar"},
		{2, "Rizki", "Jakarta", "Frontend Engineer", "Mau belajar"},
		{3, "Arul", "Bandung", "DevOps Engineer", "Mau belajar"},
		{4, "Zidan", "Semarang", "Data Engineer", "Mau belajar"},
		{5, "Oni", "Yogyakarta", "Machine Learning Engineer", "Mau belajar"},
	}

	number := ""
	if len(os.Args) > 1 {
		number = os.Args[1]
	} else {
		fmt.Println("Tidak ada input")
		return
	}

	n, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Tidak bisa mengkonversi ke angka", err)
		return
	}

	index, ok := slices.BinarySearchFunc(students, Student{n, "", "", "", ""}, func(a, b Student) int {
		return cmp.Compare(a.Id, b.Id)
	})

	if ok {
		fmt.Println("Absen:", students[index].Id)
		fmt.Println("Nama:", students[index].Name)
		fmt.Println("Alamat:", students[index].Address)
		fmt.Println("Pekerjaan:", students[index].Job)
		fmt.Println("Alasan:", students[index].Reason)
	} else {
		fmt.Println("SISWA TIDAK DITEMUKAN!")
	}
}
