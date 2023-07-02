package services

import (
	"math"
)

func CalculateBillForDiscountTypeFor(count float64, price float64, discountQuantity float64, discountPrice float64) (float64, error) {
	if count < discountQuantity {
		return count * price, nil
	} else if math.Remainder(count, discountQuantity) == 0 {
		amount := count / discountQuantity
		return amount * discountPrice, nil
	} else {
		outOfDiscount := math.Remainder(count, discountQuantity)
		amount := (count - outOfDiscount) / discountQuantity
		return amount*discountPrice + outOfDiscount*price, nil
	}
}

func CalculateBillForDiscountTypePlus(count float64, price float64, discountQuantity float64, discountPlus float64) (float64, error) {
	totalDiscount := discountQuantity + discountPlus
	discountPrice := price - (price * (discountPlus / totalDiscount))
	if count < totalDiscount {
		return count * price, nil
	} else if math.Remainder(count, totalDiscount) == 0 {
		return count * discountPrice, nil
	} else {
		outOfDiscount := math.Remainder(count, discountQuantity)
		amaunt := (count - outOfDiscount)
		return amaunt*discountPrice + outOfDiscount*price, nil
	}
}

func CalculateBillForDiscountTypeNone(count float64, price float64) float64 {
	return count * price
}

func CalculateTotalBill(billByProduct map[string]float64) (float64, error) {
	var total float64
	for _, v := range billByProduct {
		total += v
	}
	return total, nil
}
