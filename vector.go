package myCollections

import (
	"fmt"
	"log"
	"reflect"
	"sort"
)

type Vector[T any] struct {
	data []T
	size int
}

// NewVector creates a new empty Vector of type T
func NewVector[T any]() *Vector[T] {
	newVector := Vector[T]{data: make([]T, 0), size: 0}
	return &newVector
}

// Clear deletes all data of the Vector.
func (vec *Vector[T]) Clear() {
	vec.data = make([]T, 0)
	vec.size = 0
}

// PushBack adds a new object to the back of the Vector.
func (vec *Vector[T]) PushBack(object T) {
	newData := make([]T, vec.size+1)
	copy(newData, vec.data)
	newData[vec.size] = object
	vec.data = newData
	vec.size++
}

// DeleteBack deletes the last element of the Vector.
func (vec *Vector[T]) DeleteBack() {
	newData := make([]T, vec.size-1)
	copy(newData, vec.data)
	vec.data = newData
	vec.size--
}

// ToString returns a string representation of the Vector.
func (vec *Vector[T]) ToString() string {
	return fmt.Sprintf("vector(%d): %v", vec.size, vec.data)
}

// Get return value of Vector by index.
func (vec *Vector[T]) Get(index int) T {
	if index < vec.size {
		value := vec.data[index]
		return value
	}
	return *new(T)
}

// Len returns the number of elements in the Vector
func (vec *Vector[T]) Len() int {
	return vec.size
}

// Pop deletes element of Vector by index and returns value of the deleted element.
//
// Returns: element of type T
// If index is out of range, returns *Type
func (vec *Vector[T]) Pop(index int) T {
	if index < vec.size {
		vec.data = append(vec.data[:index], vec.data[index+1:]...)
		vec.size--
		return vec.data[index]
	}
	return *new(T)
}

// Reverse flips the Vector.
//
// Example: [1, 2, 3, 4] -> [4, 3, 2, 1]
func (vec *Vector[T]) Reverse() {
	for index := 0; index < vec.size/2; index++ {
		vec.data[index], vec.data[vec.size-1-index] = vec.data[vec.size-1-index], vec.data[index]
	}
}

// Sort sorts the Vector by ascending.
//
// Supported types: Int, Float, String
//
// Also you can pass sort function as the parameter.
//
// Example: [6, 2, 4, 1] -> [1, 2, 4, 6]
func (vec *Vector[T]) Sort(sortFunc ...func(i, j int) bool) {
	if len(sortFunc) == 1 {
		sort.SliceStable(vec.data, sortFunc[0])
		return
	}
	t := reflect.TypeOf(vec.data).Elem().Kind()
	switch t {
	case reflect.Int:
		sort.Ints(reflect.ValueOf(vec.data).Interface().([]int))
	case reflect.Float64:
		sort.Float64s(reflect.ValueOf(vec.data).Interface().([]float64))
	case reflect.String:
		sort.Strings(reflect.ValueOf(vec.data).Interface().([]string))
	default:
		log.Fatalf("enable to sort the vector of the type %v. try to pass sort function as the parameter.\n", t)
	}
}
