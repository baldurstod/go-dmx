package dmx

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	"github.com/baldurstod/go-vector"
)

func SerializeBinary(buf *bytes.Buffer, root *DmElement, format string, formatVersion int) error {
	context := newSerializerContext(buf)

	if _, err := buf.WriteString(fmt.Sprintf("<!-- dmx encoding binary 9 format %s %d -->\n\x00", format, formatVersion)); err != nil {
		return err
	}
	if err := binary.Write(context.buf, binary.LittleEndian, uint32(0)); err != nil {
		return err
	}

	buildElementList(context, root)

	serializeStringsBinary(context)
	/*
		err := serializeElementBinary(context, root)
		if err != nil {
			return err
		}
	*/

	return serializeDictBinary(context)
}

func serializeStringsBinary(context *serializerContext) error {
	var count uint32 = uint32(len(context.stringDictionary))
	binary.Write(context.buf, binary.LittleEndian, count)

	var terminator byte

	for _, s := range context.stringDictionary2 {
		if err := binary.Write(context.buf, binary.LittleEndian, []byte(s)); err != nil {
			return err
		}
		if err := binary.Write(context.buf, binary.LittleEndian, terminator); err != nil {
			return err
		}
	}

	return nil
}

func serializeDictBinary(context *serializerContext) error {
	if err := binary.Write(context.buf, binary.LittleEndian, uint32(len(context.dictionary2))); err != nil {
		return err
	}
	for _, e := range context.dictionary2 {
		err := serializeElementBinary(context, e)
		if err != nil {
			return err
		}
	}

	for _, e := range context.dictionary2 {
		err := serializeAttributesBinary(context, e)
		if err != nil {
			return err
		}
	}
	return nil
}

func serializeElementBinary(context *serializerContext, element *DmElement) error {
	if element == nil {
		return nil
	}

	typeId := context.stringDictionary[element.elementType]
	nameId := context.stringDictionary[element.Name]

	if err := binary.Write(context.buf, binary.LittleEndian, typeId); err != nil {
		return err
	}
	if err := binary.Write(context.buf, binary.LittleEndian, nameId); err != nil {
		return err
	}
	if err := binary.Write(context.buf, binary.LittleEndian, element.id); err != nil {
		return err
	}

	return nil
}

func serializeAttributesBinary(context *serializerContext, element *DmElement) error {
	if err := binary.Write(context.buf, binary.LittleEndian, uint32(len(element.attributes))); err != nil {
		return err
	}

	for _, a := range element.attributes {
		nameId := context.stringDictionary[a.name]
		if err := binary.Write(context.buf, binary.LittleEndian, nameId); err != nil {
			return err
		}
		if err := binary.Write(context.buf, binary.LittleEndian, a.attributeType); err != nil {
			return err
		}

		switch a.attributeType {
		case AT_ELEMENT:
			if v, ok := a.GetValue().(*DmElement); ok {
				if err := serializeElementAttribute(context, v); err != nil {
					return err
				}
			} else {
				return errors.New("attribute is of type element but doesn't contain an element")
			}
		case AT_INT:
			if err := serializeAttribute[int32](context, a); err != nil {
				return err
			}
		case AT_FLOAT:
			if err := serializeAttribute[float32](context, a); err != nil {
				return err
			}
		case AT_BOOL:
			if err := serializeAttribute[bool](context, a); err != nil {
				return err
			}
		case AT_STRING:
			if err := serializeStringAttribute(context, a); err != nil {
				return err
			}
		case AT_TIME:
			if err := serializeTimeAttribute(context, a); err != nil {
				return err
			}
		case AT_COLOR:
			if err := serializeAttribute[[4]byte](context, a); err != nil {
				return err
			}
		case AT_VECTOR2:
			if err := serializeAttribute[vector.Vector2[float32]](context, a); err != nil {
				return err
			}
		case AT_VECTOR3:
			if err := serializeAttribute[vector.Vector3[float32]](context, a); err != nil {
				return err
			}
		case AT_VECTOR4:
			if err := serializeAttribute[vector.Vector4[float32]](context, a); err != nil {
				return err
			}
		case AT_QANGLE:
			if err := serializeAttribute[vector.Vector3[float32]](context, a); err != nil {
				return err
			}
		case AT_QUATERNION:
			if err := serializeAttribute[vector.Quaternion[float32]](context, a); err != nil {
				return err
			}
		case AT_VMATRIX:
			if err := serializeAttribute[[16]float32](context, a); err != nil {
				return err
			}
		case AT_UINT64:
			if err := serializeAttribute[uint64](context, a); err != nil {
				return err
			}
		case AT_ELEMENT_ARRAY:
			if v, ok := a.GetValue().([]*DmElement); ok {
				if err := binary.Write(context.buf, binary.LittleEndian, uint32(len(v))); err != nil {
					return err
				}
				for _, e := range v {
					if err := serializeElementAttribute(context, e); err != nil {
						return err
					}
				}
			} else {
				return errors.New("attribute is of type element but doesn't contain an element")
			}
		case AT_INT_ARRAY:
			if err := serializeArrayAttribute[int32](context, a); err != nil {
				return err
			}
		case AT_FLOAT_ARRAY:
			if err := serializeArrayAttribute[float32](context, a); err != nil {
				return err
			}
		case AT_BOOL_ARRAY:
			if err := serializeArrayAttribute[bool](context, a); err != nil {
				return err
			}
		case AT_STRING_ARRAY:
			if err := serializeStringArrayAttribute(context, a); err != nil {
				return err
			}
		case AT_TIME_ARRAY:
			if err := serializeTimeArrayAttribute(context, a); err != nil {
				return err
			}
		case AT_COLOR_ARRAY:
			if err := serializeArrayAttribute[[4]byte](context, a); err != nil {
				return err
			}
		case AT_VECTOR2_ARRAY:
			if err := serializeArrayAttribute[vector.Vector2[float32]](context, a); err != nil {
				return err
			}
		case AT_VECTOR3_ARRAY:
			if err := serializeArrayAttribute[vector.Vector3[float32]](context, a); err != nil {
				return err
			}
		case AT_VECTOR4_ARRAY:
			if err := serializeArrayAttribute[vector.Vector4[float32]](context, a); err != nil {
				return err
			}
		case AT_QUATERNION_ARRAY:
			if err := serializeArrayAttribute[vector.Quaternion[float32]](context, a); err != nil {
				return err
			}
		case AT_VMATRIX_ARRAY:
			if err := serializeArrayAttribute[[16]float32](context, a); err != nil {
				return err
			}
		case AT_UINT64_ARRAY:
			if err := serializeArrayAttribute[uint64](context, a); err != nil {
				return err
			}
		default:
			return errors.New("unknown attribute type " + strconv.Itoa(int(a.attributeType)))
		}
	}
	return nil
}

