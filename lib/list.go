package main

import "fmt"

type List struct {
	Elements []interface{}
	Length Value
	StartPos, EndPos *Position
}

func NewList(el []interface{}) *List {
	l := &List{Elements: el, Length: NewNumber(float64(len(el)))}
	return l
}

func (l *List) String() string {
	str := "["
	for i, el := range l.Elements {
		if i != 0 {
			str += ", "
		}
		str += el.(Value).String()
	}
	str += "]"
	return str
}

func (l *List) SetPos(sp, ep *Position) Value {
	l.StartPos = sp
	l.EndPos = ep
	if ep == nil {
		endPos := *sp
		endPos.Advance("")
		l.EndPos = &endPos
	}
	return l
}

func (l *List) AddTo(other interface{}) (Value, *Error) {
	if o, ok := other.(*List); ok {
		el := append(l.Elements, o.Elements...)
		return NewList(el), nil
	}
	return nil, NewInvalidSyntaxError("Only lists can be concatinated with a list", l.StartPos, l.EndPos)
}

func (l *List) SubBy(other interface{}) (Value, *Error) {
	return nil, NewInvalidSyntaxError("Invalid '-' operation on a list", l.StartPos, l.EndPos)
}

func (l *List) MulBy(other interface{}) (Value, *Error) {
	return nil, NewInvalidSyntaxError("Invalid '*' operation on a list", l.StartPos, l.EndPos)
}

func (l *List) DivBy(other interface{}) (Value, *Error) {
	return nil, NewInvalidSyntaxError("Invalid '/' operation on a list", l.StartPos, l.EndPos)
}

func (l *List) Mod(other interface{}) (Value, *Error) {
	return nil, NewInvalidSyntaxError("Invalid '%' operation on a list", l.StartPos, l.EndPos)
}

func (l *List) Pow(other interface{}) (Value, *Error) {
	return nil, NewInvalidSyntaxError("Invalid '^' operation on a list", l.StartPos, l.EndPos)
}

func (l *List) IsEqualTo(other interface{}) Value {
	return NewNumber(0)
}

func (l *List) IsNotEqualTo(other interface{}) Value {
	return NewNumber(1)
}

func (l *List) IsGreaterThan(other interface{}) (Value, *Error) {
	return nil, NewInvalidSyntaxError("Can't compare lists", l.StartPos, l.EndPos)
}

func (l *List) IsGreaterThanOrEqual(other interface{}) (Value, *Error) {
	return nil, NewRuntimeError("Can't compare lists", l.StartPos, nil)
}

func (l *List) IsLessThan(other interface{}) (Value, *Error) {
	return nil, NewRuntimeError("Can't compare lists", l.StartPos, nil)
}

func (l *List) IsLessThanOrEqual(other interface{}) (Value, *Error) {
	return nil, NewRuntimeError("Can't compare lists", l.StartPos, nil)
}

func (l *List) And(other interface{}) (Value, *Error) {
	if o, ok := other.(Value); ok {
		if o.IsTrue() {
			return o, nil
		}
		return NewNumber(0), nil
	}

	return nil, NewRuntimeError("Can't compare values of different types", l.StartPos, nil)
}

func (l *List) Or(other interface{}) (Value, *Error) {
	if o, ok := other.(Value); ok {
		if o.IsTrue() {
			return o, nil
		}
		return NewNumber(0), nil
	}

	return nil, NewRuntimeError("Can't compare values of different types", l.StartPos, nil)
}

func (l *List) Not() Value {
	return NewNumber(0)
}

func (l *List) IsTrue() bool {
	return true
}

func (l *List) GetVal() interface{} {
	return l.Elements
}

func (l *List) Call(args []interface{}, ctx *Context) *RuntimeResult {
	rr := NewRuntimeResult()

	if len(args) == 1 {
		if arg, ok := args[0].(*Number); ok {
			index := int(arg.Value)
			length := int(l.Length.GetVal().(float64))

			if length > index {
				return rr.Success(l.Elements[index].(Value))
			}

			return rr.Failure(NewRuntimeError(
				fmt.Sprintf("Index out of range (%v) with length of %v", index, length),
				l.StartPos, l.EndPos))
		} else {
			return rr.Failure(NewRuntimeError("Expected a number", l.StartPos, l.EndPos))
		}
	}

	if len(args) == 2 {
		if startArg, ok := args[0].(*Number); ok {
			if endArg, ok := args[1].(*Number); ok {
				start := int(startArg.Value)
				end := int(endArg.Value)
				length := int(l.Length.GetVal().(float64))
	
				if length <= start {
					return rr.Failure(NewRuntimeError(
						fmt.Sprintf("Index out of range (%v) with length of %v", start, length),
						l.StartPos, l.EndPos))
				}
				if length < end {
					return rr.Failure(NewRuntimeError(
						fmt.Sprintf("Index out of range (%v) with length of %v", end, length),
						l.StartPos, l.EndPos))	
				}

				el := l.Elements[start:end]

				return rr.Success(NewList(el))
			}
		}

		return rr.Failure(NewRuntimeError("Expected a number", l.StartPos, l.EndPos))
	}

	return rr.Failure(NewRuntimeError("Expected an index or start & end indexes", l.StartPos, l.EndPos))
}
