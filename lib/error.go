package main

import "fmt"

type Error struct {
	Name, Details string
	StartPos *Position
	EndPos *Position
}

func NewError(n, d string, sp, ep *Position) *Error {
	e := &Error{
		Name: n,
		Details: d,
		StartPos: sp,
		EndPos: ep,
	}

	if sp != nil && ep == nil {
		endPos := *sp
		endPos.Advance("")
		e.EndPos = &endPos
	}

	return e
}

func (e *Error) String() string {
	if e.StartPos != nil && e.EndPos != nil {
		return fmt.Sprintf(
			"%vError(%v): %v.\nFile: %v - Line: %v - Col: %v:%v",
			"\033[31m",
			e.Name,
			e.Details,
			e.StartPos.FileName,
			e.StartPos.Line,
			e.StartPos.Col,
			e.EndPos.Col)
	}
	return fmt.Sprintf(
		"%vError(%v): %v.",
		"\033[31m",
		e.Name,
		e.Details)
}

func NewIlligalCharError(d string, sp, ep *Position) *Error {
	e := NewError("Illigal Char", d, sp, ep)
	return e
}

func NewInvalidSyntaxError(d string, sp, ep *Position) *Error {
	e := NewError("Invalid Syntax", d, sp, ep)
	return e
}

func NewRuntimeError(d string, sp, ep *Position) *Error {
	e := NewError("Runtime Error", d, sp, ep)
	return e
}
