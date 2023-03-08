package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_It_Gets_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Blank_Or_Negative(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "invalid price")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Blank_Or_Negative(t *testing.T) {
	order := Order{ID: "123", Price: 99.0}
	assert.Error(t, order.Validate(), "invalid id")
}

func Test_With_All_Valid_Params(t *testing.T) {
	order := Order{ID: "123", Price: 99.0, Tax: 9.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 99.0, order.Price)
	assert.Equal(t, 9.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 108.0, order.FinalPrice)
}
