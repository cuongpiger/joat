package parser

import "net/url"

type IParser interface {
	UrlMe(pObj interface{}) (*url.URL, error)
	MapMe(pObj interface{}, pParent string) (map[string]interface{}, error)
}
