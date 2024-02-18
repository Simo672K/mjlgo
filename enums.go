package main

// somthing like enums like in go

type GameCategory int

const (
	Action GameCategory = iota // assign 1 to Action and auto increment and assign other constants. Other solution is to asign each value manuly
	Adventure
	SciFi
	Horror
)

func getGameCategory(gameCategory GameCategory) int {
	switch gameCategory {
	case Action:
		return 1
	case Adventure:
		return 2
	case SciFi:
		return 3
	case Horror:
		return 4
	default:
		panic("Category provided doesn't exist")
	}
}

// func main() {
// 	fmt.Println(getGameCategory(Adventure))
// }
