package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var n food
	if order == "burger" {
		n.preptime = 15
	} else if order == "chips" {
		n.preptime = 10
	} else if order == "nuggets" {
		n.preptime = 12
	} else {
		return 404
	}
	return n.preptime
}
