package Data

import (
	"strconv"
	"strings"
)

func GetTotalProfit(selectionStr string) (revenue, cost, profit float64){
	targetSheet := "Report Data"
	totalRevenue := 0.0
	totalCost := 0.0

	results := FindAll(targetSheet, "G", selectionStr)

	for	_, r := range results{
		rev := f.GetCellValue(targetSheet, "D" + strconv.Itoa(r))
		cos := f.GetCellValue(targetSheet, "E" + strconv.Itoa(r))
		conRev, _ :=  strconv.ParseFloat(rev, 64)
		conCos, _ := strconv.ParseFloat(cos, 64)

		totalRevenue += conRev
		totalCost += conCos
	}

	totalProfit := totalRevenue - totalCost
	return totalRevenue, totalCost, totalProfit
}

func FindAll(targetSheet, targetAxis, subStr string) []int {
	var idxes []int
	cell := f.GetCellValue(targetSheet, targetAxis + "1")
	for i := 1; cell != "";  {
		if strings.Contains(cell, subStr) {
			idxes = append(idxes, i)
		}
		i++
		cell = f.GetCellValue(targetSheet, targetAxis+strconv.Itoa(i))
	}
	return idxes
}

func GetData(targetSheet string,id int) []string{
	i := GetIndex(targetSheet, id, 1)

	name := f.GetCellValue(targetSheet, "B"+ strconv.Itoa(i))
	price := f.GetCellValue(targetSheet, "C"+ strconv.Itoa(i))
	cost := f.GetCellValue(targetSheet, "D"+ strconv.Itoa(i))
	quantity := f.GetCellValue(targetSheet, "E"+ strconv.Itoa(i))
	return []string{
		name,
		price,
		cost,
		quantity,
	}
}

func GetAllData(targetSheet string) []Sale{
	i:= 1
	var data []Sale

	cell := f.GetCellValue(targetSheet, "A"+strconv.Itoa(i))
	for cell != ""{
		name := f.GetCellValue(targetSheet, "B"+ strconv.Itoa(i))
		price := f.GetCellValue(targetSheet, "C"+ strconv.Itoa(i))
		cost := f.GetCellValue(targetSheet, "D"+ strconv.Itoa(i))
		quantity := f.GetCellValue(targetSheet, "E"+ strconv.Itoa(i))

		conID, _ := strconv.Atoi(cell)
		p, c, q := ConvertStringToSale(price, cost, quantity)

		temp := Sale{
			ID:    conID,
			Name:  name,
			Price: p,
			Cost:  c,
			Quantity: q,
		}
		data = append(data, temp)
		cell = f.GetCellValue(targetSheet, "A"+strconv.Itoa(i+1))
		i++
	}
	return data
}