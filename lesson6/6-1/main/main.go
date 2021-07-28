package main

import (
	"fmt"
	"strings"

	"geometry/geometry"
)

type triangle struct {
	size int
}

type square struct {
	size int
}

type upperstring string

type coloredTriangle struct {
	triangle
	color string
}

func main() {
	t := triangle{3}
	fmt.Println("Perimeter(triangle):", t.perimeter())

	// メソッド内のポインター
	t.doubleSize()
	fmt.Println("Size(triangle):", t.size)
	fmt.Println("Perimeter(triangle):", t.perimeter())

	s := square{4}
	fmt.Println("Perimeter(square):", s.perimeter())

	// 他の型のメソッドを宣言
	str := upperstring("Learning Go!")
	fmt.Println(str)
	fmt.Println(str.Upper())

	// メソッドの埋め込み
	// メソッドのオーバーロード
	ct := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Size(coloredTriangle):", ct.size)
	fmt.Println("Perimeter(coloredTriangle):", ct.perimeter())
	fmt.Println("Perimeter(normal):", ct.triangle.perimeter())

	// カプセル化
	pt := geometry.Triangle{}
	pt.SetSize(3)
	fmt.Println("Perimeter(package)", pt.Perimeter())
	// fmt.Println("Size(package):", pt.size)
}

func (t *triangle) perimeter() int {
	return t.size * 3
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

func (s square) perimeter() int {
	return s.size * 4
}

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}
