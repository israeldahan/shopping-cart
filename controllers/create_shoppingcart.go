package controllers

import (
	"context"
	// "fmt"
	"net/http"

	"go.uber.org/zap"
	"shoppingcart.com/config"
	"shoppingcart.com/handlers"
	"shoppingcart.com/logger"
	"shoppingcart.com/models"
	"shoppingcart.com/services"
)

// BillReq is the request to create a Shoppingcart
type BillReq struct {
	Cart []*models.BillRequest `json:"cart"`
}

// BillResp is the response to create a Shoppingcart
type BillResp struct {
	Cart *models.BillResponse `json:"cart"`
}

// CreateBill handles the request to create a calculation
func CreateBill(c context.Context, urlData handlers.URLData, reqShoppingcart BillReq) (int, interface{}) {

	logger.Info("start to create bill")

	cart := reqShoppingcart.Cart

	// Count the number of each product
	cartUniqueProductsCount := make(map[string]float64)
	for _, product := range cart {
		cartUniqueProductsCount[product.Name]++
	}
	logger.Info("End to calculate the count of each product by getting all products", zap.Any("cartUniqueProductsCount", cartUniqueProductsCount))

	logger.Info("Start to calculate the bill by getting all products")
	billByProduct := make(map[string]float64)
	for k, count := range cartUniqueProductsCount {
		discount, _ := config.GetDiscountByProduct(k)
		price, _ := config.GetPriceByProduct(k)
		discountQuantity, discountType, discountPrice, err := services.GetDiscountDetails(discount)
		if err != nil {
			logger.Error("error to get discount details", err, zap.Any("discount", discount))
		}
		if discountType == "for" {
			billByProduct[k], err = services.CalculateBillForDiscountTypeFor(count, price, discountQuantity, discountPrice)
			if err != nil {
				logger.Error("error to calculate bill for discount type for", err, zap.Any("count", count), zap.Any("price", price), zap.Any("discountQuantity", discountQuantity), zap.Any("discountPrice", discountPrice))
			}
		} else if discountType == "+" {
			billByProduct[k], err = services.CalculateBillForDiscountTypePlus(count, price, discountQuantity, discountPrice)
			if err != nil {
				logger.Error("error to calculate bill for discount type plus", err, zap.Any("count", count), zap.Any("price", price), zap.Any("discountQuantity", discountQuantity), zap.Any("discountPrice", discountPrice))
			}
		} else {
			billByProduct[k] = services.CalculateBillForDiscountTypeNone(count, price)
		}

	}
	total, err := services.CalculateTotalBill(billByProduct)
	if err != nil {
		logger.Error("error to calculate total bill", err, zap.Any("billByProduct", billByProduct))
	}
	// fmt.Println("cartUniqueProductsCount", cartUniqueProductsCount)

	return http.StatusOK, BillResp{
		Cart: &models.BillResponse{
			Bill: total,
		},
	}
}
