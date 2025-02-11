package repository

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

type Repository struct {
	Client *spanner.Client
}
// Employee struct defined within the function
type Employee struct {
	ID        int64
	FirstName string
	LastName  string
	JoiningDate time.Time
	Salary    float64
}

// RegisterEmployee inserts a new employee record and returns the whole employee object
func (r *Repository) RegisterEmployee(ctx context.Context, firstName string, lastName string, joiningDate string, salary float64) (*Employee, error) {
	employee := &Employee{}
	joinDate, errdate := time.Parse("2006-01-02", joiningDate)
	if errdate != nil {
		return nil, fmt.Errorf("invalid date format (expected YYYY-MM-DD): %w", errdate)
	}
	_, err := r.Client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{
			SQL: `INSERT INTO employees (first_name, last_name, joining_date, salary)
                  VALUES (@FirstName, @LastName, @JoiningDate, @Salary)
                  THEN RETURN id, first_name, last_name, joining_date, salary`,
			Params: map[string]interface{}{
				"FirstName": firstName,
				"LastName":  lastName,
				"JoiningDate": joinDate.Format("2006-01-02"),
				"Salary":    big.NewRat(int64(salary*100), 100),
			},
		}

		iter := txn.Query(ctx, stmt)
		defer iter.Stop()

		for {
			row, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return fmt.Errorf("failed to execute query: %v", err)
			}
			var salaryNumeric big.Rat
			var date spanner.NullDate
			if err := row.Columns(&employee.ID, &employee.FirstName, &employee.LastName, &date, &salaryNumeric); err != nil {
				return fmt.Errorf("failed to parse employee data: %v", err)
			}
			employee.JoiningDate = time.Date(date.Date.Year, date.Date.Month, date.Date.Day, 0, 0, 0, 0, time.UTC)
			f64, _ := salaryNumeric.Float64()
			employee.Salary = f64
		}

		return nil
	})
	fmt.Println("Employee registered:", employee, err)

	if err != nil {
		return nil, fmt.Errorf("failed to insert employee: %v", err)
	}

	return employee, nil
}

