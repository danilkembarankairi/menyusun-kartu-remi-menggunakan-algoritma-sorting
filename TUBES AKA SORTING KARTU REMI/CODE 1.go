package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Kartu struct {
	Angka int
	Jenis string
}

var urutanJenis = map[string]int{
	"♠": 1,
	"♥": 2,
	"♣": 3,
	"♦": 4,
}

func formatAngka(angka int) string {
	switch angka {
	case 1:
		return "AS"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	default:
		return fmt.Sprintf("%d", angka)
	}
}

func bandingkanKartu(a, b Kartu) bool {
	if a.Angka != b.Angka {
		return a.Angka < b.Angka
	}
	return urutanJenis[a.Jenis] > urutanJenis[b.Jenis]
}

func selectionSortIteratif(kartu []Kartu) []Kartu {
	n := len(kartu)
	for i := 0; i < n-1; i++ {
		indeksMin := i
		for j := i + 1; j < n; j++ {
			if bandingkanKartu(kartu[j], kartu[indeksMin]) {
				indeksMin = j
			}
		}
		kartu[i], kartu[indeksMin] = kartu[indeksMin], kartu[i]
	}
	return kartu
}

func selectionSortRekursif(kartu []Kartu, mulai int) []Kartu {
	n := len(kartu)
	if mulai >= n-1 {
		return kartu
	}
	indeksMin := mulai
	for i := mulai + 1; i < n; i++ {
		if bandingkanKartu(kartu[i], kartu[indeksMin]) {
			indeksMin = i
		}
	}
	kartu[mulai], kartu[indeksMin] = kartu[indeksMin], kartu[mulai]
	return selectionSortRekursif(kartu, mulai+1)
}

func binaryInsertionSortIteratif(kartu []Kartu) []Kartu {
	n := len(kartu)
	for i := 1; i < n; i++ {
		kunci := kartu[i]
		low, high := 0, i-1
		for low <= high {
			mid := (low + high) / 2
			if bandingkanKartu(kunci, kartu[mid]) {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		for j := i; j > low; j-- {
			kartu[j] = kartu[j-1]
		}
		kartu[low] = kunci
	}
	return kartu
}

func binaryInsertionSortRekursif(kartu []Kartu, n int) []Kartu {
	if n <= 1 {
		return kartu
	}
	kartu = binaryInsertionSortRekursif(kartu, n-1)
	kunci := kartu[n-1]
	low, high := 0, n-2
	for low <= high {
		mid := (low + high) / 2
		if bandingkanKartu(kunci, kartu[mid]) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	for j := n - 1; j > low; j-- {
		kartu[j] = kartu[j-1]
	}
	kartu[low] = kunci
	return kartu
}

func tampilkanKartu(kartu []Kartu) {
	for i, k := range kartu {
		fmt.Printf("[%-7s %s] ", formatAngka(k.Angka), k.Jenis)
		if (i+1)%7 == 0 {
			fmt.Println()
		}
	}
	if len(kartu)%7 != 0 {
		fmt.Println()
	}
}

func main() {
	kartu := []Kartu{}
	jumlahKartu := 250

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Apakah anda ingin menampilkan data kartu? (y/n): ")
	tampilkanData, _ := reader.ReadString('\n')
	tampilkanData = strings.TrimSpace(tampilkanData)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < jumlahKartu; i++ {
		kartu = append(kartu, Kartu{
			Angka: i%13 + 1,
			Jenis: []string{"♠", "♥", "♣", "♦"}[i%4],
		})
	}

	if strings.ToLower(tampilkanData) == "y" {
		fmt.Println("\nData Kartu Sebelum Diurutkan:")
		tampilkanKartu(kartu)
	} else {
		fmt.Println("\nUser memilih untuk tidak menampilkan data kartu")
	}

	mulai := time.Now()
	hasilSelectionIteratif := selectionSortIteratif(append([]Kartu{}, kartu...))
	waktu := time.Since(mulai).Seconds()

	if strings.ToLower(tampilkanData) == "y" {
		fmt.Println("\nHasil Kartu Setelah Diurutkan:")
		tampilkanKartu(hasilSelectionIteratif)
	}

	fmt.Printf("\nWaktu Running Time Selection Sort (Iteratif): %.6f detik\n", waktu)

	mulai = time.Now()
	_ = binaryInsertionSortIteratif(append([]Kartu{}, kartu...))
	waktu = time.Since(mulai).Seconds()
	fmt.Printf("Waktu Running Time Binary Insertion Sort (Iteratif): %.6f detik\n", waktu)

	mulai = time.Now()
	_ = selectionSortRekursif(append([]Kartu{}, kartu...), 0)
	waktu = time.Since(mulai).Seconds()
	fmt.Printf("Waktu Running Time Selection Sort (Rekursif): %.6f detik\n", waktu)

	mulai = time.Now()
	_ = binaryInsertionSortRekursif(append([]Kartu{}, kartu...), len(kartu))
	waktu = time.Since(mulai).Seconds()
	fmt.Printf("Waktu Running Time Binary Insertion Sort (Rekursif): %.6f detik\n", waktu)
}
