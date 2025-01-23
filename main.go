package main

import (
	"fmt"
)

func romanToArabic(roman string) int {
	romanNumerals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	n := len(roman)
	total := 0
	for i := 0; i < n; i++ {
		value := romanNumerals[roman[i]]
		if i+1 < n && value < romanNumerals[roman[i+1]] {
			total -= value
		} else {
			total += value
		}
	}
	return total
}

func arabicToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman += syb[i]
		}
	}
	return roman
}

func main() {
	fmt.Println(romanToArabic("X"))     // 10
	fmt.Println(romanToArabic("XXI"))   // 21
	fmt.Println(romanToArabic("IV"))    // 4
	fmt.Println(romanToArabic("XXXIX")) // 39

	fmt.Println(arabicToRoman(10)) // X
	fmt.Println(arabicToRoman(21)) // XXI
	fmt.Println(arabicToRoman(4))  // IV
	fmt.Println(arabicToRoman(39)) // XXXIX
}
