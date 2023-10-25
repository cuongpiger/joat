package parser

import "github.com/cuongpiger/joat/parser/entity"

func GetParser() (IParser, error) {
	return newParser()
}

func newParser() (IParser, error) {
	parserOnce.Do(func() {
		parserIns = new(entity.Parser)
	})

	// currently, there is no error to return
	return parserIns, nil
}
