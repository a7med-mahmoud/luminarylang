package main

import (
	"fmt"
	"math"
)

type Number struct {
	Value float64
	Pos *Position
}

func NewNumber(v float64) *Number {
	n := &Number{Value: v}

	return n
}

func (n *Number) String() string {
	return fmt.Sprintf("%v", n.Value);
}

func (n *Number) SetPos(p *Position) *Number {
	n.Pos = p
	return n
}

func (n *Number) AddTo(other interface{}) *Number {
	if o, ok := other.(*Number); ok {
		return NewNumber(n.Value + o.Value)
	}
	return nil
}

func (n *Number) SubBy(other interface{}) *Number {
	if o, ok := other.(*Number); ok {
		return NewNumber(n.Value - o.Value)
	}
	return nil
}

func (n *Number) MulBy(other interface{}) *Number {
	if o, ok := other.(*Number); ok {
		return NewNumber(n.Value * o.Value)
	}
	return nil
}

func (n *Number) DivBy(other interface{}) *Number {
	switch o := other.(type) {
	case *Number:
		return NewNumber(n.Value / o.Value)
	}
	return nil
}

func (n *Number) Mod(other interface{}) *Number {
	switch o := other.(type) {
	case *Number:
		return NewNumber(math.Mod(n.Value, o.Value))
	}
	return nil
}

func (n *Number) Pow(other interface{}) *Number {
	switch o := other.(type) {
	case *Number:
		return NewNumber(math.Pow(n.Value, o.Value))
	}
	return nil
}
