package dmx

type DmElement struct {
	id         DmObjectId
	attributes map[string]*DmAttribute
}

func NewDmElement() *DmElement {
	return &DmElement{
		attributes: map[string]*DmAttribute{},
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

	return attribute
}

func (element *DmElement) GetId() DmObjectId {
	return element.id
}

func (element *DmElement) SetId(id DmObjectId) {
	element.id = id
}
