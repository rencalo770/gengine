package iparser

import parser "gengine/internal/iantlr/alr"

type GengineParserVisitor struct {
	parser.BasegengineVisitor
}

func NewGengineParserVisitor() *GengineParserVisitor {
	return &GengineParserVisitor{}
}
