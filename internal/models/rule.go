package models

import "fmt"

type Rule struct {
	ID     uint64
	Name   string
	Count  uint64
	UserID uint64
}

func (r Rule) String() string {
	return fmt.Sprintf("<%s: Count = %d>", r.Name, r.Count)
}

func (r *Rule) IncCount() {
	r.Count++
}
