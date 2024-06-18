package dmx

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
		attribute.value = int64(0)
	case AT_FLOAT:
		attribute.value = float64(0)
	case AT_BOOL:
		attribute.value = false
	case AT_STRING:
		attribute.value = ""
	case AT_OBJECTID:
		attribute.value = new(DmObjectId)
	case AT_COLOR:
		attribute.value = [...]float64{0, 0, 0, 0}
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
		attribute.value = make([]int64, 0)
	case AT_FLOAT_ARRAY:
		attribute.value = make([]float64, 0)
	case AT_BOOL_ARRAY:
		attribute.value = make([]bool, 0)
	case AT_STRING_ARRAY:
		attribute.value = make([]string, 0)
	case AT_OBJECTID_ARRAY:
		attribute.value = make([]*DmObjectId, 0)
	case AT_COLOR_ARRAY:
		attribute.value = make([][]float64, 0)
	case AT_VECTOR2_ARRAY:
		attribute.value = make([][]float64, 0)
	case AT_VECTOR3_ARRAY:
		attribute.value = make([][]float64, 0)
	case AT_VECTOR4_ARRAY:
		attribute.value = make([][]float64, 0)
	case AT_QANGLE_ARRAY:
		attribute.value = make([][]float64, 0)
	case AT_QUATERNION_ARRAY:
		attribute.value = make([][]float64, 0)
	case AT_VMATRIX_ARRAY:
		attribute.value = make([][]float64, 0)
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
