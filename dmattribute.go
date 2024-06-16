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
}

func (attribute *DmAttribute) GetValue() interface{} {
	return attribute.value
}

func (attribute *DmAttribute) SetValue(value interface{}) {
	//TODO: check if value is compatible with type
	attribute.value = value
}

func (attribute *DmAttribute) GetOwner() *DmElement {
	return attribute.owner
}
