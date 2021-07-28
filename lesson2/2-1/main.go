package main

func main() {
	firstName, lastName := "John", "Doe"
	age := 32
	println(firstName, lastName, age)

	const (
		StatusOK              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)
	println(StatusOK, StatusConnectionReset, StatusOtherError)
}
