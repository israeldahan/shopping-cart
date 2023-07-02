package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"shoppingcart.com/services"
)

func TestCalculateBillForDiscountTypeFor(t *testing.T) {

	t.Run("should return the right amount for discount type for", func(t *testing.T) {
		count := 2.0
		price := 10.0
		discountQuantity := 3.0
		discountPrice := 25.0
		expected := 20.0
		actual, err := services.CalculateBillForDiscountTypeFor(count, price, discountQuantity, discountPrice)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should return the right amount for discount type for", func(t *testing.T) {
		count := 5.0
		price := 10.0
		discountQuantity := 3.0
		discountPrice := 25.0
		expected := 45.0
		actual, err := services.CalculateBillForDiscountTypeFor(count, price, discountQuantity, discountPrice)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should return the right amount for discount type for", func(t *testing.T) {
		count := 6.0
		price := 10.0
		discountQuantity := 3.0
		discountPrice := 25.0
		expected := 50.0
		actual, err := services.CalculateBillForDiscountTypeFor(count, price, discountQuantity, discountPrice)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should return the right amount for discount type Plus", func(t *testing.T) {
		count := 4.0
		price := 10.0
		discountQuantity := 3.0
		discountPlus := 1.0
		expected := 30.0
		actual, err := services.CalculateBillForDiscountTypePlus(count, price, discountQuantity, discountPlus)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should return the right amount for discount type Plus", func(t *testing.T) {
		count := 5.0
		price := 10.0
		discountQuantity := 3.0
		discountPlus := 1.0
		expected := 40.0
		actual, err := services.CalculateBillForDiscountTypePlus(count, price, discountQuantity, discountPlus)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("should return the right amount for discount type Plus", func(t *testing.T) {
		count := 2.0
		price := 10.0
		discountQuantity := 3.0
		discountPlus := 1.0
		expected := 20.0
		actual, err := services.CalculateBillForDiscountTypePlus(count, price, discountQuantity, discountPlus)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

}