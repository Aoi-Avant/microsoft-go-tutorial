package main

import "fmt"

func main() {
	// マップ
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}

	fmt.Println(studentsAge)

	// マップの初期化
	studentsAge2 := make(map[string]int)

	fmt.Println(studentsAge2)

	// 要素の追加
	studentsAge2["john"] = 32
	studentsAge2["bob"] = 31

	fmt.Println(studentsAge2)

	// 追加時にエラーになるパターン（nilマップ）
	var studentsAge3 map[string]int
	// studentsAge3["john"] = 32
	// studentsAge3["bob"] = 31
	fmt.Println(studentsAge3)

	// 要素にアクセスする
	fmt.Println("Bob's age is", studentsAge2["bob"])
	fmt.Println("Christy's age is", studentsAge2["christy"]) // ない場合は既定値を返す

	// 存在の確認
	age, exist := studentsAge2["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}

	// 項目の削除
	delete(studentsAge2, "john")
	delete(studentsAge2, "christy")
	fmt.Println(studentsAge2)

	// マップ内でのループ
	studentsAge2["john"] = 32
	for name, age := range studentsAge2 {
		fmt.Printf("%s\t%d\n", name, age)
	}
	for _, age := range studentsAge2 {
		fmt.Printf("Ages %d\n", age)
	}
	for name := range studentsAge2 {
		fmt.Printf("Names %s\n", name)
	}
}
