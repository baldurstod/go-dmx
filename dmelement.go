package dmx

import "github.com/baldurstod/go-vector"

type DmElement struct {
	Name        string
	id          DmObjectId
	elementType string
	attributes  map[string]*DmAttribute
}

func NewDmElement(name string, elementType string) *DmElement {
	return &DmElement{
		Name:        name,
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

func (element *DmElement) CreateFloatAttribute(name string, value float32) *DmAttribute {
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

func (element *DmElement) CreateTimeAttribute(name string, value float32) *DmAttribute {
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

func (element *DmElement) CreateVector2Attribute(name string, value vector.Vector2[float32]) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_VECTOR2)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateVector3Attribute(name string, value vector.Vector3[float32]) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_VECTOR3)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateVector4Attribute(name string, value vector.Vector4[float32]) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_VECTOR4)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateQAngleAttribute(name string, value vector.Vector3[float32]) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_QANGLE)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateQuaternionAttribute(name string, value vector.Quaternion[float32]) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_QUATERNION)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateMatrixAttribute(name string, value [16]float32) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_VMATRIX)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}

func (element *DmElement) CreateUint64Attribute(name string, value uint64) *DmAttribute {
	attribute := element.CreateAttribute(name, AT_UINT64)

	if attribute != nil {
		attribute.SetValue(value)
	}

	return attribute
}
