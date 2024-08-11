package memento

// Memento is a generic type that represents a memento object. It stores the state of an object and can be used to restore that state later. The state is of type T, which can be any type specified during instantiation.
type Memento[T any] struct {
	State T
}

// Originator is a generic type that represents an object with a state.
// It has a single field called State of type T, which holds the current state of the object.
// Originator provides methods to apply changes to the state, save the current state to a Memento,
// and restore the state from a Memento.
type Originator[T any] struct {
	State T
}

// ApplyChange sets the state of the Originator to the new state provided as a parameter.
// The new state is assigned to the "State" field of the Originator.
func (o *Originator[T]) ApplyChange(newState T) {
	o.State = newState
}

// SaveToMemento saves the current state of the Originator object to a Memento object.
// It creates a new Memento object with the same state as the current state of the Originator
// and returns it.
// The Memento object stores the state of the Originator and can be used later to restore that state.
// The state is assigned to the "State" field of the Memento.
func (o *Originator[T]) SaveToMemento() Memento[T] {
	return Memento[T]{State: o.State}
}

// RestoreFromMemento restores the state of the Originator object from a Memento object.
// The state of the Originator is assigned the value of the "State" field of the Memento.
func (o *Originator[T]) RestoreFromMemento(m Memento[T]) {
	o.State = m.State
}

// Caretaker is a generic type that represents a caretaker object. It is responsible for storing and managing
// a collection of memento objects. The type parameter T represents the type of the state stored in the mementos.
type Caretaker[T any] struct {
	mementos []Memento[T]
}

// AddMemento adds a Memento object to the caretaker's list of mementos.
// The provided Memento object is appended to the "mementos" slice of the caretaker.
// The "mementos" slice stores the state of the objects and can be used to restore
// the state later.
func (c *Caretaker[T]) AddMemento(m Memento[T]) {
	c.mementos = append(c.mementos, m)
}

// GetMemento retrieves a memento object from the caretaker's mementos list based on the given index.
// The method returns the memento object at the specified index.
// If the index is out of range, the behavior is undefined.
// Note: The index starts from 0.
func (c *Caretaker[T]) GetMemento(index int) Memento[T] {
	return c.mementos[index]
}
