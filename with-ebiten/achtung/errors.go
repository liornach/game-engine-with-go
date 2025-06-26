package achtung

import "fmt"

type InvalidRadError struct {
	r Rad
}

func (e InvalidRadError) Error() string {
	var err string
	if e.r < MinRad {
		err = fmt.Sprintf("radian must be greater than or equal to %f", MinRad)
	} else if e.r >= MaxRad {
		err = fmt.Sprintf("radian must be less than %f", MaxRad)
	} else {
		panic("no error found, consider debugging")
	}

	return err
}

type PlayerNotFoundError struct {
	p Player
}

func (e PlayerNotFoundError) Error() string {
	err := fmt.Sprintf("player with uuid %v was not found", e.p)
	return err
}
