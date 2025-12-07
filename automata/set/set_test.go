package set

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s1 Set[int]
		s2   Set[int]
		want Set[int]
	}{
		{
			name: "Union of two sets with no commmon elements",
			s1: Set[int]{5:{}, 2:{}},
			s2: Set[int]{3:{}, 4: {}},
			want: Set[int]{2:{}, 3:{}, 4:{}, 5:{}},
		},
		{
			name: "Union of a set with elements and an empty set",
			s1: Set[int]{5:{}, 2:{}},
			s2: Set[int]{},
			want: Set[int]{2:{}, 5:{}},
		},
		{
			name: "Union of two sets with common elements",
			s1: Set[int]{5:{}, 2:{}},
			s2: Set[int]{5:{}, 3:{}},
			want: Set[int]{2:{}, 3:{}, 5:{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Union(tt.s1, tt.s2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s1   Set[int]
		s2   Set[int]
		want Set[int]
	}{
		{
			name: "Intersection of two sets with no commmon elements",
			s1: Set[int]{5:{}, 2:{}},
			s2: Set[int]{3:{}, 4: {}},
			want: Set[int]{},
		},
		{
			name: "Intersection of two sets with common elements",
			s1: Set[int]{5:{}, 2:{}},
			s2: Set[int]{5:{}, 3:{}},
			want: Set[int]{5:{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersection(tt.s1, tt.s2)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
