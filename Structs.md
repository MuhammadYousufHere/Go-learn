# Structs

Aggregate data type
--> groups together other objects arbitray types into one
type Person struct {
name string
age number
address string
}

var p1 Person
or

p2 := new(Person)
or
p3 := Person(name:"Ali",age:19,address:"karachi")

can access the properties using dot notation
add = p3.name
