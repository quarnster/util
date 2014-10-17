// Copyright 2013 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

// This file was generated with the following command:
// ["/Users/quarnster/code/go/bin/pegparser", "-peg=encoding/binary/expression/expression.peg", "-notest", "-ignore=Spacing,Primary,Op,Expression,Grouping,BooleanOp", "-testfile=", "-outpath", "encoding/binary/expression/", "-generator=go"]

package expression

import (
	"github.com/limetext/text"
	. "github.com/quarnster/parser"
)

type EXPRESSION struct {
	ParserData  Reader
	IgnoreRange text.Region
	Root        Node
	LastError   int
}

func (p *EXPRESSION) RootNode() *Node {
	return &p.Root
}

func (p *EXPRESSION) SetData(data string) {
	p.ParserData = NewReader(data)
	p.Root = Node{Name: "EXPRESSION", P: p}
	p.IgnoreRange = text.Region{}
	p.LastError = 0
}

func (p *EXPRESSION) Parse(data string) bool {
	p.SetData(data)
	ret := p.realParse()
	p.Root.UpdateRange()
	return ret
}

func (p *EXPRESSION) Data(start, end int) string {
	return p.ParserData.Substring(start, end)
}

func (p *EXPRESSION) Error() Error {
	errstr := ""
	line, column := p.ParserData.LineCol(p.LastError)

	if p.LastError == p.ParserData.Len() {
		errstr = "Unexpected EOF"
	} else {
		p.ParserData.Seek(p.LastError)
		if r := p.ParserData.Read(); r == '\r' || r == '\n' {
			errstr = "Unexpected new line"
		} else {
			errstr = "Unexpected " + string(r)
		}
	}
	return NewError(line, column, errstr)
}

