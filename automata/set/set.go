package set

type Set[T comparable] map[T]struct{}

func Union[T comparable](s1, s2 Set[T]) Set[T] {
	union := Set[T]{}

	for k := range s1 {
		union[k] = struct{}{}
	}
	for k := range s2 {
		union[k] = struct{}{}
	}

	return union
}

func Intersection[T comparable](s1, s2 Set[T]) Set[T] {
	intersection := Set[T]{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for k := range s1 {
		if _, ok := s2[k]; ok {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

func Add[T comparable](s *Set[T], element T) {
	(*s)[element] = struct{}{}
}

func Remove[T comparable](s *Set[T], element T) {
	delete(*s, element)
}