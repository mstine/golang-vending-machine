package vending

import "container/list"

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
	coinsInserted *list.List
}

type VendingMachineError struct {
	Message string
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

	v.coinsInserted = list.New()

	return v
}

func (e *VendingMachineError) Error() string {
	return e.Message
}

func (v *VendingMachine) Get(item string) (string, error) {
	if v.items[item].count > 0 {
		v.items[item].count--
		return item, nil
	} else {
		return "", &VendingMachineError{"No " + item + " items available!"}
	}
}

func (v *VendingMachine) Service() {
	v.items["A"].count = 50
	v.items["B"].count = 50
	v.items["C"].count = 50	
}

func (v *VendingMachine) CoinReturn() string {
	e := v.coinsInserted.Front();
	coinReturn := "";
	if e != nil {
		coinReturn += e.Value.(*Coin).label
		for e != nil {
			e = e.Next()
			if e != nil {
				coinReturn += ", "	
				coinReturn += e.Value.(*Coin).label	
			}
		}		
		v.coinsInserted = list.New()
	}
	return coinReturn
}

func (v *VendingMachine) Insert(coin string) {
	v.coinsInserted.PushBack(v.coins[coin])
}

func (v *VendingMachine) AmountInserted() int {
	amount := 0
	for e := v.coinsInserted.Front(); e != nil; e = e.Next() {
		amount += e.Value.(*Coin).value
	}
	return amount
}