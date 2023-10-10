package tax

import "github.com/stretchr/testify/mock"

type Repository interface {
	Save(amount float64) error
}

type TextRepositoryMock struct {
	mock.Mock
}

func (m *TextRepositoryMock) Save(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax(amount)
	return repository.Save(tax)
}

func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}
