package set

func New(items ...any) Set {
	st := set{
		elements: make(map[any]bool),
	}
	for _, item := range items {
		st.Add(item)
	}
	return &st
}

type Set interface {
	Add(iem any)                      // adds new element to the set
	Delete(item any)                  // deletes the passed. element from the set if present
	Len() int                         // gives the length of the set (total no. of elements in set)
	GetItems() []any                  // gives the array([]any) of elements of the set
	In(item any) bool                 // checks whether item is present in set or not
	IsSubsetOf(set2 Set) bool         // checks whether set is subset of set2 or not
	IsSupersetOf(set2 Set) bool       // checks whether set is superset of set2 or nor
	Union(set2 Set) Set               // gives new union set of both sets
	Intersection(set2 Set) Set        // gives new intersection set of both sets
	Difference(set2 Set) Set          // gives new difference set of both sets
	SymmetricDifference(set2 Set) Set // gives new symmetric difference set of both sets
}

type set struct {
	elements map[any]bool
}

func (st *set) Add(value any) {
	st.elements[value] = true
}

func (st *set) Delete(value any) {
	delete(st.elements, value)
}

func (st *set) GetItems() []any {
	keys := make([]any, 0, len(st.elements))
	for k := range st.elements {
		keys = append(keys, k)
	}
	return keys
}

func (st *set) Len() int {
	return len(st.elements)
}

func (st *set) In(value any) bool {
	if _, in := st.elements[value]; in {
		return true
	}
	return false
}

func (st *set) IsSubsetOf(superSet Set) bool {
	if st.Len() > superSet.Len() {
		return false
	}

	for _, item := range st.GetItems() {
		if !superSet.In(item) {
			return false
		}
	}
	return true
}

func (st *set) IsSupersetOf(subSet Set) bool {
	return subSet.IsSubsetOf(st)
}

func (st *set) Union(st2 Set) Set {
	unionSet := New()
	for _, item := range st.GetItems() {
		unionSet.Add(item)
	}
	for _, item := range st2.GetItems() {
		unionSet.Add(item)
	}
	return unionSet
}

func (st *set) Intersection(st2 Set) Set {
	intersectionSet := New()
	var minSet, maxSet Set
	if st.Len() > st2.Len() {
		minSet = st2
		maxSet = st
	} else {
		minSet = st
		maxSet = st2
	}
	for _, item := range minSet.GetItems() {
		if maxSet.In(item) {
			intersectionSet.Add(item)
		}
	}
	return intersectionSet
}

func (st *set) Difference(st2 Set) Set {
	differenceSet := New()
	for _, item := range st.GetItems() {
		if !st2.In(item) {
			differenceSet.Add(item)
		}
	}
	return differenceSet
}

func (st *set) SymmetricDifference(st2 Set) Set {
	symmetricDifferenceSet := New()
	dropSet := New()
	for _, item := range st.GetItems() {
		if st2.In(item) {
			dropSet.Add(item)
		} else {
			symmetricDifferenceSet.Add(item)
		}
	}

	for _, item := range st2.GetItems() {
		if !dropSet.In(item) {
			symmetricDifferenceSet.Add(item)
		}
	}

	return symmetricDifferenceSet
}
