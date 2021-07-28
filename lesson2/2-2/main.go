package main

import (
	"math"
	"strconv"
)

func main() {
	// 整数
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	println(integer8, integer16, integer32, integer64)

	// 32bit以上になるから定義できない
	// var overInteger32 int32 = 2147483648
	// println(overInteger32)

	// 符号付きだから定義できない
	// var integer uint = -10
	// println(integer)

	// 浮動小数点
	var float32var float32 = 2147483647
	var float64var float64 = 9223372036854775807
	println(float32var, float64var)
	println(math.MaxFloat32, math.MaxFloat64)

	const (
		e        = 2.71828
		Avogadro = 6.02214129e23
		Planck   = 6.62606957e-34
	)
	println(e, Avogadro, Planck)

	// bool値
	var featureFlag bool = true
	println(featureFlag)

	// 文字列
	var firstName string = "John"
	lastName := "Doe"
	println(firstName, lastName)
	fullName := "John Doe \t(alias \"Foo\")\n"
	println(fullName)

	// 既定値
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)

	// キャスト
	var int16var int16 = 127
	var int32var int32 = 32767
	println(int32(int16var) + int32var)
	// パッケージを使ったキャスト
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	println(i, s)
}
