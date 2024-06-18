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
