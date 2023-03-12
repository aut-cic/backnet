package model

import (
	"fmt"
	"time"
)

type Discount interface{}

type TimeDiscount struct {
	Since time.Time
	Until time.Time
}

func (t TimeDiscount) Type() string {
	return "time"
}

func (t TimeDiscount) Value() string {
	return fmt.Sprintf(
		"%s-%s",
		t.Since.Format("15:04"),
		t.Until.Format("15:04"),
	)
}

type DayDiscount struct {
	Day time.Weekday
}

// nolint: gomnd
// Value converts weekday into a database compatible string.
// https://pkg.go.dev/time#Weekday
// https://www.geeksforgeeks.org/python-datetime-weekday-method-with-example/
func (d DayDiscount) Value() string {
	return fmt.Sprintf("%d", ((int(d.Day) + 6) % 7))
}

func (d DayDiscount) Type() string {
	return "day_of_week"
}

type DateDiscount struct {
	Date time.Time
}

func (d DateDiscount) Type() string {
	return "date"
}

func (d DateDiscount) Value() string {
	return d.Date.Format(time.DateOnly)
}

type DiscountFactor struct {
	ID     uint `gorm:"<-:false,autoIncrement"`
	Type   string
	Value  string
	Factor float64
}

func (DiscountFactor) TableName() string {
	return "discountfactor"
}
