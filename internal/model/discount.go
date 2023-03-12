package model

import (
	"fmt"
	"time"
)

type Discount interface {
	Value() string
	Type() string
	Factor() float64
}

type DefaultDiscount struct {
	F float64
}

func (DefaultDiscount) Value() string {
	return "default"
}

func (DefaultDiscount) Type() string {
	return "default"
}

func (d DefaultDiscount) Factor() float64 {
	return d.F
}

type TimeDiscount struct {
	Since time.Time
	Until time.Time
	F     float64
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

func (t TimeDiscount) Factor() float64 {
	return t.F
}

type DayDiscount struct {
	Day time.Weekday
	F   float64
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

func (d DayDiscount) Factor() float64 {
	return d.F
}

type DateDiscount struct {
	Date time.Time
	F    float64
}

func (d DateDiscount) Type() string {
	return "date"
}

func (d DateDiscount) Value() string {
	return d.Date.Format(time.DateOnly)
}

func (d DateDiscount) Factor() float64 {
	return d.F
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
