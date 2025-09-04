package tools

import (
	"time"
)

type mockDB struct{}

// Mock de logins
var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "Alex",
	},
	"maria": {
		AuthToken: "456DEF",
		Username:  "Maria",
	},
	"joao": {
		AuthToken: "789GHI",
		Username:  "João",
	},
	"ana": {
		AuthToken: "321JKL",
		Username:  "Ana",
	},
	"carlos": {
		AuthToken: "654MNO",
		Username:  "Carlos",
	},
}

// Mock de saldo de moedas
var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"maria": {
		Coins:    250,
		Username: "maria",
	},
	"joao": {
		Coins:    50,
		Username: "joao",
	},
	"ana": {
		Coins:    400,
		Username: "ana",
	},
	"carlos": {
		Coins:    0,
		Username: "carlos",
	},
}

// GetUserLoginDetails simula buscar login de um usuário
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	client, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &client
}

// GetUserCoins simula buscar saldo de moedas de um usuário
func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	client, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &client
}

// SetupDatabase apenas simula setup do banco
func (d *mockDB) SetupDatabase() error {
	return nil
}