func (p *EXPRESSION) realParse() bool {
	return p.Expression()
}
func (p *EXPRESSION) Expression() bool {
	// Expression      <-      (Op / Grouping) EndOfFile
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			save := p.ParserData.Pos()
			accept = p.Op()
			if !accept {
				accept = p.Grouping()
				if !accept {
				}
			}
			if !accept {
				p.ParserData.Seek(save)
			}
		}
		if accept {
			accept = p.EndOfFile()
			if accept {
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *EXPRESSION) Op() bool {
	// Op              <-      ShiftRight / ShiftLeft / AndNot / Mask / Add / Sub / Mul / BooleanOp
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.ShiftRight()
		if !accept {
			accept = p.ShiftLeft()
			if !accept {
				accept = p.AndNot()
				if !accept {
					accept = p.Mask()
					if !accept {
						accept = p.Add()
						if !accept {
							accept = p.Sub()
							if !accept {
								accept = p.Mul()
								if !accept {
									accept = p.BooleanOp()
									if !accept {
									}
								}
							}
						}
					}
				}
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *EXPRESSION) BooleanOp() bool {
	// BooleanOp       <-      Eq / Lt / Gt / Le / Ge / Ne
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Eq()
		if !accept {
			accept = p.Lt()
			if !accept {
				accept = p.Gt()
				if !accept {
					accept = p.Le()
					if !accept {
						accept = p.Ge()
						if !accept {
							accept = p.Ne()
							if !accept {
							}
						}
					}
				}
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *EXPRESSION) ShiftRight() bool {
	// ShiftRight      <-      Grouping ">>" Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			{
				accept = true
				s := p.ParserData.Pos()
				if p.ParserData.Read() != '>' || p.ParserData.Read() != '>' {
					p.ParserData.Seek(s)
					accept = false
				}
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "ShiftRight"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) ShiftLeft() bool {
	// ShiftLeft       <-      Grouping "<<" Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			{
				accept = true
				s := p.ParserData.Pos()
				if p.ParserData.Read() != '<' || p.ParserData.Read() != '<' {
					p.ParserData.Seek(s)
					accept = false
				}
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "ShiftLeft"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Mask() bool {
	// Mask            <-      Grouping '&' Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			if p.ParserData.Read() != '&' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Mask"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Add() bool {
	// Add             <-      Grouping '+' Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			if p.ParserData.Read() != '+' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Add"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Sub() bool {
	// Sub             <-      Grouping '-' Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			if p.ParserData.Read() != '-' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Sub"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Mul() bool {
	// Mul             <-      Grouping '*' Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			if p.ParserData.Read() != '*' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Mul"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) AndNot() bool {
	// AndNot          <-      Grouping "&^" Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			{
				accept = true
				s := p.ParserData.Pos()
				if p.ParserData.Read() != '&' || p.ParserData.Read() != '^' {
					p.ParserData.Seek(s)
					accept = false
				}
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "AndNot"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Eq() bool {
	// Eq              <-      Grouping "==" Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			{
				accept = true
				s := p.ParserData.Pos()
				if p.ParserData.Read() != '=' || p.ParserData.Read() != '=' {
					p.ParserData.Seek(s)
					accept = false
				}
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Eq"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Ne() bool {
	// Ne              <-      Grouping "!=" Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			{
				accept = true
				s := p.ParserData.Pos()
				if p.ParserData.Read() != '!' || p.ParserData.Read() != '=' {
					p.ParserData.Seek(s)
					accept = false
				}
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Ne"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Lt() bool {
	// Lt              <-      Grouping '<' Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			if p.ParserData.Read() != '<' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Lt"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Le() bool {
	// Le              <-      Grouping "<=" Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			{
				accept = true
				s := p.ParserData.Pos()
				if p.ParserData.Read() != '<' || p.ParserData.Read() != '=' {
					p.ParserData.Seek(s)
					accept = false
				}
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Le"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Gt() bool {
	// Gt              <-      Grouping '>' Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			if p.ParserData.Read() != '>' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Gt"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Ge() bool {
	// Ge              <-      Grouping ">=" Grouping
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Grouping()
		if accept {
			{
				accept = true
				s := p.ParserData.Pos()
				if p.ParserData.Read() != '>' || p.ParserData.Read() != '=' {
					p.ParserData.Seek(s)
					accept = false
				}
			}
			if accept {
				accept = p.Grouping()
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Ge"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Grouping() bool {
	// Grouping        <-      Spacing? ('(' Op ')' / Constant / DotIdentifier) Spacing?
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Spacing()
		accept = true
		if accept {
			{
				save := p.ParserData.Pos()
				{
					save := p.ParserData.Pos()
					if p.ParserData.Read() != '(' {
						p.ParserData.UnRead()
						accept = false
					} else {
						accept = true
					}
					if accept {
						accept = p.Op()
						if accept {
							if p.ParserData.Read() != ')' {
								p.ParserData.UnRead()
								accept = false
							} else {
								accept = true
							}
							if accept {
							}
						}
					}
					if !accept {
						if p.LastError < p.ParserData.Pos() {
							p.LastError = p.ParserData.Pos()
						}
						p.ParserData.Seek(save)
					}
				}
				if !accept {
					accept = p.Constant()
					if !accept {
						accept = p.DotIdentifier()
						if !accept {
						}
					}
				}
				if !accept {
					p.ParserData.Seek(save)
				}
			}
			if accept {
				accept = p.Spacing()
				accept = true
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *EXPRESSION) DotIdentifier() bool {
	// DotIdentifier   <-      Identifier ('.' Identifier)*
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Identifier()
		if accept {
			{
				accept = true
				for accept {
					{
						save := p.ParserData.Pos()
						if p.ParserData.Read() != '.' {
							p.ParserData.UnRead()
							accept = false
						} else {
							accept = true
						}
						if accept {
							accept = p.Identifier()
							if accept {
							}
						}
						if !accept {
							if p.LastError < p.ParserData.Pos() {
								p.LastError = p.ParserData.Pos()
							}
							p.ParserData.Seek(save)
						}
					}
				}
				accept = true
			}
			if accept {
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "DotIdentifier"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Identifier() bool {
	// Identifier      <-      [A-Z] [_A-Za-z0-9]*
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		c := p.ParserData.Read()
		if c >= 'A' && c <= 'Z' {
			accept = true
		} else {
			p.ParserData.UnRead()
			accept = false
		}
		if accept {
			{
				accept = true
				for accept {
					{
						save := p.ParserData.Pos()
						c := p.ParserData.Read()
						if c >= 'A' && c <= 'Z' {
							accept = true
						} else {
							p.ParserData.UnRead()
							accept = false
						}
						if !accept {
							c := p.ParserData.Read()
							if c >= 'a' && c <= 'z' {
								accept = true
							} else {
								p.ParserData.UnRead()
								accept = false
							}
							if !accept {
								c := p.ParserData.Read()
								if c >= '0' && c <= '9' {
									accept = true
								} else {
									p.ParserData.UnRead()
									accept = false
								}
								if !accept {
									{
										accept = false
										c := p.ParserData.Read()
										if c == '_' {
											accept = true
										} else {
											p.ParserData.UnRead()
										}
									}
									if !accept {
									}
								}
							}
						}
						if !accept {
							p.ParserData.Seek(save)
						}
					}
				}
				accept = true
			}
			if accept {
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Identifier"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Constant() bool {
	// Constant        <-      ("0x" [a-fA-F0-9]+) / [0-9]+
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			save := p.ParserData.Pos()
			{
				accept = true
				s := p.ParserData.Pos()
				if p.ParserData.Read() != '0' || p.ParserData.Read() != 'x' {
					p.ParserData.Seek(s)
					accept = false
				}
			}
			if accept {
				{
					save := p.ParserData.Pos()
					{
						save := p.ParserData.Pos()
						c := p.ParserData.Read()
						if c >= 'a' && c <= 'f' {
							accept = true
						} else {
							p.ParserData.UnRead()
							accept = false
						}
						if !accept {
							c := p.ParserData.Read()
							if c >= 'A' && c <= 'F' {
								accept = true
							} else {
								p.ParserData.UnRead()
								accept = false
							}
							if !accept {
								c := p.ParserData.Read()
								if c >= '0' && c <= '9' {
									accept = true
								} else {
									p.ParserData.UnRead()
									accept = false
								}
								if !accept {
								}
							}
						}
						if !accept {
							p.ParserData.Seek(save)
						}
					}
					if !accept {
						p.ParserData.Seek(save)
					} else {
						for accept {
							{
								save := p.ParserData.Pos()
								c := p.ParserData.Read()
								if c >= 'a' && c <= 'f' {
									accept = true
								} else {
									p.ParserData.UnRead()
									accept = false
								}
								if !accept {
									c := p.ParserData.Read()
									if c >= 'A' && c <= 'F' {
										accept = true
									} else {
										p.ParserData.UnRead()
										accept = false
									}
									if !accept {
										c := p.ParserData.Read()
										if c >= '0' && c <= '9' {
											accept = true
										} else {
											p.ParserData.UnRead()
											accept = false
										}
										if !accept {
										}
									}
								}
								if !accept {
									p.ParserData.Seek(save)
								}
							}
						}
						accept = true
					}
				}
				if accept {
				}
			}
			if !accept {
				if p.LastError < p.ParserData.Pos() {
					p.LastError = p.ParserData.Pos()
				}
				p.ParserData.Seek(save)
			}
		}
		if !accept {
			{
				save := p.ParserData.Pos()
				c := p.ParserData.Read()
				if c >= '0' && c <= '9' {
					accept = true
				} else {
					p.ParserData.UnRead()
					accept = false
				}
				if !accept {
					p.ParserData.Seek(save)
				} else {
					for accept {
						c := p.ParserData.Read()
						if c >= '0' && c <= '9' {
							accept = true
						} else {
							p.ParserData.UnRead()
							accept = false
						}
					}
					accept = true
				}
			}
			if !accept {
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Constant"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *EXPRESSION) Spacing() bool {
	// Spacing         <-      [ \t\n\r]+
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			accept = false
			c := p.ParserData.Read()
			if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
				accept = true
			} else {
				p.ParserData.UnRead()
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		} else {
			for accept {
				{
					accept = false
					c := p.ParserData.Read()
					if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
						accept = true
					} else {
						p.ParserData.UnRead()
					}
				}
			}
			accept = true
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *EXPRESSION) EndOfFile() bool {
	// EndOfFile       <-      !.
	accept := false
	accept = true
	start := p.ParserData.Pos()
	s := p.ParserData.Pos()
	if p.ParserData.Pos() >= p.ParserData.Len() {
		accept = false
	} else {
		p.ParserData.Read()
		accept = true
	}
	p.ParserData.Seek(s)
	p.Root.Discard(s)
	accept = !accept
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "EndOfFile"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}
