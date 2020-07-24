package iparser

import parser "gengine/iantlr/alr"

type GengineParserVisitor struct {
	parser.BasegengineVisitor
}

func NewGengineParserVisitor() *GengineParserVisitor{
	return &GengineParserVisitor{}
}





