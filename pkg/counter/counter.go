package counter

type Counter[A comparable] struct {
	Counts map[A]int
}

func NewCounter[A comparable]() *Counter[A] {
	return &Counter[A]{Counts: map[A]int{}}
}

func (c *Counter[A]) Add(a A) int {
	if _, ok := c.Counts[a]; !ok {
		c.Counts[a] = 0
	}
	c.Counts[a]++
	return c.Counts[a]
}

func (c *Counter[A]) Remove(a A) bool {
	if _, ok := c.Counts[a]; !ok {
		return false
	}
	c.Counts[a]--
	if c.Counts[a] == 0 {
		delete(c.Counts, a)
	}
	return true
}

func (c *Counter[A]) Get(a A) int {
	return c.Counts[a]
}
