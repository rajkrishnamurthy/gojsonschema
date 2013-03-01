// @author       sigu-399
// @description  An implementation of JSON Schema, draft v4 - Go language
// @created      27-02-2013

package gojsonschema

import (
	"errors"
	"gojsonreference"
	"regexp"
)

type JsonSchema struct {
	id          *string
	title       *string
	description *string
	types       JsonSchemaType

	ref *gojsonreference.JsonReference

	definitionsChildren []*JsonSchema
	itemsChild          *JsonSchema
	propertiesChildren  []*JsonSchema

	parent *JsonSchema

	property string

	schema *gojsonreference.JsonReference

	// validation : number / integer
	multipleOf       *float64
	maximum          *float64
	exclusiveMaximum bool
	minimum          *float64
	exclusiveMinimum bool

	// validation : string
	minLength *int
	maxLength *int
	pattern   *regexp.Regexp

	// validation : object
	minProperties *int
	maxProperties *int

	required []string
}

func (s *JsonSchema) AddRequired(value string) error {

	if isStringInSlice(s.required, value) {
		return errors.New("required items must be unique")
	}

	s.required = append(s.required, value)

	return nil
}

func (s *JsonSchema) AddDefinitionChild(child *JsonSchema) {
	s.definitionsChildren = append(s.definitionsChildren, child)
}

func (s *JsonSchema) SetItemsChild(child *JsonSchema) {
	s.itemsChild = child
}

func (s *JsonSchema) AddPropertiesChild(child *JsonSchema) {
	s.propertiesChildren = append(s.propertiesChildren, child)
}

func (s *JsonSchema) HasProperty(name string) bool {

	for _, v := range s.propertiesChildren {
		if v.property == name {
			return true
		}
	}
	return false
}