func serializeElementAttribute(context *serializerContext, v *DmElement) error {
	if v == nil {
		if err := binary.Write(context.buf, binary.LittleEndian, int32(-1)); err != nil {
			return err
		}
		return nil
	}

	if e, ok := context.dictionary[v]; ok {
		if err := binary.Write(context.buf, binary.LittleEndian, e.id); err != nil {
			return err
		}
	} else {
		return errors.New("missing dictionnary entry for element")
	}
	return nil
}

func serializeAttribute[T int32 | float32 | bool | [4]byte | vector.Vector2[float32] | vector.Vector3[float32] | vector.Vector4[float32] | vector.Quaternion[float32] | [16]float32 | uint64](context *serializerContext, attribute *DmAttribute) error {
	if v, ok := attribute.value.(T); ok {
		if err := binary.Write(context.buf, binary.LittleEndian, v); err != nil {
			return err
		}
	} else {
		return errors.New("unable to cast attribute value")
	}

	return nil
}

func serializeStringAttribute(context *serializerContext, attribute *DmAttribute) error {
	if v, ok := attribute.value.(string); ok {
		stringId := context.stringDictionary[v]

		if err := binary.Write(context.buf, binary.LittleEndian, stringId); err != nil {
			return err
		}

		/*
			if err := binary.Write(context.buf, binary.LittleEndian, []byte(v)); err != nil {
				return err
			}
			if err := binary.Write(context.buf, binary.LittleEndian, byte(0)); err != nil {
				return err
			}
		*/
	} else {
		return errors.New("unable to cast attribute value")
	}

	return nil
}

func serializeTimeAttribute(context *serializerContext, attribute *DmAttribute) error {
	if v, ok := attribute.value.(float32); ok {
		if err := binary.Write(context.buf, binary.LittleEndian, int32(v*10000)); err != nil {
			return err
		}
	} else {
		return errors.New("unable to cast attribute value")
	}

	return nil
}

func serializeArrayAttribute[T int32 | float32 | bool | [4]byte | vector.Vector2[float32] | vector.Vector3[float32] | vector.Vector4[float32] | vector.Quaternion[float32] | [16]float32 | uint64](context *serializerContext, attribute *DmAttribute) error {
	if v, ok := attribute.value.([]T); ok {
		if err := binary.Write(context.buf, binary.LittleEndian, uint32(len(v))); err != nil {
			return err
		}
		for _, e := range v {
			if err := binary.Write(context.buf, binary.LittleEndian, e); err != nil {
				return err
			}
		}
	} else {
		return errors.New("unable to cast attribute value")
	}

	return nil
}

func serializeStringArrayAttribute(context *serializerContext, attribute *DmAttribute) error {
	if v, ok := attribute.value.([]string); ok {
		if err := binary.Write(context.buf, binary.LittleEndian, uint32(len(v))); err != nil {
			return err
		}
		for _, e := range v {
			if err := binary.Write(context.buf, binary.LittleEndian, []byte(e)); err != nil {
				return err
			}
			if err := binary.Write(context.buf, binary.LittleEndian, byte(0)); err != nil {
				return err
			}
		}
	} else {
		return errors.New("unable to cast attribute value")
	}

	return nil
}

func serializeTimeArrayAttribute(context *serializerContext, attribute *DmAttribute) error {
	if v, ok := attribute.value.([]float32); ok {
		if err := binary.Write(context.buf, binary.LittleEndian, uint32(len(v))); err != nil {
			return err
		}
		for _, e := range v {
			if err := binary.Write(context.buf, binary.LittleEndian, int32(e*10000)); err != nil {
				return err
			}
		}
	} else {
		return errors.New("unable to cast attribute value")
	}

	return nil
}
