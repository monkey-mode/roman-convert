package main

import (
    "testing"
)

func TestRomanToArabic(t *testing.T) {
    tests := []struct {
        roman   string
        arabic  int
    }{
        {"I", 1},
        {"IV", 4},
        {"IX", 9},
        {"X", 10},
        {"XXI", 21},
        {"XXXIX", 39},
        {"XL", 40},
        {"L", 50},
        {"XC", 90},
        {"C", 100},
        {"CD", 400},
        {"D", 500},
        {"CM", 900},
        {"M", 1000},
    }

    for _, test := range tests {
        result := romanToArabic(test.roman)
        if result != test.arabic {
            t.Errorf("romanToArabic(%s) = %d; want %d", test.roman, result, test.arabic)
        }
    }
}

func TestArabicToRoman(t *testing.T) {
    tests := []struct {
        arabic  int
        roman   string
    }{
        {1, "I"},
        {4, "IV"},
        {9, "IX"},
        {10, "X"},
        {21, "XXI"},
        {39, "XXXIX"},
        {40, "XL"},
        {50, "L"},
		{69,"LXIX"},
        {90, "XC"},
        {100, "C"},
        {400, "CD"},
        {500, "D"},
        {900, "CM"},
        {1000, "M"},
    }

    for _, test := range tests {
        result := arabicToRoman(test.arabic)
        if result != test.roman {
            t.Errorf("arabicToRoman(%d) = %s; want %s", test.arabic, result, test.roman)
        }
    }
}