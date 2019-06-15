package day297

import "sort"

// Customer is the customer's id.
type Customer int

// Drink is the drink's id.
type Drink int

// LazyBartender returns the fewest number of drinks
// to satisfy all customers.
func LazyBartender(prefs map[Customer][]Drink) []Drink {
	var drinks []Drink
	inverted := make(map[Drink][]Customer)
	for cust, favorites := range prefs {
		for _, drink := range favorites {
			inverted[drink] = append(inverted[drink], cust)
		}
	}
	totalCustomers := len(prefs)
	served := make(map[Customer]struct{}, totalCustomers)
	for len(served) < totalCustomers {
		var largest []Customer
		var largestDrink Drink
		var largestIncrease int
		for drink, custs := range inverted {
			var newlyServed int
			for _, cust := range custs {
				if _, found := served[cust]; !found {
					newlyServed++
				}
			}
			if newlyServed > largestIncrease {
				largestIncrease = newlyServed
				largestDrink = drink
				largest = custs
			}
		}
		delete(inverted, largestDrink)
		drinks = append(drinks, largestDrink)
		for _, cust := range largest {
			served[cust] = struct{}{}
		}
	}
	sort.Slice(drinks, func(i, j int) bool {
		return drinks[i] < drinks[j]
	})
	return drinks
}
