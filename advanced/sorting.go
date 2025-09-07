package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	People []Person
	by     func(p1, p2 *Person) bool
}

func (ps *personSorter) Len() int { return len(ps.People) }

func (ps *personSorter) Less(i, j int) bool {
	return ps.by(&ps.People[i], &ps.People[j])
}

func (ps *personSorter) Swap(i, j int) {
	ps.People[i], ps.People[j] = ps.People[j], ps.People[i]
}

func (by By) Sort(people []Person) {
	ps := &personSorter{
		People: people,
		by:     by,
	}
	sort.Sort(ps)
}

// ===============================================
// type ByAge []Person
// type ByName []Person

// func (a ByAge) Len() int { return len(a) }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

// func (a ByName) Len() int { return len(a) }

// func (a ByName) Less(i, j int) bool {
// 	return a[i].Name < a[j].Name
// }

// func (a ByName) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }
// ===============================================

func main_for_sorting() {
	people := []Person{
		{"Steve", 27},
		{"Charlie", 35},
		{"Alice", 30},
		{"Bob", 25},
	}
	fmt.Println("Unsorted by age:", people)

	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	By(age).Sort(people)
	fmt.Println("Sorted by age:", people)

	By(func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}).Sort(people)
	fmt.Println("Sorted by name:", people)

	// ======= SORT.SLICE
	stringSlice := []string{"john", "alice", "bob", "steve", "victor"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorted by last letter:", stringSlice)

	// ===============================================
	// sort.Sort(ByAge(people))
	// fmt.Println("Sorted by age:", people)

	// sort.Sort(ByName(people))
	// fmt.Println("Sorted by age:", people)

	// numbers := []int{5, 3, 8, 6, 2}
	// sort.Ints(numbers)
	// fmt.Println("Sorted numbers:", numbers)

	// stringSlice := []string{"john", "alice", "bob", "steve", "victor"}
	// sort.Strings(stringSlice)
	// fmt.Println("Sorted strings:", stringSlice)
}
