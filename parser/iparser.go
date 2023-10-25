package parser

type IParser interface {
	MapMe(pObj interface{}, pParent string) (map[string]interface{}, error)
}
