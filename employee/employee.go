package employee 

import "fmt"     

type Employee struct {
    Name   string
    Age    int
    Salary float64
}

func NewEmployee(name string, age int, salary float64) *Employee {
    return &Employee{name, age, salary}
}


func (e *Employee) Work() {
    fmt.Printf("%s работает\n", e.Name)
}

func (e *Employee) Rest() {
    fmt.Printf("%s отдыхает\n", e.Name)
}

func (e *Employee) GetPaid() {
    fmt.Printf("%s получил %.2f зарплаты\n", e.Name, e.Salary)
}

func (e *Employee) Meet() {
    fmt.Printf("%s на совещании\n", e.Name)
}

func (e *Employee) Report() {
    fmt.Printf("%s сдаёт отчёт\n", e.Name)
}

// Геттеры и сеттеры (были в твоём коде)
func (e *Employee) GetName() string {
    return e.Name
}

func (e *Employee) GetAge() int {
    return e.Age
}

func (e *Employee) GetSalary() float64 {
    return e.Salary
}

func (e *Employee) SetName(name string) {
    if name != "" {
        e.Name = name
    }
}

func (e *Employee) SetAge(age int) {
    if age > 0 && age < 100 {
        e.Age = age
    }
}

func (e *Employee) SetSalary(salary float64) {
    if salary > 0 {
        e.Salary = salary
    }
}
