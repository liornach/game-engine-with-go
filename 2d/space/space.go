package space

import "errors"

type ObjectInSpace struct {
	Name string
}

type objectsMap map[Coordinate]*ObjectInSpace

type space struct {
	min     Coordinate
	max     Coordinate
	objects objectsMap
}

func NewSpace(min Coordinate, max Coordinate) *space {
	return &space{
		min:     min,
		max:     max,
		objects: make(objectsMap),
	}
}

func (s *space) IsCoodinateInBoundries(coord Coordinate) {
	
}

func (s *space) IsFree(coord Coordinate) bool, error {
	_, found := s.objects[coord]
	return !found
}

func (s *space) OwnedBy(coordinate Coordinate) (*ObjectInSpace, bool) {
	owner, found := s.objects[coordinate]
	return owner, found
}

func (s *space) GetMaxCoordinate() Coordinate {
	return s.max
}

func (s *space) GetMinCoordinate() Coordinate {
	return s.min
}

func (s *space) TryAdd(coord Coordinate, obj *ObjectInSpace) (*ObjectInSpace, error) {
	var owner, found = s.OwnedBy(coord)
	if found {
		var err = errors.New("coordinate already inhabited")
		return owner, err
	}

	s.objects[coord] = obj
	return obj, nil
}

func (s *space) ForceAdd(coord Coordinate, obj *ObjectInSpace) {
	s.objects[coord] = obj
}
