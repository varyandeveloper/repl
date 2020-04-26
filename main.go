package main

import (
    "fmt"
)

func BuilEmployee(name string) EmployeeInterface {
  return &AbstractEmployee{
    name: name,
  }
}

func main() {
  employee := BuilEmployee("Varazdat Stepanyan")
  
  fmt.Println(employee.GetName())
}

type EmployeeInterface interface {
  GetName() string
  SetName(name string) EmployeeInterface
}

type ManagerInterface interface {
  EmployeeInterface
}

type AbstractEmployee struct {
  EmployeeInterface
  name string
}

func (employee *AbstractEmployee) GetName() string {
  return employee.name
}

func (employee *AbstractEmployee) SetName(name string) EmployeeInterface {
  employee.name = name
  return employee
}