package services

import (
	"log"
	"strconv"
	"strings"
)

func GetDiscountDetails(discount string) (float64, string, float64, error) {
	// discount = "3 for 2"
	// discount = "2 + 1"
	
	discountDetails := strings.Split(discount, " ")
	quantity := discountDetails[0]
	discountType := discountDetails[1]
	sale := discountDetails[2]
	return convertStringToFloat(quantity), discountType, convertStringToFloat(sale), nil
}

func convertStringToFloat(str string) float64 {
    f, err := strconv.ParseFloat(str, 64)
    if err != nil {
        log.Fatal(err)
    }
    return f
}
