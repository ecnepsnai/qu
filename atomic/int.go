package atomic

import "sync"

// Integer describes a threadsave integer
type Integer struct {
	mutex *sync.Mutex
	value int
}

// NewInteger create a new threadsafe interger with the initial value
func NewInteger(initialValue int) *Integer {
	return &Integer{
		mutex: &sync.Mutex{},
		value: initialValue,
	}
}

// Get get the current value. May block.
func (i *Integer) Get() int {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	return i.value
}

// IncrementAndGet increment the current value and return it. May block.
func (i *Integer) IncrementAndGet() int {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.value++
	return i.value
}

// DecrementAndGet drecrement the current value and return it. May block.
func (i *Integer) DecrementAndGet() int {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.value--
	return i.value
}

// SetAndGet set the current value and return it. May block.
func (i *Integer) SetAndGet(newValue int) int {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.value = newValue
	return i.value
}
