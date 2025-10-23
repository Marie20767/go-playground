package algorithms

import (
	"time"
)

// employee A works at company A, given T is the total number of holidays that employee A gets per year,
// r is the remaning number of holidays she has for the current year and d is the date she is planning to leave
// (as a string in DD/MM/YYYY format), write a function that takes t, r and d and returns the number of holidays that employee A is due in refund or has in excess.
// E.g. if employee A leaves exactly halfway through the year and has used exactly half her holiday, the function should return 0.

func HolidayExitBalance(t int, r int, d string) float32 {
	const workingDaysPerYear = 260
	const workingDaysPerMonth = 21.6
	const dateTimeLayout = "02/01/2006"

	datetime, err := time.Parse(dateTimeLayout, d)

	if err != nil {
		panic(err)
	}

	dailyHolAllowance := float32(t) / workingDaysPerYear

	_, month, day := datetime.Date()
	monthNum := int(month)

	totalDaysWorkedUpToPrevMonth := float32(monthNum-1) * workingDaysPerMonth
	totalDaysWorkedLeavingMonth := float32(day) / 30 * workingDaysPerMonth
	totalDaysWorked := totalDaysWorkedUpToPrevMonth + totalDaysWorkedLeavingMonth

	holAllowanceUntilExit := dailyHolAllowance * float32(totalDaysWorked)

	return holAllowanceUntilExit - float32(t-r)
}

