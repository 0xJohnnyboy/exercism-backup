package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    return map[string]int{
        "quarter_of_a_dozen":	3,
        "half_of_a_dozen":	6,
        "dozen":	12,
        "small_gross":	120,
        "gross":	144,
        "great_gross":	1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill map[string]int, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
        return false
    }

    bill[item] += units[unit]
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
        return false
    }

    found := false
    for k, _ := range bill {
        if item == k {
            found = true
            break
        }
    }
    if found == false {
        return false
    }

    billQty := bill[item]
    targetQty := billQty - units[unit]

    if targetQty < 0 {
        return false
    }
    if targetQty == 0 {
        delete(bill, item)
        return true
    }
    bill[item] = targetQty
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    if _, ok := bill[item]; !ok  {
        return 0, false
    }
	
    return bill[item], true
}
