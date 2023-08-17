package salary

import "testing"

func TestCalculateSalary(t *testing.T) {

	t.Run("Calculate Developer salary", func(*testing.T) {

		salary := Salary(10, 10, DeveloperSalary)

		got := salary
		want := 100

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

	t.Run("Calculate Manager salary", func(*testing.T) {

		salary := Salary(10, 10, ManagerSalary)

		got := salary
		want := 20

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

}
