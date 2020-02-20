package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	var builder strings.Builder

	for _, cat := range listCat {
		var count int

		for _, art := range listArt {
			if art[0] == cat[0] {
				artCount, _ := strconv.Atoi(strings.Split(art, " ")[1])
				count += artCount
			}
		}

		builder.WriteString("(" + cat + " : " + strconv.Itoa(count) + ") - ")
	}

	return builder.String()[:builder.Len()-3]
}

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}

	fmt.Println(StockList(b, c))
}
