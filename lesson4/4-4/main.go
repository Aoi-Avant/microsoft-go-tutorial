package main

import (
	"encoding/json"
	"fmt"
)

// 構造体
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

// 構造体の埋め込み
// jsonでエンコード、デコード
type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Student struct {
	Information Person
	StudentID   int
}

type Teacher struct {
	Person
	TeacherID int
}

func main() {
	// 初期化
	var john Employee
	fmt.Println(john)

	employee := Employee{1001, "John", "Doe", "Doe's Street"}
	fmt.Println(employee)

	employee2 := Employee{LastName: "Doe", FirstName: "John"}

	// フィールドへのアクセス
	employee2.ID = 1001
	fmt.Println(employee2.FirstName)

	// 構造体へのポインター
	fmt.Println(employee2)
	employeeCopy := &employee2
	employeeCopy.FirstName = "Devid"
	fmt.Println(employee2)

	// 構造体の埋め込み
	var student Student
	student.Information.FirstName = "John"

	teacher := Teacher{
		Person: Person{
			FirstName: "John",
		},
	}
	teacher.LastName = "Due"
	fmt.Println(teacher.FirstName, teacher.LastName)

	teachers := []Teacher{
		Teacher{
			Person: Person{
				LastName: "Doe", FirstName: "John",
			},
		},
		Teacher{
			Person: Person{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	data, _ := json.Marshal(teachers)
	fmt.Printf("%s\n", data)

	var decoded []Teacher
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)
}
