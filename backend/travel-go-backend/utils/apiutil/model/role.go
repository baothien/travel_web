package model

type Role int

const (
	CUSTOMER Role = iota
	DRIVER
	EMPLOYEE
)

func (r Role) Name() string {
	return [...]string{"CUSTOMER", "DRIVER", "EMPLOYEE"}[r]
}
