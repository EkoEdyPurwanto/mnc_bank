package usecase

import (
	"EkoEdyPurwanto/mnc-bank/model"
	"EkoEdyPurwanto/mnc-bank/model/req"
	"EkoEdyPurwanto/mnc-bank/utils/common"
	"EkoEdyPurwanto/mnc-bank/utils/security"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
)

type CustomerUseCase interface {
	Register(payload req.RegisterRequest) error
	Login(payload req.LoginRequest) (string, error)
	Logout() error
	Payment() error
}

type customerUseCase struct {
	Customers    []*model.Customer
	dataFilePath string
}

func NewCustomerUseCase(dataFilePath string) (CustomerUseCase, error) {
	data, err := os.ReadFile(dataFilePath)
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

func (uc *customerUseCase) Register(payload req.RegisterRequest) error {
	// Validate the payload
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return err
	}

	// hash password
	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return err
	}
	// Create a new customer with a unique ID
	newCustomer := &model.Customer{
		ID:       common.GenerateID(),
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashPassword,
	}

	// Check if the email is already registered
	for _, customer := range uc.Customers {
		if customer.Email == payload.Email {
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
	customerData := model.CustomerData{
		Customers: uc.Customers,
	}

	data, err := json.MarshalIndent(customerData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal customer data: %v", err)
	}

	err = os.WriteFile(uc.dataFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write customer data to file: %v", err)
	}

	return nil
}

func (uc *customerUseCase) Login(payload req.LoginRequest) (string, error) {
	// Validate the payload
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return "", err
	}

	var identifier string

	if payload.Identifier.Name != "" {
		identifier = payload.Identifier.Name
	} else if payload.Identifier.Email != "" {
		identifier = payload.Identifier.Email
	}

	if identifier == "" {
		return "", errors.New("unauthorized: Invalid credential")
	}

	// Cari customer yang sesuai dengan email atau nama
	var matchedCustomer *model.Customer
	for _, customer := range uc.Customers {
		if customer.Email == payload.Identifier.Email || customer.Name == payload.Identifier.Name {
			matchedCustomer = customer
			break
		}
	}

	// Validasi Password
	err = security.VerifyPassword(matchedCustomer.Password, payload.Password)
	if err != nil {
		return "", fmt.Errorf("unauthorized: invalid credential")
	}

	// Generate Token
	token, err := security.GenerateJWTToken(matchedCustomer)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *customerUseCase) Logout() error {
	// Implement logout logic here

	return nil
}

func (uc *customerUseCase) Payment() error {
	// Implement payment logic here
	panic("implement me")
}
