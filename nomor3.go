// Masalah utamanya adalah bahwa perubahan pada objek Bebek di dalam fungsi tidak akan tercermin di luar fungsi
// karena parameter Bebek diteruskan dengan nilai (pass by value), bukan referensi. Untuk memperbaiki masalah ini,
// Anda bisa menggunakan pointer untuk mengubah nilai Bebek di dalam fungsi. Berikut adalah versi yang telah diperbaiki:
package main

import "fmt"

type Bebek struct {
	energi       int
	hidup        bool
	bisaTerbang  bool
	suaraTerbang string
}

func Mati(b *Bebek) {
	b.hidup = false
}

func Terbang(b *Bebek) {
	if b.energi > 0 && b.hidup == true && b.bisaTerbang {
		fmt.Println(b.suaraTerbang)
		b.energi -= 1
		if b.energi == 0 {
			Mati(b)
		}
	}
}

func Makan(b *Bebek) {
	if b.energi > 0 && b.hidup == true {
		b.energi += 1
	}
}

func main() {
	bebek := Bebek{energi: 3, hidup: true, bisaTerbang: true, suaraTerbang: "Kwek kwek!"}

	Terbang(&bebek)
	fmt.Println("Energi setelah terbang:", bebek.energi)

	Makan(&bebek)
	fmt.Println("Energi setelah makan:", bebek.energi)

	Terbang(&bebek)
	fmt.Println("Energi setelah terbang lagi:", bebek.energi)
}
