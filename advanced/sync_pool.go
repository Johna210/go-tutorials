package main

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func main_sync_pool() {
	var pool = sync.Pool{
		// New: func() interface{} {
		// 	fmt.Println("Creating a new Person.")
		// 	return &person{}
		// },
	}
	pool.Put(&person{name: "John", age: 25})

	// Get an object from the pool
	person1 := pool.Get().(*person)
	fmt.Printf("Person1 - Name: %s, Age: %d\n", person1.name, person1.age)
	// person1.name = "John"
	// person1.age = 30
	// fmt.Printf("Got person: %+v\n", person1)
	//

	pool.Put(person1)
	fmt.Println("Returned person1 to the pool.")

	person2 := pool.Get().(*person)
	fmt.Printf("Person2 - Name: %s, Age: %d\n", person2.name, person2.age)
	// fmt.Printf("Got person2: %+v\n", person2)

	person3 := pool.Get()
	if person3 != nil {
		fmt.Printf("Got person3: %+v\n", person3)
		// fmt.Printf("Person3 - Name: %s, Age: %d\n", person3.name, person3.age)
		person3.(*person).name = "Alice"
		pool.Put(person3)
	} else {
		fmt.Println("Sync pool is empty")
	}

	// Returning person2 and person3 to the pool
	pool.Put(person2)
	fmt.Println("Returned person2 and person3 to the pool.")

	person4 := pool.Get().(*person)
	// fmt.Printf("Got person4: %+v\n", person4)
	fmt.Printf("Person4 - Name: %s, Age: %d\n", person4.name, person4.age)

	person5 := pool.Get().(*person)
	fmt.Printf("Got person5: %+v\n", person5)
	// fmt.Printf("Person5 - Name: %s, Age: %d\n", person5.name, person5.age)
}
