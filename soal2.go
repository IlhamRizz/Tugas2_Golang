package main

import "fmt"

func main() {
	// Data 1
	weightMark1 := 78.0 // kg
	heightMark1 := 1.69 // meter
	weightJohn1 := 92.0 // kg
	heightJohn1 := 1.95 // meter

	// Data 2
	weightMark2 := 95.0 // kg
	heightMark2 := 1.88 // meter
	weightJohn2 := 85.0 // kg
	heightJohn2 := 1.76 // meter

	// Menghitung BMI
	bmiMark1 := calculateBMI(weightMark1, heightMark1)
	bmiJohn1 := calculateBMI(weightJohn1, heightJohn1)

	bmiMark2 := calculateBMI(weightMark2, heightMark2)
	bmiJohn2 := calculateBMI(weightJohn2, heightJohn2)

	// Membandingkan BMI
	markHigherBMI1 := bmiMark1 > bmiJohn1
	markHigherBMI2 := bmiMark2 > bmiJohn2

	// Menampilkan hasil
	fmt.Println("Data 1:")
	fmt.Printf("BMI Mark: %.2f\n", bmiMark1)
	fmt.Printf("BMI John: %.2f\n", bmiJohn1)
	fmt.Printf("Mark memiliki BMI lebih tinggi dari John? %t\n", markHigherBMI1)

	fmt.Println("\nData 2:")
	fmt.Printf("BMI Mark: %.2f\n", bmiMark2)
	fmt.Printf("BMI John: %.2f\n", bmiJohn2)
	fmt.Printf("Mark memiliki BMI lebih tinggi dari John? %t\n", markHigherBMI2)
}

// Fungsi untuk menghitung BMI
func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}
