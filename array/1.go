package main

import "fmt"

func main(){
// var arrName [arraySize]dataType
// var nmubers [5]int
// Method 1: Specify all elements
var numbers [5]int = [5]int{1, 2, 3, 4, 5}

// Method 2: Using shorthand (Go infers the type)
names := [3]string{"Alice", "Bob", "Charlie"}

// Method 3: Using ... to let the compiler determine the size
grades := [...]int{95, 88, 76, 92} // Size will be 4

firstNumber := numbers[0]
fmt.Println(firstNumber)
numbers[4] =60
fmt.Println(numbers)
arrayLength :=len(grades)
fmt.Println(arrayLength)

for i :=0;i<len(names);i++{
  fmt.Println(names[i])
}

}
