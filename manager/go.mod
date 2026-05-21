module myproject/manager

go 1.22.2

replace myproject/employee => ../employee

require myproject/employee v0.0.0-00010101000000-000000000000
