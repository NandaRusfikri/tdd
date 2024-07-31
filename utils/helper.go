package utils

import "errors"

func CalculateTotalPrice(originalPrice, discountPercentage int) (int, error) {
	if originalPrice < 0 {
		return 0, errors.New("original price must be non-negative")
	}
	if discountPercentage < 0 || discountPercentage > 100 {
		return 0, errors.New("discount percentage must be between 0 and 100")
	}

	discountAmount := originalPrice * discountPercentage / 100
	totalPrice := originalPrice - discountAmount
	return totalPrice, nil
}
