package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	WeekDays = 7

	DefaultType = "default"
	DateType    = "date"
	DayType     = "day_of_week"
	TimeType    = "time"
)

var ErrInvalidDiscountType = errors.New("unknown discount type")

type Discount interface {
	Value() string
	Type() string
	Factor() float64
}

func ToDiscount(df DiscountFactor) (Discount, error) {
	switch df.Type {
	case DateType:
		dd, err := NewDateDiscount(df)
		if err != nil {
			return nil, fmt.Errorf("invalid date discount record %w", err)
		}

		return dd, nil
	case DefaultType:
		return NewDefaultDiscount(df), nil
	case TimeType:
		td, err := NewTimeDiscount(df)
		if err != nil {
			return nil, fmt.Errorf("invalid time discount record %w", err)
		}

		return td, nil
	case DayType:
		dd, err := NewDayDiscount(df)
		if err != nil {
			return nil, fmt.Errorf("invalid day_of_week discount record %w", err)
		}

		return dd, nil
	default:
		return nil, ErrInvalidDiscountType
	}
}

type DefaultDiscount struct {
	F float64
}

func NewDefaultDiscount(df DiscountFactor) DefaultDiscount {
	return DefaultDiscount{
		F: df.Factor,
	}
}

func (DefaultDiscount) Value() string {
	return DefaultType
}

func (DefaultDiscount) Type() string {
	return DefaultType
}

func (d DefaultDiscount) Factor() float64 {
	return d.F
}

type TimeDiscount struct {
	Since time.Time
	Until time.Time
	F     float64
}

func NewTimeDiscount(df DiscountFactor) (TimeDiscount, error) {
	values := strings.Split(df.Value, "-")

	since, err := time.Parse("15:04", values[0])
	if err != nil {
		return TimeDiscount{}, fmt.Errorf("cannot parse since field: %w", err)
	}

	until, err := time.Parse("15:04", values[1])
	if err != nil {
		return TimeDiscount{}, fmt.Errorf("cannot parse until field: %w", err)
	}

	return TimeDiscount{
		F:     df.Factor,
		Until: until,
		Since: since,
	}, nil
}

func (t TimeDiscount) Type() string {
	return TimeType
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

func NewDayDiscount(df DiscountFactor) (DayDiscount, error) {
	dw, err := strconv.Atoi(df.Value)
	if err != nil {
		return DayDiscount{}, fmt.Errorf("cannot parse day field: %w", err)
	}

	return DayDiscount{
		F:   df.Factor,
		Day: time.Weekday((dw + 1) % WeekDays),
	}, nil
}

// Value converts weekday into a database compatible string.
// https://pkg.go.dev/time#Weekday
// https://www.geeksforgeeks.org/python-datetime-weekday-method-with-example/
func (d DayDiscount) Value() string {
	return fmt.Sprintf("%d", ((int(d.Day) - 1) % WeekDays))
}

func (d DayDiscount) Type() string {
	return DayType
}

func (d DayDiscount) Factor() float64 {
	return d.F
}

type DateDiscount struct {
	Date time.Time
	F    float64
}

func NewDateDiscount(df DiscountFactor) (DateDiscount, error) {
	date, err := time.Parse(time.DateOnly, df.Value)
	if err != nil {
		return DateDiscount{}, fmt.Errorf("cannot parse date field: %w", err)
	}

	return DateDiscount{
		F:    df.Factor,
		Date: date,
	}, nil
}

func (d DateDiscount) Type() string {
	return DateType
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
