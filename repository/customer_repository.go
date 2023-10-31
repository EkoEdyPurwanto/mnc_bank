package repository

import (
	"EkoEdyPurwanto/mnc-bank/model"
	"encoding/json"
	"fmt"
	"os"
)

type CustomerRepository interface {
	Save(customers []*model.Customer) error
	Load() ([]*model.Customer, error)
}

type customerRepository struct {
	customer []*model.Customer
}

func NewCustomerRepository(customer []*model.Customer) CustomerRepository {
	return &customerRepository{
		customer: customer,
	}
}

func (c *customerRepository) Save(customers []*model.Customer) error {
	customerData := model.CustomerData{
		Customers: customers,
	}

	data, err := json.MarshalIndent(customerData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal customer data: %v", err)
	}

	err = os.WriteFile("../data/json/customer.json", data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write customer data to file: %v", err)
	}

	return nil
}

func (c *customerRepository) Load() ([]*model.Customer, error) {
	data, err := os.ReadFile("../data/json/customer.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read customer data from file: %v", err)
	}

	var customerData model.CustomerData
	if err := json.Unmarshal(data, &customerData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}

	return customerData.Customers, nil
}
