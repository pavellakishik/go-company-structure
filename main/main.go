package main

import (
    "github.com/pavellakishik/go-company-structure/employee"
    "github.com/pavellakishik/go-company-structure/manager"
    "github.com/pavellakishik/go-company-structure/director"
)

func showBossActions(boss manager.Boss, name string) {
    fmt.Printf("\n=== Действия руководителя %s ===\n", name)
    boss.Lead()
    boss.Plan()
    boss.Hire()
    boss.Delegate()
    boss.Evaluate()
    boss.Strategy()
    boss.Negotiation()
    boss.Approve()
    boss.Represent()
    boss.Revenue()
}

func main() {
    fmt.Println("СОЗДАНИЕ ОБЪЕКТОВ")
    
    emp := employee.NewEmployee("Дмитрий", 35, 75000)
    mgr := manager.NewManager("Анна", 30, 90000, 8, "Разработка")
    dir := director.NewDirector("Олег", 50, 150000, 50, "Управление", "ТехноПром", 10000000)
    
    fmt.Println("\nХАРАКТЕРИСТИКИ")
    fmt.Printf("Сотрудник: %s, %d лет, зарплата %.0f\n", emp.GetName(), emp.GetAge(), emp.GetSalary())
    fmt.Printf("Менеджер: %s, %d лет, команда %d, отдел %s\n", mgr.GetName(), mgr.GetAge(), mgr.GetTeamSize(), mgr.GetDepartment())
    fmt.Printf("Директор: %s, %d лет, компания %s, доход %.0f, отдел %s\n", 
        dir.GetName(), dir.GetAge(), dir.GetCompanyName(), dir.GetCompanyRevenue(), dir.GetDepartment())
    
    fmt.Println("\nМЕТОДЫ СОТРУДНИКА")
    emp.Work()
    emp.Rest()
    emp.GetPaid()
    emp.Meet()
    emp.Report()
    
    fmt.Println("\nМЕТОДЫ МЕНЕДЖЕРА")
    mgr.Work()
    mgr.Lead()
    mgr.Plan()
    mgr.Hire()
    mgr.Delegate()
    mgr.Evaluate()
    
    fmt.Println("\nМЕТОДЫ ДИРЕКТОРА")
    dir.Work()
    dir.Strategy()
    dir.Negotiation()
    dir.Approve()
    dir.Represent()
    dir.Revenue()
    
    fmt.Println("\nПОЛИМОРФИЗМ")
    bosses := []manager.Boss{mgr, dir}
    for _, boss := range bosses {
        switch b := boss.(type) {
        case *manager.Manager:
            showBossActions(b, b.GetName())
        case *director.Director:
            showBossActions(b, b.GetName())
        }
    }
    
    fmt.Println("\nИЗМЕНЕНИЕ ХАРАКТЕРИСТИК")
    emp.SetName("Иван Петров")
    emp.SetAge(31)
    emp.SetSalary(55000)
    fmt.Printf("Обновлённый сотрудник: %s, %d лет, зарплата %.0f\n", emp.GetName(), emp.GetAge(), emp.GetSalary())
    
    mgr.SetDepartment("Маркетинг")
    mgr.SetTeamSize(10)
    fmt.Printf("Обновлённый менеджер: отдел %s, команда %d\n", mgr.GetDepartment(), mgr.GetTeamSize())
    
    dir.SetCompanyRevenue(12000000)
    dir.SetCompanyName("ТехноИнновации")
    dir.SetDepartment("Стратегическое развитие")
    fmt.Printf("Обновлённый директор: компания %s, доход %.0f, отдел %s\n", 
        dir.GetCompanyName(), dir.GetCompanyRevenue(), dir.GetDepartment())
    
    fmt.Println("\nЗАЩИТА ДАННЫХ")
    fmt.Println("Попытка установить некорректные значения:")
    emp.SetAge(-5)
    emp.SetSalary(-100)
    mgr.SetTeamSize(-3)
    
    fmt.Printf("Корректные данные сохранены: возраст %d, зарплата %.0f, команда %d\n", 
        emp.GetAge(), emp.GetSalary(), mgr.GetTeamSize())
}
