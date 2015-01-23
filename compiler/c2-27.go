package main

import (
	"bufio"
	"os"
)

type Parser struct {
	lookahead byte
	reader    *bufio.Reader
}

func (p *Parser) Init(f *os.File) {
	p.reader = bufio.NewReader(f)
	lookahead, err := p.reader.ReadByte()
	if err != nil {
		panic("cannot read byte")
	}
	p.lookahead = lookahead
}

func (p *Parser) expr() {
	p.term()
	for {
		if p.lookahead == '+' {
			p.match('+')
			p.term()
			print("+")
		} else if p.lookahead == '-' {
			p.match('-')
			p.term()
			print("-")
		} else {
			return
		}
	}
}

func (p *Parser) term() {
	if p.lookahead >= 48 && p.lookahead <= 57 {
		print(string(p.lookahead))
		p.match(p.lookahead)
	}
}

func (p *Parser) match(t byte) {
	if p.lookahead == t {
		lookahead, err := p.reader.ReadByte()
		if err != nil {
			panic("cannot read byte")
		}
		p.lookahead = lookahead
	} else {
		panic("syntax error")
	}
}

func main() {
	p := Parser{}
	p.Init(os.Stdin)
	p.expr()
	println()
}
