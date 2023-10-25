package parser

import "github.com/cuongpiger/joat/parser/entity"

// ************************************************* PUBLIC FUNCTIONS **************************************************

func GetParser() (IParser, error) {
	return newParser()
}

// ************************************************* PRIVATE FUNCTIONS *************************************************

func newParser() (IParser, error) {
	parserOnce.Do(func() {
		parserIns = new(entity.Parser)
	})

	// currently, there is no error to return
	return parserIns, nil
}
