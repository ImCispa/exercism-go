package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var m = map[string]int{}
	m["quarter_of_a_dozen"] = 3
	m["half_of_a_dozen"] = 6
	m["dozen"] = 12
	m["small_gross"] = 120
	m["gross"] = 144
	m["great_gross"] = 1728
	return m
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	m := Units()
	value, isPresent := m[unit]
	if !isPresent {
		return false
	}
	_, isPresentItem := bill[item]
	if isPresentItem {
		bill[item] += value
	} else {
		bill[item] = value
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	m := Units()
	valueItem, isPresentItem := bill[item]
	if !isPresentItem {
		return false
	}
	value, isPresent := m[unit]
	if !isPresent {
		return false
	}
	if valueItem < value {
		return false
	} else if valueItem == value {
		delete(bill, item)
	} else {
		bill[item] -= value
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	valueItem, isPresentItem := bill[item]
	if !isPresentItem {
		return 0, false
	}
	return valueItem, true
}
