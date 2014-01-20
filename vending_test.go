package vending

import "testing"

func TestBuyFromAEmptyMachine(t *testing.T) {
	v := NewVendingMachine()
	v.Insert(QUARTER)
	v.Insert(QUARTER)
	v.Insert(DIME)
	v.Insert(NICKLE)
	result, error := v.Get("A")
	
	if result == "A" {
		t.Errorf("Empty machine should not return an A!")
	}

	if error.Error() != "No A items available!" {
		t.Errorf("Wrong error message!")
	}
}

func TestBuyFromBEmptyMachine(t *testing.T) {
	v := NewVendingMachine()
	v.Insert(DOLLAR)
	result, error := v.Get("B")
	
	if result == "B" {
		t.Errorf("Empty machine should not return a B")
	}

	if error.Error() != "No B items available!" {
		t.Errorf("Wrong error message!")
	}
}

func TestBuyFromCEmptyMachine(t *testing.T) {
	v := NewVendingMachine()
	v.Insert(DOLLAR)
	v.Insert(QUARTER)
	v.Insert(QUARTER)
	result, error := v.Get("C")
	
	if result == "C" {
		t.Errorf("Empty machine should not return a C")
	}

	if error.Error() != "No C items available!" {
		t.Errorf("Wrong error message!")
	}
}

func TestServiceMachine(t *testing.T) {
	v := NewVendingMachine()
	v.Service()

	if v.items["A"].count != 50 || v.items["B"].count != 50 || v.items["C"].count != 50 {
		t.Errorf("Machine was not properly serviced!")
	}
}

func TestBuyAWithNoMoney(t *testing.T) {
	v := NewVendingMachine()
	v.Service()
	_, error := v.Get("A")
	if error.Error() != "You didn't insert enough money for A! Inserted: 0, Required: 65" {
		t.Errorf("Improper error message!")
	}	
}

func TestBuyBWithNoMoney(t *testing.T) {
	v := NewVendingMachine()
	v.Service()
	_, error := v.Get("B")
	if error.Error() != "You didn't insert enough money for B! Inserted: 0, Required: 100" {
		t.Errorf("Improper error message!")
	}	
}

func TestBuyCWithNoMoney(t *testing.T) {
	v := NewVendingMachine()
	v.Service()
	_, error := v.Get("C")
	if error.Error() != "You didn't insert enough money for C! Inserted: 0, Required: 150" {
		t.Errorf("Improper error message!")
	}	
}

func TestBuyAFromServicedMachine(t *testing.T) {
	v := NewVendingMachine()
	v.Service()
	v.Insert(QUARTER)
	v.Insert(QUARTER)
	v.Insert(DIME)
	v.Insert(NICKLE)
	result, _ := v.Get("A")
	if result != "A" {
		t.Errorf("Machine should have returned an A!")
	}
	if v.items["A"].count != 49 {
		t.Errorf("Machine did not properly decrement inventory!")
	}
}

func TestBuyBFromServicedMachine(t *testing.T) {
	v := NewVendingMachine()
	v.Service()
	v.Insert(DOLLAR)
	result, _ := v.Get("B")
	if result != "B" {
		t.Errorf("Machine should have returned a B!")
	}
	if v.items["B"].count != 49 {
		t.Errorf("Machine did not properly decrement inventory!")
	}
}

func TestBuyCFromServicedMachine(t *testing.T) {
	v := NewVendingMachine()
	v.Service()
	v.Insert(DOLLAR)
	v.Insert(QUARTER)
	v.Insert(QUARTER)
	result, _ := v.Get("C")
	if result != "C" {
		t.Errorf("Machine should have returned a C!")
	}
	if v.items["C"].count != 49 {
		t.Errorf("Machine did not properly decrement inventory!")
	}
}

func TestEmptyCoinReturn(t *testing.T) {
	v := NewVendingMachine()
	result := v.CoinReturn()
	if result != "" {
		t.Errorf("CoinReturn should be empty string with no money inserted!")
	}
}

