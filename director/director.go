package director

import (
    "fmt"
    "myproject/manager"
)

type Director struct {
    manager.Manager
    CompanyName    string
    CompanyRevenue float64
}

func NewDirector(name string, age int, salary float64, teamsize int, department string, companyName string, companyRevenue float64) *Director {
    return &Director{
        Manager:        *manager.NewManager(name, age, salary, teamsize, department),
        CompanyName:    companyName,
        CompanyRevenue: companyRevenue,
    }
}

func (d *Director) GetCompanyName() string {
    return d.CompanyName
}

func (d *Director) GetCompanyRevenue() float64 {
    return d.CompanyRevenue
}

func (d *Director) SetCompanyName(name string) {
    if name != "" {
        d.CompanyName = name
    }
}

func (d *Director) SetCompanyRevenue(revenue float64) {
    if revenue > 0 {
        d.CompanyRevenue = revenue
    }
}

func (d *Director) Strategy() {
    fmt.Printf("%s разрабатывает стратегию для компании %s\n", d.Name, d.CompanyName)
}

func (d *Director) Negotiation() {
    fmt.Printf("%s ведёт переговоры\n", d.Name)
}

func (d *Director) Approve() {
    fmt.Printf("%s утверждает бюджет для отдела %s\n", d.Name, d.Department)
}

func (d *Director) Represent() {
    fmt.Printf("%s представляет компанию %s\n", d.Name, d.CompanyName)
}

func (d *Director) Revenue() {
    fmt.Printf("%s подсчитывает доходы компании: %.2f\n", d.Name, d.CompanyRevenue)
}
