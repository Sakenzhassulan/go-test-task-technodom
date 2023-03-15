package store

type Cache interface {
	Add(key, value string)
	Get(key string) (value string, ok bool)
	Len() int
}

type Store struct {
	data map[string]string
	size int
}

func NewCache(size int) *Store {
	return &Store{
		data: make(map[string]string),
		size: size,
	}
}

func (c *Store) Add(key, value string) {

	if len(c.data) >= c.size {
		// Remove the oldest entry
		for k := range c.data {
			delete(c.data, k)
			break
		}
	}
	c.data[key] = value
}

func (c *Store) Get(key string) (value string, ok bool) {
	value, ok = c.data[key]
	return
}

func (c *Store) Len() int {
	return len(c.data)
}
