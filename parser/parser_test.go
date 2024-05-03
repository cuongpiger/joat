package parser

import (
	"testing"
)

type (
	ListOpts struct {
		Name string `q:"name,beempty"`
		Page int    `q:"page"`
		Size int    `q:"size"`

		Tags map[string]string `q:"tags"`
	}
)

func TestUrlMe(t *testing.T) {
	_, _ = GetParser()

}
