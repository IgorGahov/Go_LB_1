package service

import "webapp/internal/model"

func IsSalaryInRange(e *model.Employee, p *model.Position) bool {
	return e.Salary >= p.MinSalary && e.Salary <= p.MaxSalary
}
