package set

// Set представляет собой структуру множества.
type Set[T comparable] struct {
	items map[T]struct{}
}

// New создает новое множество и инициализирует его элементами.
func New[T comparable](elements ...T) *Set[T] {
	s := &Set[T]{items: make(map[T]struct{})}
	for _, element := range elements {
		s.Add(element)
	}
	return s
}

// Add добавляет элемент в множество.
func (s *Set[T]) Add(element T) {
	s.items[element] = struct{}{}
}

// Remove удаляет элемент из множества.
func (s *Set[T]) Remove(element T) {
	delete(s.items, element)
}

// Contains проверяет, содержится ли элемент в множестве.
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.items[element]
	return exists
}

// Size возвращает количество элементов в множестве.
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Clear очищает множество.
func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

// Elements возвращает все элементы множества в виде среза.
func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, len(s.items))
	for item := range s.items {
		elements = append(elements, item)
	}
	return elements
}
