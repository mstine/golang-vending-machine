package vending

import "fmt"

type Item struct {
	price int
	count int
}

type Coin struct {
	value int
	label string
}

type VendingMachine struct {
	items map[string]*Item
	coins map[string]*Coin
	coinsInserted []*Coin
}

func NewVendingMachine() *VendingMachine {
	v := new(VendingMachine)

	v.items = make(map[string]*Item)
	v.items["A"] = &Item{65,0}
	v.items["B"] = &Item{100,0}
	v.items["C"] = &Item{150,0}

	v.coins = make(map[string]*Coin)
	v.coins["N"] = &Coin{5,"N"}
	v.coins["D"] = &Coin{10,"D"}
	v.coins["Q"] = &Coin{25,"Q"}
	v.coins["DD"] = &Coin{100,"DD"}

	v.coinsInserted = make([]*Coin, 0)

	return v
}

func (v *VendingMachine) Get(item string) (string, error) {
	if v.items[item].count > 0 {
		if v.AmountInserted() == v.items[item].price {
			v.items[item].count--
			v.coinsInserted = make([]*Coin, 0)
			return item, nil
		} else {
			return "", fmt.Errorf("You didn't insert enough money for %v! Inserted: %v, Required: %v", item, v.AmountInserted(), v.items[item].price)
		}
	} else {
		return "", fmt.Errorf("No %v items available!", item)
	}
}

func (v *VendingMachine) Service() {
	v.items["A"].count = 50
	v.items["B"].count = 50
	v.items["C"].count = 50	
}

func (v *VendingMachine) CoinReturn() string {
	coinReturn := "";
	for i, coin := range v.coinsInserted {
		coinReturn += coin.label
		if i < len(v.coinsInserted) - 1 {
			coinReturn += ", "
		}
	}	
	return coinReturn
}

func (v *VendingMachine) Insert(coin string) {
	v.coinsInserted = append(v.coinsInserted, v.coins[coin])
}

func (v *VendingMachine) AmountInserted() int {
	amount := 0
	for _, coin := range v.coinsInserted {
		amount += coin.value
	}	
	return amount
}