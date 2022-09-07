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
// Example: [6, 2, 4, 1] -> [1, 2, 4, 6]
func (vec *Vector[T]) Sort() {
	v := reflect.TypeOf(vec.data)
	switch v.Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		vecInt := any(vec.data).([]int)
		if !sort.IntsAreSorted(vecInt) {
			sort.Ints(vecInt)
		}
	case reflect.Float32, reflect.Float64:
		vecFloat64 := any(vec.data).([]float64)
		if !sort.Float64sAreSorted(vecFloat64) {
			sort.Float64s(vecFloat64)
		}
	case reflect.String:
		vecString := any(vec.data).([]string)
		if !sort.StringsAreSorted(vecString) {
			sort.Strings(vecString)
		}
	default:
		log.Panicf("enable to sort vector of type %v\n", v.Elem().Kind())
	}
}
