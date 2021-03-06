package Database

import (
	"strconv"
	"time"
)

func BuyCart(ShoppingCart []Sale) []Sale {
	for _, v := range ShoppingCart {
		y, month, day := time.Now().Date()
		year, _ := strconv.Atoi(strconv.Itoa(y)[1:])
		v.Day = uint8(day)
		v.Month = uint8(month)
		v.Year = uint8(year)
		Reports[0] = append(Reports[0], v)
		Items[ItemKeys[v.ID].Idxes[0]].Quantity -= v.Quantity
	}
	SaveData()
	return ShoppingCart[:0]
}

func AddToCart(item Sale, ShoppingCart []Sale) []Sale {
	for i, v := range ShoppingCart {
		if v.ID == item.ID && v.Price == item.Price {
			ShoppingCart[i].Quantity++
			return ShoppingCart
		}
	}
	return append(ShoppingCart, item)
}

func DecreaseFromCart(item Sale, ShoppingCart []Sale) []Sale {
	for i, v := range ShoppingCart {
		if v.ID != item.ID || v.Price != item.Price {
			continue
		}

		if v.Quantity-1 > 0 {
			ShoppingCart[i].Quantity -= 1
		} else {
			ShoppingCart[i] = ShoppingCart[len(ShoppingCart)-1] // Copy last element to index i.
			ShoppingCart = ShoppingCart[:len(ShoppingCart)-1]   // Truncate slice.
		}
	}

	return ShoppingCart
}

func GetCartTotal(ShoppingCart []Sale) (total float32) {
	for _, v := range ShoppingCart {
		total += v.Price * float32(v.Quantity)
	}
	return
}

func ConvertCart(shoppingCart []Sale) (intercart []interface{}) {
	for i := range shoppingCart {
		intercart = append(intercart, shoppingCart[i])
	}
	return
}

func ConvertItemKeys() (inter []int) {
	for k := range ItemKeys {
		inter = append(inter, int(k))
	}
	return
}

func ConvertItemIdxes(target uint64) (list []int) {
	for _, v := range ItemKeys[target].Idxes {
		list = append(list, v)
	}
	return
}

func RemoveItem(idx int, id uint64) {
	Free_Spaces = append(Free_Spaces, ItemKeys[id].Idxes[idx])
	ItemKeys[id].Idxes[idx] = ItemKeys[id].Idxes[len(ItemKeys[id].Idxes)-1]
	ItemKeys[id].Idxes = ItemKeys[id].Idxes[:len(ItemKeys[id].Idxes)-1]
}

func ConvertExpenses() (inter []interface{}) {
	for i := range Expenses {
		inter = append(inter, Expenses[i])
	}
	return
}

func RemoveExpense(index int) {
	Expenses[index] = Expenses[len(Expenses)-1]
	Expenses = Expenses[:len(Expenses)-1]
}

func ConvertString(Price, Cost, Quantity string) (float32, float32, uint16) {
	newPrice, _ := strconv.ParseFloat(Price, 64)
	newCost, _ := strconv.ParseFloat(Cost, 64)
	newQuantity, _ := strconv.Atoi(Quantity)
	return float32(newPrice), float32(newCost), uint16(newQuantity)
}

func ConvertItem(id uint64) (result Sale) {
	vals := ItemKeys[id]

	result.ID = id
	result.Price = vals.Price
	result.Cost = Items[vals.Idxes[0]].Cost
	result.Quantity = 1
	return
}
