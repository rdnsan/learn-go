package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 90

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 75

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)
}
