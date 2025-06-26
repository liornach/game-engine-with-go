package logic

import "fmt"

type NotInWorldError struct{}

type NotFreeError struct {
	Pos   Pos
	Owner Owner
}

func (e NotInWorldError) Error() string {
	return "not in world"
}

func (e NotFreeError) Error() string {
	ret := fmt.Sprintf("position %s is owned by %v", e.Pos.String(), e.Owner)
	return ret
}
