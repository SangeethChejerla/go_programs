package main

import "fmt"

func main(){
  x:=10
name,age := "Alice",30
  
  if true{
    x:=20
    fmt.Println(x)
    fmt.Println(name,age)
  }
  fmt.Println(x)
}
