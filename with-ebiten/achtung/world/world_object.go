package world

import "image/color"

type Uid = string

type WorldObject struct {
	color color.RGBA
	uid   Uid
}

func NewWorldObject(c color.RGBA, u Uid) WorldObject {
	return WorldObject{
		color: c,
		uid:   u}
}

func (w WorldObject) Color() color.RGBA {
	return w.color
}

func (w WorldObject) Uid() Uid {
	return w.uid
}
