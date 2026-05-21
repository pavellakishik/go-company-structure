package manager

import (
    "fmt"
    "myproject/employee"
)

type Boss interface {
    Lead()
    Plan()
    Hire()
    Delegate()
    Evaluate()
    Strategy()
    Negotiation()
    Approve()
    Represent()
    Revenue()
}

type Manager struct {
    employee.Employee     
    TeamSize   int
    Department string
}

func NewManager(name string, age int, salary float64, teamsize int, department string) *Manager {
    return &Manager{
        Employee:   *employee.NewEmployee(name, age, salary),
        TeamSize:   teamsize,
        Department: department,
    }
}

func (m *Manager) GetTeamSize() int {
    return m.TeamSize
}

func (m *Manager) GetDepartment() string {
    return m.Department
}

func (m *Manager) SetTeamSize(size int) {
    if size > 0 {
        m.TeamSize = size
    }
}

func (m *Manager) SetDepartment(dept string) {
    if dept != "" {
        m.Department = dept
    }
}

// Все методы интерфейса Boss
func (m *Manager) Lead() {
    fmt.Printf("%s ведёт команду из %d человек в отделе %s\n", m.Name, m.TeamSize, m.Department)
}

func (m *Manager) Plan() {
    fmt.Printf("%s планирует работу\n", m.Name)
}

func (m *Manager) Hire() {
    fmt.Printf("%s нанимает новых людей в отдел %s\n", m.Name, m.Department)
}

func (m *Manager) Delegate() {
    fmt.Printf("%s делегирует задачи\n", m.Name)
}

func (m *Manager) Evaluate() {
    fmt.Printf("%s оценивает работу отдела %s\n", m.Name, m.Department)
}

func (m *Manager) Strategy() {
    fmt.Printf("%s участвует в стратегическом планировании отдела %s\n", m.Name, m.Department)
}

func (m *Manager) Negotiation() {
    fmt.Printf("%s ведёт переговоры о ресурсах для отдела %s\n", m.Name, m.Department)
}

func (m *Manager) Approve() {
    fmt.Printf("%s утверждает планы отдела %s\n", m.Name, m.Department)
}

func (m *Manager) Represent() {
    fmt.Printf("%s представляет отдел %s на совещаниях\n", m.Name, m.Department)
}

func (m *Manager) Revenue() {
    fmt.Printf("%s анализирует доходы отдела %s\n", m.Name, m.Department)
}
