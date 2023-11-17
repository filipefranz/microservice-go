package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name, emali string) (*Client, error) {
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     emali,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := client.Validate()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (client *Client) Validate() error {
	if client.Name == "" {
		return errors.New("name is required")
	}
	if client.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

func (client *Client) Update(name, email string) error {
	client.Name = name
	client.Email = email
	client.UpdatedAt = time.Now()
	err := client.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) AddAccount(account *Account) error {
	if account.Client.ID != client.ID {
		return errors.New("account does not belong to client")
	}
	client.Accounts = append(client.Accounts, account)
	return nil
}
