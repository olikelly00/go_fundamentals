package main

import (
	"fmt"
)

// type drink struct {
// 	name   string
// 	volume float32
// 	price  int
// }

// func (myDrink drink) calculatePriceAtCheckout(quantity int) int {
// 	return myDrink.price * quantity
// }

// func main() {
// 	juice := drink{name: "orange juice", volume: 0.33, price: 3}
// 	fmt.Printf("I bough two %v, and cost me %d", juice.name, juice.calculatePriceAtCheckout(17))
// }

type person struct {
	name     string
	age      int
	location string
}

func (myPerson person) announce() string {
	return fmt.Sprintf("I am %s! I am %d years old and I live in %s./n", myPerson.name, myPerson.age, myPerson.location)
}


func (myPerson person) introduce(myFriend person) string {
	return fmt.Sprintf("I, %s, would like to introduce %v of %s", myPerson.name, myFriend.name, myPerson.location)
}


// func main() {
// 	oli := person{name: "oli", age: 26, location: "london"}
// 	myFriend := person{name: "rachel", age: 31, location: "paris"}
// 	fmt.Printf("I am %s! I am %d years old and I live in %s./n", oli.name, oli.age, oli.location)
// 	fmt.Println(person.introduce(oli, myFriend))
// }



func main() {
	var friendArray [5]person
	friendArray[0] = person{name: "Oli", age: 28, location: "London"}
	friendArray[1] = person{name: "Josh", age: 29, location: "Paris"}
	friendArray[2] = person{name: "Fara", age: 24, location: "Milan"}
	friendArray[3] = person{name: "Emily", age: 23, location: "Barcelona"}
	friendArray[4] = person{name: "John", age: 21, location: "Tokyo"}
	for friend := range len(friendArray)-1 {
		fmt.Printf("Hi, my name is %v. I am %v years old and I'm living in %v. I'd like to introduce my friend %v.\n", friendArray[friend].name, friendArray[friend].age, friendArray[friend].location, friendArray[friend+1].name)
	}
	fmt.Printf("Hi, my name is %v. I am %v years old and I'm living in %v.\n", friendArray[len(friendArray)-1].name, friendArray[len(friendArray)-1].age, friendArray[len(friendArray)-1].location)
}



