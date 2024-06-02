package hashset

type Set[T comparable] struct {
	set map[T]bool
}

// return false if already added
func (this *Set[T]) Add(val T) bool {
	ok := this.Exists(val)
	if ok {
		return false
	}
	this.set[val] = true
	return true
}

func (this *Set[T]) Exists(val T) bool {
	_, ok := this.set[val]
	return ok
}

func (this *Set[T]) Delete(val T) {
	delete(this.set, val)
}

func (this *Set[T]) Size(val T) int {
	return len(this.set)
}
