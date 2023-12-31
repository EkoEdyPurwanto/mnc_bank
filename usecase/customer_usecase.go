package usecase

import (
	"EkoEdyPurwanto/mnc-bank/model"
	"EkoEdyPurwanto/mnc-bank/model/req"
	"EkoEdyPurwanto/mnc-bank/repository"
	"EkoEdyPurwanto/mnc-bank/utils/common"
	"EkoEdyPurwanto/mnc-bank/utils/security"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CustomerUseCase interface {
	Register(payload req.RegisterRequest) error
	Login(payload req.LoginRequest) (string, error)
	Logout() error
	Payment() error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		repo: repo,
	}
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

	// Baca data pelanggan dari file JSON
	customers, err := uc.repo.Load()
	if err != nil {
		return err
	}

	// Check if the email is already registered
	for _, customer := range customers {
		if customer.Email == newCustomer.Email {
			return errors.New("email is already registered")
		}
	}

	// Append the new customer to the list of customers
	customers = append(customers, newCustomer)

	// Save the updated customer data to the JSON file
	if err := uc.repo.Save(customers); err != nil {
		return fmt.Errorf("failed to save customer data: %v", err)
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

	// Baca data pelanggan dari file JSON
	customers, err := uc.repo.Load()
	if err != nil {
		return "", err
	}

	// Cari customer yang sesuai dengan email atau nama
	var matchedCustomer *model.Customer
	for _, customer := range customers {
		if customer.Email == payload.Identifier.Email || customer.Name == payload.Identifier.Name {
			matchedCustomer = customer
			break
		}
	}

	if matchedCustomer == nil {
		return "", errors.New("unauthorized: Customer not found")
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

	// Save the token data to the JSON file yang sudah terbuat tadi,

	return token, nil
}

func (uc *customerUseCase) Logout() error {
	// Implement payment logic here
	panic("implement me")
}

func (uc *customerUseCase) Payment() error {
	// Implement payment logic here
	panic("implement me")
}
