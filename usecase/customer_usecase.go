package usecase

import (
	"EkoEdyPurwanto/mnc-bank/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CustomerUseCase interface {
	Register(name, email, password string) error
	Login(email, password string) error
	Logout() error
	Payment() error
}

type customerUseCase struct {
	Customers    []*model.Customer
	dataFilePath string
}

func NewCustomerUseCase(dataFilePath string) (CustomerUseCase, error) {
	data, err := ioutil.ReadFile(dataFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from JSON file: %v", err)
	}

	var customerData model.CustomerData
	if err := json.Unmarshal(data, &customerData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}

	return &customerUseCase{
		Customers:    customerData.Customers,
		dataFilePath: dataFilePath,
	}, nil
}

func (uc *customerUseCase) Register(name, email, password string) error {
	// Create a new customer with a unique ID
	newCustomer := &model.Customer{
		ID:       len(uc.Customers) + 1, // You can use a more robust ID generation method
		Name:     name,
		Email:    email,
		Password: password,
	}

	// Check if the email is already registered
	for _, customer := range uc.Customers {
		if customer.Email == email {
			return fmt.Errorf("email is already registered")
		}
	}

	// Append the new customer to the list of customers
	uc.Customers = append(uc.Customers, newCustomer)

	// Save the updated customer data to the JSON file
	if err := uc.saveCustomerDataToFile(); err != nil {
		return err
	}

	return nil
}

func (uc *customerUseCase) saveCustomerDataToFile() error {
	data, err := json.Marshal(uc.Customers)
	if err != nil {
		return fmt.Errorf("failed to marshal customer data: %v", err)
	}

	err = ioutil.WriteFile(uc.dataFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write customer data to file: %v", err)
	}

	return nil
}

func (uc *customerUseCase) Login(email, password string) error {
	// Implement login logic here
	// Check if the email and password match any customer's credentials

	return nil
}

func (uc *customerUseCase) Logout() error {
	// Implement logout logic here

	return nil
}

func (uc *customerUseCase) Payment() error {
	// Implement payment logic here
	panic("implement me")
}
