package entity

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	lError "github.com/cuongpiger/joat/error"
)

type Parser struct {
}

func (s *Parser) MapMe(pObj interface{}, parent string) (map[string]interface{}, error) {
	objValue := reflect.ValueOf(pObj)

	// if the pObj is a pointer, get the underlying value
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	// if the pObj is a struct, get the underlying type
	objType := reflect.TypeOf(pObj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	objMap := make(map[string]interface{})
	if objValue.Kind() == reflect.Struct {
		for i := 0; i < objValue.NumField(); i++ {
			v := objValue.Field(i)
			f := objType.Field(i)

			if f.Name != cases.Title(language.Und, cases.NoLower).String(f.Name) {
				continue
			}

			zero := isZero(v)

			// if the field has a required tag that's set to "true"
			if requiredTag := f.Tag.Get("required"); requiredTag == "true" {
				// if the field's value is zero, return a missing-argument error
				if zero {
					// if the field has a 'required' tag, it can't have a zero-value
					return nil, lError.NewErrMissingInput(f.Name, "")
				}
			}

			if xorTag := f.Tag.Get("xor"); xorTag != "" {
				xorField := objValue.FieldByName(xorTag)
				var xorFieldIsZero bool
				if reflect.ValueOf(xorField.Interface()) == reflect.Zero(xorField.Type()) {
					xorFieldIsZero = true
				} else {
					if xorField.Kind() == reflect.Ptr {
						xorField = xorField.Elem()
					}
					xorFieldIsZero = isZero(xorField)
				}
				if !(zero != xorFieldIsZero) {
					return nil, lError.NewErrMissingInput(
						fmt.Sprintf("%s/%s", f.Name, xorTag),
						fmt.Sprintf("exactly one of %s and %s must be provided", f.Name, xorTag))
				}
			}

			if orTag := f.Tag.Get("or"); orTag != "" {
				if zero {
					orField := objValue.FieldByName(orTag)
					var orFieldIsZero bool
					if reflect.ValueOf(orField.Interface()) == reflect.Zero(orField.Type()) {
						orFieldIsZero = true
					} else {
						if orField.Kind() == reflect.Ptr {
							orField = orField.Elem()
						}
						orFieldIsZero = isZero(orField)
					}
					if orFieldIsZero {
						return nil, lError.NewErrMissingInput(
							fmt.Sprintf("%s/%s", f.Name, orTag),
							fmt.Sprintf("at least one of %s and %s must be provided", f.Name, orTag))
					}
				}
			}

			jsonTag := f.Tag.Get("json")
			if jsonTag == "-" {
				continue
			}

			if v.Kind() == reflect.Slice || (v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Slice) {
				sliceValue := v
				if sliceValue.Kind() == reflect.Ptr {
					sliceValue = sliceValue.Elem()
				}

				for i := 0; i < sliceValue.Len(); i++ {
					element := sliceValue.Index(i)
					if element.Kind() == reflect.Struct || (element.Kind() == reflect.Ptr && element.Elem().Kind() == reflect.Struct) {
						_, err := s.MapMe(element.Interface(), f.Name)
						if err != nil {
							return nil, err
						}
					}
				}
			}

			if v.Kind() == reflect.Struct || (v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct) {
				if zero {
					if jsonTag != "" {
						jsonTagPieces := strings.Split(jsonTag, ",")
						if len(jsonTagPieces) > 1 && jsonTagPieces[1] == "omitempty" {
							if v.CanSet() {
								if !v.IsNil() {
									if v.Kind() == reflect.Ptr {
										v.Set(reflect.Zero(v.Type()))
									}
								}
							}
						}
					}
					continue
				}

				_, err := s.MapMe(v.Interface(), f.Name)
				if err != nil {
					return nil, err
				}
			}
		}

		b, err := json.Marshal(pObj)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(b, &objMap)
		if err != nil {
			return nil, err
		}

		if parent != "" {
			objMap = map[string]interface{}{parent: objMap}
		}

		return objMap, nil
	}

	// return an error if the underlying type of 'opts' isn't a struct
	return nil, fmt.Errorf("options type is not a struct")
}
