package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) (out []Record) {

	// for each record, apply predicate
	for _, rec := range in {
		if predicate(rec) {
			out = append(out, rec)
		}
	}
	return
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) (out float64) {

	// filter records by days period and add the amount
	for _, rec := range Filter(in, ByDaysPeriod(p)) {
		out += rec.Amount
	}
	return
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (out float64, err error) {

	// get filtered records for given category
	filteredRec := Filter(in, ByCategory(c))

	// if no records found, set error else calculate total amount for given period
	if len(filteredRec) == 0 {
		err = errors.New("unknown category " + c)
	} else {
		out = TotalByPeriod(filteredRec, p)
	}
	return
}
