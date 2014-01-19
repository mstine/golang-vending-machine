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
	bank map[*Coin]int
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

	v.bank = make(map[*Coin]int)
	v.bank[v.coins["N"]] = 0
	v.bank[v.coins["D"]] = 0
	v.bank[v.coins["Q"]] = 0
	v.bank[v.coins["DD"]] = 0

	return v
}

func (v *VendingMachine) Get(item string) (string, error) {
	if v.items[item].count > 0 {
		if v.AmountInserted() == v.items[item].price {
			v.items[item].count--
			v.addAmountInsertedToBank()
			return item, nil
	    } else if v.AmountInserted() > v.items[item].price {
	    	changeDue := v.AmountInserted() - v.items[item].price
	    	v.addAmountInsertedToBank()
			v.items[item].count--	    	
			return "A" + v.returnChange(changeDue), nil    	
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

	v.bank[v.coins["N"]] = 50
	v.bank[v.coins["D"]] = 50
	v.bank[v.coins["Q"]] = 50	
}

func (v *VendingMachine) CoinReturn() string {
	coinReturn := ""
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

func (v *VendingMachine) addAmountInsertedToBank() {
	for _, coin := range v.coinsInserted {
		v.bank[v.coins[coin.label]]++
	}
	v.coinsInserted = make([]*Coin, 0)
}

func (v *VendingMachine) returnChange(changeDue int) string {
	coinReturn := ""

	quarters := changeDue / v.coins["Q"].value
	for i := 1; i <= quarters; i++ {
		coinReturn += ", Q"		
	}

	changeDue -= quarters * v.coins["Q"].value

	if changeDue == 0 {
		return coinReturn
	}

	dimes := changeDue / v.coins["D"].value
	for i := 1; i <= dimes; i++ {
		coinReturn += ", D"		
	}

	changeDue -= dimes * v.coins["D"].value

	if changeDue == 0 {
		return coinReturn
	}

	return "ERROR"
}
