package main

import (
	"fmt"
)

// Структура посади
type Position struct {
	Title     string
	MinSalary float64
	MaxSalary float64
}

// Структура працівника
type Employee struct {
	ID       int
	Name     string
	Position Position
	Salary   float64
}

// Метод перевірки відповідності зарплати діапазону посади
func (e Employee) IsSalaryInRange() bool {
	return e.Salary >= e.Position.MinSalary && e.Salary <= e.Position.MaxSalary
}

// Функція пошуку працівника з найближчою до максимальної зарплатою
func findClosestToMax(employees map[int]Employee, position Position) *Employee {
	var closest *Employee
	minDiff := position.MaxSalary

	for _, e := range employees {
		if e.Position.Title == position.Title {
			diff := position.MaxSalary - e.Salary
			if diff >= 0 && diff < minDiff {
				minDiff = diff
				closest = &e
			}
		}
	}

	return closest
}

func main() {
	// Створення позицій
	dev := Position{"Developer", 2000, 5000}
	qa := Position{"QA", 1800, 4000}

	// Працівники
	employees := map[int]Employee{
		1: {1, "Alice", dev, 4800},
		2: {2, "Bob", dev, 5050},
		3: {3, "Charlie", qa, 3900},
		4: {4, "Diana", qa, 3700},
	}

	// Перевірка відповідності зарплати
	for _, e := range employees {
		fmt.Printf("%s: зарплата в межах? %v\n", e.Name, e.IsSalaryInRange())
	}

	fmt.Println()

	// Пошук працівника з найвищою зарплатою в межах
	bestDev := findClosestToMax(employees, dev)
	if bestDev != nil {
		fmt.Printf("Розробник з найвищою зарплатою в межах: %s (%.2f)\n", bestDev.Name, bestDev.Salary)
	}
}
