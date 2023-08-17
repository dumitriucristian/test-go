package salary

func Salary(firstValue int, secondValue int, f func(int, int) int) int {
	total := f(firstValue, secondValue)
	return total
}

func DeveloperSalary(hours int, price int) int {
	return hours * price
}

func ManagerSalary(bonus int, regularSalary int) int {
	return bonus + regularSalary
}
