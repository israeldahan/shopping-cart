package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// InitConfig initializes the config
func InitConfig() {
	viper.SetConfigName("default")
	viper.SetConfigType("json")
	viper.AddConfigPath(getConfigPath())
	viper.AutomaticEnv()
	viper.SetEnvPrefix("SC")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.ReadInConfig()
}

func getConfigPath() string {
	wd, _ := os.Getwd()
	return filepath.Join(wd, "configs")
}

// GetConfig returns the config
func GetConfig() *viper.Viper {
	return viper.GetViper()
}

type Product struct {
	Price    float64
	Discount string
}

// GetPriceAndDiscount returns the PriceAndDiscount
func GetPriceAndDiscount() (map[string]*Product, error) {
    priceAndDiscountMap := viper.Get("priceAndDiscount").(map[string]interface{})
	pruducts := make(map[string]*Product)
	for k := range priceAndDiscountMap {
		marshalled := priceAndDiscountMap[k].(map[string]interface{})
		// fmt.Println("marshalled", marshalled)
		price, ok := marshalled["price"].(float64)
		if !ok {
			return nil, errors.New("price is not a float64")
		}
		discount, ok := marshalled["discount"].(string)
		if !ok {
			return nil, errors.New("discount is not a float64")
		}
		pruducts[k] = &Product{
			Price:    price,
			Discount: discount,
		}
	}
    return pruducts, nil
}

func GetProductByName(product string) (*Product, error) {
	p := &Product{}
	priceAndDiscount, err := GetPriceAndDiscount()
	if err != nil {
		return nil, err
	}
	p, ok := priceAndDiscount[product]
	if !ok {
		return nil, errors.New("product not found")
	}
	return p, nil
}

func GetDiscountByProduct(product string) (string, error) {
	priceAndDiscount, err := GetPriceAndDiscount()
	if err != nil {
		return "", err
	}
	p, ok := priceAndDiscount[product]
	if !ok {
		return "", errors.New("product not found")
	}
	return p.Discount, nil
}

func GetPriceByProduct(product string) (float64, error) {
	priceAndDiscount, err := GetPriceAndDiscount()
	if err != nil {
		return 0, err
	}
	p, ok := priceAndDiscount[product]
	if !ok {
		return 0, errors.New("product not found")
	}
	return p.Price, nil
}