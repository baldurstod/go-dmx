package dmx

type DmElement struct {
	id          DmObjectId
	elementType string
	attributes  map[string]*DmAttribute
}

func NewDmElement(elementType string) *DmElement {
	return &DmElement{
		id:          CreateObjectId(),
		elementType: elementType,
		attributes:  map[string]*DmAttribute{},
	}
}

func (element *DmElement) CreateAttribute(name string, attributeType DmAttributeType) *DmAttribute {
	attribute, exist := element.attributes[name]

	if exist {
		if attribute.attributeType == attributeType {
			return attribute
		}
		return nil
	}

	attribute = newDmAttribute(name, attributeType, element)
	element.attributes[name] = attribute

	return attribute
}

func (element *DmElement) GetId() DmObjectId {
	return element.id
}

func (element *DmElement) SetId(id DmObjectId) {
	element.id = id
}

func (element *DmElement) GetType() string {
	return element.elementType
}

func (element *DmElement) SetType(elementType string) {
	element.elementType = elementType
}

func (element *DmElement) CreateElementAttribute(name string, value *DmElement) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_ELEMENT)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateIntAttribute(name string, value int32) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_INT)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateFloatAttribute(name string, value float64) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_FLOAT)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateBoolAttribute(name string, value bool) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_BOOL)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateStringAttribute(name string, value string) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_STRING)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateTimeAttribute(name string, value float64) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_TIME)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateColorAttribute(name string, value [4]byte) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_COLOR)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateVector2Attribute(name string, value [2]float64) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_VECTOR2)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateVector3Attribute(name string, value [3]float64) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_VECTOR3)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateVector4Attribute(name string, value [4]float64) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_VECTOR4)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateQAngleAttribute(name string, value [3]float64) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_QANGLE)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateQuaternionAttribute(name string, value [4]float64) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_QUATERNION)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateMatrixAttribute(name string, value [16]float64) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_VMATRIX)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}
