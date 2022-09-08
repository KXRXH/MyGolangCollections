package myCollections

import (
	"fmt"
	"log"
	"reflect"
)

// Pair is two objects with the same types.
type Pair[T any] struct {
	first, second T
}

// NewPair creates a new Pair.
//
// Usage:
//
// p1 := NewPair[type](a, b)
//
// p2 := NewPair[type]([]type{a, b}...)
//
// p3 := NewPair[type]()
func NewPair[T any](args ...T) *Pair[T] {
	var newPair *Pair[T]
	switch len(args) {
	case 2:
		newPair = &Pair[T]{first: args[0], second: args[1]}
	case 0:
		newPair = &Pair[T]{first: *new(T), second: *new(T)}
	default:
		log.Panicf("unable to create a pair by %d value(s), expected 0 or 2.\n", len(args))
	}
	return newPair
}

func (p *Pair[T]) SetFirst(value T) {
	p.first = value
}

func (p *Pair[T]) SetSecond(value T) {
	p.second = value
}

// Swap swaps first and second object of the Pair
//
// Example: {5, 6} -> {6, 5}
func (p *Pair[T]) Swap() {
	p.first, p.second = p.second, p.first
}

func (p *Pair[T]) Copy() *Pair[T] {
	return &Pair[T]{first: p.first, second: p.second}
}

func (p *Pair[T]) First() T {
	return p.first
}

func (p *Pair[T]) Second() T {
	return p.second
}

// ToString returns a string representation of the Pair.
func (p *Pair[T]) ToString() string {
	return fmt.Sprintf("pair{%v %v}", p.first, p.second)
}

func (p *Pair[T]) Sum() T {
	var sum T
	// Convert sum to the type T
	toTypeT := func(i any) T { return reflect.ValueOf(i).Interface().(T) }
	t := reflect.TypeOf(p.first).Kind()
	switch t {
	case reflect.Int:
		sum = toTypeT(reflect.ValueOf(p.first).Interface().(int) + reflect.ValueOf(p.second).Interface().(int))
	case reflect.Int8:
		sum = toTypeT(reflect.ValueOf(p.first).Interface().(int8) + reflect.ValueOf(p.second).Interface().(int8))
	case reflect.Int32:
		sum = toTypeT(reflect.ValueOf(p.first).Interface().(int32) + reflect.ValueOf(p.second).Interface().(int32))
	case reflect.Int64:
		sum = toTypeT(reflect.ValueOf(p.first).Interface().(int64) + reflect.ValueOf(p.second).Interface().(int64))
	case reflect.Float64:
		sum = toTypeT(reflect.ValueOf(p.first).Interface().(float64) + reflect.ValueOf(p.second).Interface().(float64))
	case reflect.Float32:
		sum = toTypeT(reflect.ValueOf(p.first).Interface().(float32) + reflect.ValueOf(p.second).Interface().(float32))
	default:
		log.Fatalf("unable to sum the pair of type %v\n", t.String())
	}
	return sum
}
