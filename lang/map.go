package main

import "fmt"

type user struct {
	name    string
	surname string
}

func main() {

	users1 := make(map[string]user)

	users1["Roy"] = user{"Rob", "Roy"}
	users1["Ford"] = user{"Henry", "Ford"}
	users1["Mouse"] = user{"Mickey", "Mouse"}
	users1["Jackson"] = user{"Michael", "Jackson"}

	fmt.Printf("\n=>Iterate over map\n")
	for key, value := range users1 {
		fmt.Println(key, value)
	}

	users2 := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	fmt.Printf("\n=>Map literals\n")
	for key, value := range users2 {
		fmt.Println(key, value)
	}

	delete(users2, "Roy")

	u1, found1 := users2["Roy"]
	u2, found2 := users2["Ford"]

	fmt.Printf("\n=>Find key\n")
	fmt.Println("Roy", found1, u1)
	fmt.Println("Ford", found2, u2)

	
}
