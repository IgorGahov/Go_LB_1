package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSalaryInRange(t *testing.T) {
	tests := []struct {
		name     string
		employee Employee
		expected bool
	}{
		{"В межах", Employee{Salary: 3000, Position: Position{"Developer", 2000, 5000}}, true},
		{"Нижче мінімуму", Employee{Salary: 1500, Position: Position{"Developer", 2000, 5000}}, false},
		{"Вище максимуму", Employee{Salary: 6000, Position: Position{"Developer", 2000, 5000}}, false},
		{"Рівно мінімум", Employee{Salary: 2000, Position: Position{"Developer", 2000, 5000}}, true},
		{"Рівно максимум", Employee{Salary: 5000, Position: Position{"Developer", 2000, 5000}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.employee.IsSalaryInRange()
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestFindClosestToMax(t *testing.T) {
	dev := Position{"Developer", 2000, 5000}
	employees := map[int]Employee{
		1: {1, "Alice", dev, 4800},
		2: {2, "Bob", dev, 4999},
		3: {3, "Charlie", dev, 5100}, // поза межами
		4: {4, "Diana", Position{"QA", 1800, 4000}, 3900},
	}

	result := findClosestToMax(employees, dev)

	require.NotNil(t, result)
	require.Equal(t, "Bob", result.Name)
	require.Equal(t, 4999.0, result.Salary)
}
