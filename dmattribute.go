package dmx

import (
	"fmt"
	"strconv"
)

type DmAttribute struct {
	name          string
	attributeType DmAttributeType
	value         interface{}
	owner         *DmElement
}

func newDmAttribute(name string, attributeType DmAttributeType, owner *DmElement) *DmAttribute {
	attribute := DmAttribute{owner: owner}
	attribute.SetName(name)
	attribute.SetType(attributeType)
	return &attribute
}

func (attribute *DmAttribute) GetName() string {
	return attribute.name
}

func (attribute *DmAttribute) SetName(name string) {
	attribute.name = name
}

func (attribute *DmAttribute) GetType() DmAttributeType {
	return attribute.attributeType
}

func (attribute *DmAttribute) SetType(attributeType DmAttributeType) {
	attribute.attributeType = attributeType
	switch attributeType {
	case AT_ELEMENT:
		attribute.value = nil
	case AT_INT:
		attribute.value = int32(0)
	case AT_FLOAT:
		attribute.value = float64(0)
	case AT_BOOL:
		attribute.value = false
	case AT_STRING:
		attribute.value = ""
	case AT_TIME:
		attribute.value = float64(0)
	case AT_COLOR:
		attribute.value = [...]byte{0, 0, 0, 0}
	case AT_VECTOR2:
		attribute.value = [...]float64{0, 0}
	case AT_VECTOR3:
		attribute.value = [...]float64{0, 0, 0}
	case AT_VECTOR4:
		attribute.value = [...]float64{0, 0, 0, 0}
	case AT_QANGLE:
		attribute.value = [...]float64{0, 0, 0}
	case AT_QUATERNION:
		attribute.value = [...]float64{0, 0, 0, 1}
	case AT_VMATRIX:
		attribute.value = [...]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}

	case AT_ELEMENT_ARRAY:
		attribute.value = make([]*DmElement, 0)
	case AT_INT_ARRAY:
		attribute.value = make([]int32, 0)
	case AT_FLOAT_ARRAY:
		attribute.value = make([]float64, 0)
	case AT_BOOL_ARRAY:
		attribute.value = make([]bool, 0)
	case AT_STRING_ARRAY:
		attribute.value = make([]string, 0)
	case AT_TIME_ARRAY:
		attribute.value = make([]float64, 0)
	case AT_COLOR_ARRAY:
		attribute.value = make([][4]byte, 0)
	case AT_VECTOR2_ARRAY:
		attribute.value = make([][2]float64, 0)
	case AT_VECTOR3_ARRAY, AT_QANGLE_ARRAY:
		attribute.value = make([][3]float64, 0)
	case AT_VECTOR4_ARRAY, AT_QUATERNION_ARRAY:
		attribute.value = make([][4]float64, 0)
	case AT_VMATRIX_ARRAY:
		attribute.value = make([][16]float64, 0)
	default:
		panic("Unknown attribute type in SetType " + type_to_string[attributeType])
	}
}

func (attribute *DmAttribute) GetValue() interface{} {
	return attribute.value
}

func (attribute *DmAttribute) SetValue(value interface{}) {
	//TODO: check if value is compatible with type
	attribute.value = value
}

func (attribute *DmAttribute) SetValues(value ...interface{}) {
	//TODO: check if value is compatible with type
	//attribute.value = value
}

func (attribute *DmAttribute) GetOwner() *DmElement {
	return attribute.owner
}

func (attribute *DmAttribute) StringValue() string {
	switch attribute.attributeType {
	case AT_INT:
		return strconv.FormatInt(int64(attribute.value.(int32)), 10)
	case AT_FLOAT, AT_TIME: // Time is stored as a float in txt version
		return strconv.FormatFloat(attribute.value.(float64), 'g', -1, 64)
	case AT_BOOL:
		if attribute.value.(bool) {
			return "1"
		} else {
			return "0"
		}
	case AT_STRING:
		return attribute.value.(string)
	case AT_COLOR:
		v := attribute.value.([4]byte)
		c := fmt.Sprintf("%d %d %d %d", v[0], v[1], v[2], v[3])
		return c
	case AT_VECTOR2:
		v := attribute.value.([2]float64)
		c := fmt.Sprintf("%g %g", v[0], v[1])
		return c
	case AT_VECTOR3, AT_QANGLE:
		v := attribute.value.([3]float64)
		c := fmt.Sprintf("%g %g %g", v[0], v[1], v[2])
		return c
	case AT_VECTOR4, AT_QUATERNION:
		v := attribute.value.([4]float64)
		c := fmt.Sprintf("%g %g %g %g", v[0], v[1], v[2], v[3])
		return c
	case AT_VMATRIX:
		v := attribute.value.([16]float64)
		c := fmt.Sprintf("%g %g %g %g %g %g %g %g %g %g %g %g %g %g %g %g", v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11], v[12], v[13], v[14], v[15])
		return c
	default:
		panic("Unknown attribute type in StringValue " + type_to_string[attribute.attributeType])
	}
	return ""
}

func (attribute *DmAttribute) PushElement(element *DmElement) {
	a := attribute.value.([]*DmElement)
	attribute.value = append(a, element)
}

func (attribute *DmAttribute) PushInt(i int32) {
	a := attribute.value.([]int32)
	attribute.value = append(a, i)
}

func (attribute *DmAttribute) PushFloat(f float64) {
	a := attribute.value.([]float64)
	attribute.value = append(a, f)
}

func (attribute *DmAttribute) PushBool(b bool) {
	a := attribute.value.([]bool)
	attribute.value = append(a, b)
}

func (attribute *DmAttribute) PushString(s string) {
	a := attribute.value.([]string)
	attribute.value = append(a, s)
}

func (attribute *DmAttribute) PushTime(t float64) {
	a := attribute.value.([]float64)
	attribute.value = append(a, t)
}

func (attribute *DmAttribute) PushColor(v [4]byte) {
	a := attribute.value.([][4]byte)
	attribute.value = append(a, v)
}

func (attribute *DmAttribute) PushVector2(v [2]float64) {
	a := attribute.value.([][2]float64)
	attribute.value = append(a, v)
}

func (attribute *DmAttribute) PushVector3(v [3]float64) {
	a := attribute.value.([][3]float64)
	attribute.value = append(a, v)
}

func (attribute *DmAttribute) PushVector4(v [4]float64) {
	a := attribute.value.([][4]float64)
	attribute.value = append(a, v)
}
