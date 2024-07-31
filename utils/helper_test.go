package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//TestCalculateTotalPrice_ValidInputs_NoDiscount
//TestCalculateTotalPrice_ValidInputs_FullDiscount
//TestCalculateTotalPrice_ValidInputs_HalfDiscount
//TestCalculateTotalPrice_NegativeOriginalPrice
//TestCalculateTotalPrice_DiscountPercentageBelowZero
//TestCalculateTotalPrice_DiscountPercentageAboveOneHundred
//TestCalculateTotalPrice_ZeroOriginalPrice
//TestCalculateTotalPrice_ZeroDiscountPercentage

func TestCalculateTotalPrice(t *testing.T) {

	tests := []struct {
		name          string
		inputPrice    int
		inputDiscount int
		expectedPrice int
		expectedErr   error
	}{
		{
			name:          "No Discount",
			inputPrice:    20000,
			inputDiscount: 0,
			expectedPrice: 20000,
			expectedErr:   nil,
		},
		{
			name:          "No Discount",
			inputPrice:    20000,
			inputDiscount: 0,
			expectedPrice: 20000,
			expectedErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, err := CalculateTotalPrice(tt.inputPrice, tt.inputDiscount)

			assert.Equal(t, tt.expectedPrice, res)
			assert.Equal(t, tt.expectedErr, err)
		})
	}

}
