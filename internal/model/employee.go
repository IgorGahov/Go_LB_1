package model

type Position struct {
	ID        int     `db:"id"`
	Title     string  `db:"title"`
	MinSalary float64 `db:"min_salary"`
	MaxSalary float64 `db:"max_salary"`
}

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	PositionID int     `db:"position_id"`
	Salary     float64 `db:"salary"`
}
