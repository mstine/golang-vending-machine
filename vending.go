package vending

import "fmt"

const (
	NICKLE = iota
	DIME = iota
	QUARTER = iota
	DOLLAR = iota
)

type Item struct {
	price int
	count int
}

type Coin struct {
	key int
	value int
	label string
}

type VendingMachine struct {
	items map[string]*Item
	coins map[int]*Coin
	coinsInserted []*Coin
	bank map[*Coin]int
}

func NewVendingMachine() *VendingMachine {
	v := new(VendingMachine)

	v.items = make(map[string]*Item)
	v.items["A"] = &Item{65,0}
	v.items["B"] = &Item{100,0}
	v.items["C"] = &Item{150,0}

	v.coins = make(map[int]*Coin)
	v.coins[NICKLE] = &Coin{NICKLE, 5,"N"}
	v.coins[DIME] = &Coin{DIME, 10,"D"}
	v.coins[QUARTER] = &Coin{QUARTER, 25,"Q"}
	v.coins[DOLLAR] = &Coin{DOLLAR, 100,"DD"}

	v.coinsInserted = make([]*Coin, 0)

	v.bank = make(map[*Coin]int)
	v.bank[v.coins[NICKLE]] = 0
	v.bank[v.coins[DIME]] = 0
	v.bank[v.coins[QUARTER]] = 0
	v.bank[v.coins[DOLLAR]] = 0

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
			change, error := v.returnChange(changeDue)
			if error != nil {
				return "", error
			}    	
			return "A" + change, nil    	
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

	v.bank[v.coins[NICKLE]] = 50
	v.bank[v.coins[DIME]] = 50
	v.bank[v.coins[QUARTER]] = 50	
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

func (v *VendingMachine) Insert(coin int) {
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
		v.bank[v.coins[coin.key]]++
	}
	v.coinsInserted = make([]*Coin, 0)
}

func (v *VendingMachine) returnChange(changeDue int) (string, error) {
	coinReturn := ""

	CoinLoop:
	for i := QUARTER; i >= NICKLE; i-- {
		coins := changeDue / v.coins[i].value

		if v.bank[v.coins[i]] - coins < 0 {
			continue CoinLoop
		} else {
			v.bank[v.coins[i]] -= coins
		}
		
		for j := 1; j <= coins; j++ {
			coinReturn += ", " + v.coins[i].label		
		}
		
		changeDue -= coins * v.coins[i].value


		if changeDue == 0 {
			return coinReturn, nil
		}
	}	

	return "", fmt.Errorf("Can't make change!")
}
