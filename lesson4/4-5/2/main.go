package main

import (
	"fmt"
	"strings"
)

func main() {
	var val string
	fmt.Print("Enter String : ")
	fmt.Scanf("%s", &val)

	fmt.Println(changeNum(val))
}

func changeNum(val string) int {
	roma := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	split_val := strings.Split(val, "")

	numebers := make([]int, len(split_val)+1)
	for i := 0; i < len(split_val); i++ {
		num, exist := roma[split_val[i]]

		if !exist {
			fmt.Println("Error!")
			return 0
		}

		numebers[i] = num
	}

	result := 0
	for i := 0; i < len(split_val); i++ {
		if numebers[i] < numebers[i+1] {
			result -= numebers[i]
		} else {
			result += numebers[i]
		}
	}

	return result
}

// 模範解答
// func romanToArabic(numeral string) int {
//     romanMap := map[rune]int{
//         'M': 1000,
//         'D': 500,
//         'C': 100,
//         'L': 50,
//         'X': 10,
//         'V': 5,
//         'I': 1,
//     }

//     arabicVals := make([]int, len(numeral)+1)

//     for index, digit := range numeral {
//         if val, present := romanMap[digit]; present {
//             arabicVals[index] = val
//         } else {
//             fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", numeral, digit)
//             return 0
//         }
//     }

//     total := 0

//     for index := 0; index < len(numeral); index++ {
//         if arabicVals[index] < arabicVals[index+1] {
//             arabicVals[index] = -arabicVals[index]
//         }
//         total += arabicVals[index]
//     }

//     return total
// }

// func main() {
//     fmt.Println("MCLX is", romanToArabic("MCLX"))
//     fmt.Println("MCMXCIX is ", romanToArabic("MCMXCIX"))
//     fmt.Println("MCMZ is", romanToArabic("MCMZ"))
// }