func TestInsertNickel(t *testing.T) {
	v := NewVendingMachine()
	v.Insert(NICKLE)
	if v.AmountInserted() != 5 {
		t.Errorf("Inserting a single nickle should result in 5 cents total inserted!")
	}
}

func TestCoinReturnWithSingleNickle(t *testing.T) {
	v := NewVendingMachine()
	v.Insert(NICKLE)
	result := v.CoinReturn()
	if result != "N" {
		t.Errorf("CoinReturn should be 'N' after inserting single nickle!")
	}
}

func TestTwoNickles(t *testing.T) {
	v := NewVendingMachine()
	v.Insert(NICKLE)
	v.Insert(NICKLE)

	if v.AmountInserted() != 10 {
		t.Errorf("Inserting two nickles should result in 10 cents total inserted!")
	}

	result := v.CoinReturn()

	if result != "N, N" {
		t.Errorf("CoinReturn should be 'N, N' after inserting two nickles. Was: '%v'", result)
	}
}

func TestSomeCoins(t *testing.T) {
	v := NewVendingMachine()
	v.Insert(NICKLE)
	v.Insert(QUARTER)
	v.Insert(DOLLAR)
	v.Insert(DIME)
	v.Insert(NICKLE)
	v.Insert(DIME)

	if v.AmountInserted() != 155 {
		t.Errorf("Amoint should be 155 cents!")
	}

	result := v.CoinReturn()

	if result != "N, Q, DD, D, N, D" {
		t.Errorf("CoinReturn should be 'N, Q, DD, D, N, D', but was: '%v'", result)
	}
}

func TestMakingChangeReducesBank(t *testing.T) {
	v := NewVendingMachine()
	v.Service()

	if (v.bank[v.coins[NICKLE]] != 50) {
		t.Errorf("Newly serviced machine should have 50 nickles!");
	}

	if (v.bank[v.coins[DIME]] != 50) {
		t.Errorf("Newly serviced machine should have 50 dimes!");
	}

	if (v.bank[v.coins[QUARTER]] != 50) {
		t.Errorf("Newly serviced machine should have 50 quarters!");
	}

	if (v.bank[v.coins[DOLLAR]] != 0) {
		t.Errorf("Newly serviced machine should have 0 dollars!");
	}

	v.Insert(DOLLAR)

	result, _ := v.Get("A")

	if result != "A, Q, D" {
		t.Errorf("Result should be 'A, Q, D' but was %v", result)
	}

	if (v.bank[v.coins[NICKLE]] != 50) {
		t.Errorf("After this purchase machine should have 50 nickles!");
	}

	if (v.bank[v.coins[DIME]] != 49) {
		t.Errorf("After this purchase machine should have 49 dimes!");
	}

	if (v.bank[v.coins[QUARTER]] != 49) {
		t.Errorf("After this purchase machine should have 49 quarters!");
	}

	if (v.bank[v.coins[DOLLAR]] != 1) {
		t.Errorf("After this purchase machine should have 1 dollar!");
	}
}

// Vending Machine Spec Tests

func TestBuyBWithExactChange(t *testing.T) {
	v := NewVendingMachine()
	v.Service()
	v.Insert(QUARTER)
	v.Insert(QUARTER)
	v.Insert(QUARTER)
	v.Insert(QUARTER)

	result, _ := v.Get("B")

	if result != "B" {
		t.Errorf("Result should be 'B' but was %v", result)	
	}	
}

func TestCoinReturn(t *testing.T) {
	v := NewVendingMachine()
	v.Insert(QUARTER)
	v.Insert(QUARTER)

	result := v.CoinReturn()

	if result != "Q, Q" {
		t.Errorf("Result should be 'Q' but was %v", result)	
	}	
}

func TestBuyAWithTooMuchMoney(t *testing.T) {
	v := NewVendingMachine()
	v.Service()
	v.Insert(DOLLAR)

	result, _ := v.Get("A")

	if result != "A, Q, D" {
		t.Errorf("Result should be 'A, Q, D' but was %v", result)
	}
}