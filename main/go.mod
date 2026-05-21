module myproject/main

go 1.22.2

replace myproject/employee => ../employee

replace myproject/manager => ../manager

replace myproject/director => ../director

require (
	myproject/director v0.0.0-00010101000000-000000000000
	myproject/employee v0.0.0-00010101000000-000000000000
	myproject/manager v0.0.0-00010101000000-000000000000
)
