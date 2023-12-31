package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	data := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range data {
		result := CalculateTax(item.amount)

		if result != item.expected {
			t.Errorf("Expected %.2f but got %.2f", item.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1501.0}

	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Errorf("Expected 0 but got %.2f", result)

		}
	})
}

// Testing using testify
func TestCalculateTaxTestify(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result := CalculateTax(amount)
	assert.Equal(t, result, expected)
}

// Testing using mocks
func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TextRepositoryMock{}
	repository.On("Save", 10.0).Return(nil)
	repository.On("Save", 0.0).Return(errors.New("error"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0, repository)
	assert.Error(t, err)

	repository.AssertExpectations(t)
}
