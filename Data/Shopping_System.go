package Data

import (
	"fmt"
	"strconv"
)

type Sale struct {
	ID       int
	Quantity int
	Price    float64
	Cost     float64
	Name     string
}

func BuyCart(ShoppingCart []Sale) []Sale {
	targetSheet := "Report Data"

	for _, v := range ShoppingCart {
		UpdateData(v, targetSheet, 1)
	}

	ReadVal(targetSheet)
	SaveFile()
	return ClearCart(ShoppingCart)
}

func AddToCart(item Sale, ShoppingCart []Sale) []Sale {
	for i, v := range ShoppingCart {
		if v.ID == item.ID && v.Price == item.Price {
			ShoppingCart[i].Quantity += 1
			return ShoppingCart
		}
	}

	fmt.Println(item)
	ShoppingCart = append(ShoppingCart, item)
	return ShoppingCart
}

func DecreaseFromCart(item Sale, ShoppingCart []Sale) []Sale {
	for i, v := range ShoppingCart {
		if v.ID == item.ID && v.Price == item.Price {
			if v.Quantity-1 > 0 {
				ShoppingCart[i].Quantity -= 1
			} else {
				ShoppingCart = RemoveFromCart(i, ShoppingCart)
			}
		}
	}

	return ShoppingCart
}

func RemoveFromCart(i int, ShoppingCart []Sale) []Sale {
	ShoppingCart[i] = ShoppingCart[len(ShoppingCart)-1] // Copy last element to index i.
	ShoppingCart[len(ShoppingCart)-1] = Sale{}          // Erase last element (write zero value).
	ShoppingCart = ShoppingCart[:len(ShoppingCart)-1]   // Truncate slice.
	return ShoppingCart
}

func GetCartTotal(ShoppingCart []Sale) float64 {
	total := 0.0
	for _, v := range ShoppingCart {
		total += v.Price * float64(v.Quantity)
	}
	return total
}

//Removes all items from the shopping cart
func ClearCart(ShoppingCart []Sale) []Sale {
	ShoppingCart = ShoppingCart[:0]
	return ShoppingCart
}

func ConvertStringToSale(Price, Cost, Quantity string) (float64, float64, int) {
	newPrice, _ := strconv.ParseFloat(Price, 64)
	newCost, _ := strconv.ParseFloat(Cost, 64)
	newQuantity, _ := strconv.Atoi(Quantity)
	return newPrice, newCost, newQuantity
}

func ConvertString(price, cost float64, inventory int) []string {
	p := fmt.Sprint(price)
	c := fmt.Sprint(cost)
	inven := strconv.Itoa(inventory)

	return []string{
		p,
		c,
		inven,
	}
}
