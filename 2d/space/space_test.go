package space

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewCoordinate(t *testing.T) {
	x, y, z := 3, 4, 5
	exp := Coordinate{
		X: x,
		Y: y,
		Z: z,
	}
	res := NewCoordinate(x, y, z)
	failIfDismatchCoordinate(t, res, exp)
}

func TestNewSpace(t *testing.T) {
	expMinCoord := NewCoordinate(-3, -4, -5)
	expMaxCoord := NewCoordinate(3, 4, 5)
	expSpace := space{
		min:     expMinCoord,
		max:     expMaxCoord,
		objects: make(map[Coordinate]*ObjectInSpace),
	}

	resSpace := NewSpace(expMinCoord, expMaxCoord)
	failifDismatchSpace(t, resSpace, &expSpace)
}

func TestGetMaxCoordinate(t *testing.T) {
	minCoordinate := NewCoordinate(-3, -4, -5)
	expMaxCoordinate := NewCoordinate(3, 4, 5)
	s := NewSpace(minCoordinate, expMaxCoordinate)
	resMaxCoordinate := s.GetMaxCoordinate()

	failIfDismatchCoordinate(t, resMaxCoordinate, expMaxCoordinate)
}

func TestGetMinCoordinate(t *testing.T) {
	expMinCoordinate := NewCoordinate(-3, -4, -5)
	maxCoordinate := NewCoordinate(3, 4, 5)
	s := NewSpace(expMinCoordinate, maxCoordinate)
	resMinCoordinate := s.GetMinCoordinate()

	failIfDismatchCoordinate(t, resMinCoordinate, expMinCoordinate)
}

func failifDismatchSpace(t *testing.T, res *space, exp *space) {
	failIfDismatchCoordinate(t, res.min, exp.min)
	failIfDismatchCoordinate(t, res.max, exp.max)
	failIfDismatchObjectMap(t, res.objects, exp.objects)
}

func failIfDismatchObjectMap(t *testing.T, res objectsMap, exp objectsMap) {
	failIfDismatchInt(t, "res length", len(res), "exp length", len(exp))

	for coord, o := range res {
		failIfDismatchObjectInSpace(t, o, exp[coord])
	}
}

func failIfDismatchCoordinate(t *testing.T, res Coordinate, exp Coordinate) {
	failIfDismatchInt(t, "res.x", res.X, "exp.x", exp.X)
	failIfDismatchInt(t, "res.y", res.Y, "exp.y", exp.Y)
	failIfDismatchInt(t, "res.z", res.Z, "exp.z", exp.Z)
}

func failIfDismatchObjectInSpace(t *testing.T, res *ObjectInSpace, exp *ObjectInSpace) error {
	return failIfDismatchString(t, "res.Name", res.Name, "exp.Name", exp.Name)
}

func failIfDismatchInt(t *testing.T, resname string, resvalue int, expectname string, expectvalue int) {
	if resvalue != expectvalue {
		t.Errorf("%s : %d, %s : %d", resname, resvalue, expectname, expectvalue)
	}
}

func failIfDismatchString(t *testing.T, resname string, resvalue string, expectname string, expectvalue string) error {
	if resvalue != expectvalue {
		errStr := fmt.Sprintf("%s : %s, %s : %s", resname, resvalue, expectname, expectvalue)
		t.Error(errStr)
		return errors.New(errStr)
	}

	return nil
}
