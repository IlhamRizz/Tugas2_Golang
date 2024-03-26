package main

import "fmt"

func main() {
	// Data uji
	scores1 := []int{96, 108, 89}
	scores2 := []int{88, 91, 110}
	scores3 := []int{97, 112, 101}
	scores4 := []int{109, 95, 123}
	scores5 := []int{97, 112, 101}
	scores6 := []int{109, 95, 106}

	// Menghitung rata-rata skor untuk setiap tim
	avgScoreTeamLumba := calculateAverage(scores1)
	avgScoreTeamKoala := calculateAverage(scores2)
	avgScoreTeamLumbaBonus1 := calculateAverage(scores3)
	avgScoreTeamKoalaBonus1 := calculateAverage(scores4)
	avgScoreTeamLumbaBonus2 := calculateAverage(scores5)
	avgScoreTeamKoalaBonus2 := calculateAverage(scores6)

	// Menentukan pemenang
	winner1 := determineWinner(avgScoreTeamLumba, avgScoreTeamKoala)
	winner2 := determineWinnerWithMinimum(avgScoreTeamLumbaBonus1, avgScoreTeamKoalaBonus1, 100)
	winner3 := determineWinnerWithMinimumAndTie(avgScoreTeamLumbaBonus2, avgScoreTeamKoalaBonus2, 100)

	// Menampilkan hasil
	fmt.Printf("Rata-rata skor Tim Lumba-lumba (Data 1): %.2f\n", avgScoreTeamLumba)
	fmt.Printf("Rata-rata skor Tim Koala (Data 1): %.2f\n", avgScoreTeamKoala)
	fmt.Println("Hasil (Data 1):", winner1)
	fmt.Printf("\nRata-rata skor Tim Lumba-lumba (Data Bonus 1): %.2f\n", avgScoreTeamLumbaBonus1)
	fmt.Printf("Rata-rata skor Tim Koala (Data Bonus 1): %.2f\n", avgScoreTeamKoalaBonus1)
	fmt.Println("Hasil (Data Bonus 1):", winner2)
	fmt.Printf("\nRata-rata skor Tim Lumba-lumba (Data Bonus 2): %.2f\n", avgScoreTeamLumbaBonus2)
	fmt.Printf("Rata-rata skor Tim Koala (Data Bonus 2): %.2f\n", avgScoreTeamKoalaBonus2)
	fmt.Println("Hasil (Data Bonus 2):", winner3)
}

// Fungsi untuk menghitung rata-rata
func calculateAverage(scores []int) float64 {
	total := 0
	for _, score := range scores {
		total += score
	}
	return float64(total) / float64(len(scores))
}

// Fungsi untuk menentukan pemenang tanpa mempertimbangkan skor minimum
func determineWinner(avgScoreTeamLumba, avgScoreTeamKoala float64) string {
	if avgScoreTeamLumba > avgScoreTeamKoala {
		return "Tim Lumba-lumba menang!"
	} else if avgScoreTeamKoala > avgScoreTeamLumba {
		return "Tim Koala menang!"
	}
	return "Hasil seri!"
}

// Fungsi untuk menentukan pemenang dengan mempertimbangkan skor minimum
func determineWinnerWithMinimum(avgScoreTeamLumba, avgScoreTeamKoala, minimumScore float64) string {
	if avgScoreTeamLumba >= minimumScore && avgScoreTeamLumba > avgScoreTeamKoala {
		return "Tim Lumba-lumba menang!"
	} else if avgScoreTeamKoala >= minimumScore && avgScoreTeamKoala > avgScoreTeamLumba {
		return "Tim Koala menang!"
	} else if avgScoreTeamLumba >= minimumScore && avgScoreTeamKoala >= minimumScore && avgScoreTeamLumba == avgScoreTeamKoala {
		return "Hasil seri!"
	}
	return "Tidak ada pemenang!"
}

// Fungsi untuk menentukan pemenang dengan mempertimbangkan skor minimum dan kondisi seri yang diperbarui
func determineWinnerWithMinimumAndTie(avgScoreTeamLumba, avgScoreTeamKoala, minimumScore float64) string {
	if avgScoreTeamLumba >= minimumScore && avgScoreTeamLumba > avgScoreTeamKoala {
		return "Tim Lumba-lumba menang!"
	} else if avgScoreTeamKoala >= minimumScore && avgScoreTeamKoala > avgScoreTeamLumba {
		return "Tim Koala menang!"
	} else if avgScoreTeamLumba >= minimumScore && avgScoreTeamKoala >= minimumScore && avgScoreTeamLumba == avgScoreTeamKoala {
		return "Hasil seri!"
	}
	return "Tidak ada pemenang!"
}
