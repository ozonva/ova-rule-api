package models

import "fmt"

type Rule struct {
	Name   string
	Count  uint64
	UserID int64
}

func (r Rule) String() string {
	return fmt.Sprintf("<%s: Count = %d>", r.Name, r.Count)
}

func (r *Rule) IncCount() {
	r.Count++
}
