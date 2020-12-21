// package main

// import "fmt"

// type Skills []string

// type Human struct {
// 	name   string
// 	age    int
// 	weight int
// }

// type Student struct {
// 	Human
// 	Skills
// 	int
// 	speciality string
// }

// func main() {
// 	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}

// 	fmt.Println("Her name is", jane.name)
// 	fmt.Println("Her age is", jane.age)
// 	fmt.Println("Her weight is", jane.weight)
// 	fmt.Println("Her speciality is", jane.speciality)

// 	jane.Skills = []string{"anatomy"}
// 	fmt.Println("Her skills is", jane.Skills)
// 	fmt.Println("She acquired two new ones")
// 	jane.Skills = append(jane.Skills, "Physics", "golang")
// 	fmt.Println("Her skills now are", jane.Skills)

// 	jane.int = 3
// 	fmt.Println("Her preferred number is", jane.int)

// }

package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string // Human型がもつフィールド
}

type Employee struct {
	Human      // 匿名フィールドHuman
	speciality string
	phone      string // 社員のphoneフィールド
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// もし我々がHumanのphoneフィールドにアクセスする場合は
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}
