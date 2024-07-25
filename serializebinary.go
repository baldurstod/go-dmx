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
		binary.Write(context.buf, binary.LittleEndian, []byte(s))
		binary.Write(context.buf, binary.LittleEndian, terminator)
	}

	return nil
}

func serializeDictBinary(context *serializerContext) error {
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
			if err := serializeAttribute[float32](context, a); err != nil {
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
			if err := serializeArrayAttribute[float32](context, a); err != nil {
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
			/*
				case AT_COLOR_ARRAY:
					attribute.value = make([][4]byte, 0)
				case AT_VECTOR2_ARRAY:
					attribute.value = make([]vector.Vector2[float32], 0)
				case AT_VECTOR3_ARRAY, AT_QANGLE_ARRAY:
					attribute.value = make([]vector.Vector3[float32], 0)
				case AT_VECTOR4_ARRAY:
					attribute.value = make([]vector.Vector4[float32], 0)
				case AT_QUATERNION_ARRAY:
					attribute.value = make([]vector.Quaternion[float32], 0)
				case AT_VMATRIX_ARRAY:
					attribute.value = make([][16]float32, 0)
				case AT_UINT64_ARRAY:
					attribute.value = make([]uint64, 0)*/
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
		if err := binary.Write(context.buf, binary.LittleEndian, []byte(v)); err != nil {
			return err
		}
		if err := binary.Write(context.buf, binary.LittleEndian, byte(0)); err != nil {
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

func serializeArrayBinary(context *serializerContext, attribute *DmAttribute) error {
	buf := context.buf
	switch attribute.attributeType {
	case AT_ELEMENT_ARRAY:
		a := attribute.value.([]*DmElement)
		l := len(a)
		for k, element := range a {
			if shouldInlineElement(context, element) {
				writeTabs(context)
				err := serializeElementText(context, element)
				if err != nil {
					return err
				}
			} else {
				writeTabs(context)
				buf.WriteString("\"element\" ")
				uuid := fmt.Sprintf("\"%x-%x-%x-%x-%x\"", element.id[0:4], element.id[4:6], element.id[6:8], element.id[8:10], element.id[10:])
				buf.WriteString(uuid)
				//buf.WriteString("\"")
				//newLine(context)
			}
			if k < l-1 {
				buf.WriteString(",")
			}
			newLine(context)
		}

	case AT_INT_ARRAY:
		a := attribute.value.([]int32)
		l := len(a)
		for k, i := range a {
			writeTabs(context)
			buf.WriteString(strconv.FormatInt(int64(i), 10))
			if k < l-1 {
				buf.WriteString(",")
			}
			newLine(context)
		}
	case AT_FLOAT_ARRAY, AT_TIME_ARRAY:
		a := attribute.value.([]float32)
		l := len(a)
		for k, f := range a {
			writeTabs(context)
			buf.WriteString("\"" + strconv.FormatFloat(float64(f), 'g', -1, 32) + "\"")
			if k < l-1 {
				buf.WriteString(",")
			}
			newLine(context)
		}
	case AT_BOOL_ARRAY:
		a := attribute.value.([]bool)
		l := len(a)
		for k, b := range a {
			writeTabs(context)

			if b {
				buf.WriteString("\"1\"")
			} else {
				buf.WriteString("\"0\"")
			}

			if k < l-1 {
				buf.WriteString(",")
			}
			newLine(context)
		}
	case AT_STRING_ARRAY:
		a := attribute.value.([]string)
		l := len(a)
		for k, s := range a {
			writeTabs(context)
			buf.WriteString("\"")
			buf.WriteString(s)
			buf.WriteString("\"")
			if k < l-1 {
				buf.WriteString(",")
			}
			newLine(context)
		}
	case AT_COLOR_ARRAY:
		a := attribute.value.([][4]byte)
		l := len(a)
		for k, v := range a {
			writeTabs(context)
			buf.WriteString("\"")
			buf.WriteString(fmt.Sprintf("%d %d %d %d", v[0], v[1], v[2], v[3]))
			buf.WriteString("\"")
			if k < l-1 {
				buf.WriteString(",")
			}
			newLine(context)
		}
	case AT_VECTOR2_ARRAY:
		a := attribute.value.([]vector.Vector2[float32])
		l := len(a)
		for k, v := range a {
			writeTabs(context)
			buf.WriteString("\"")
			buf.WriteString(fmt.Sprintf("%g %g", v[0], v[1]))
			buf.WriteString("\"")
			if k < l-1 {
				buf.WriteString(",")
			}
			newLine(context)
		}
	case AT_VECTOR3_ARRAY, AT_QANGLE_ARRAY:
		a := attribute.value.([]vector.Vector3[float32])
		l := len(a)
		for k, v := range a {
			writeTabs(context)
			buf.WriteString("\"")
			buf.WriteString(fmt.Sprintf("%g %g %g", v[0], v[1], v[2]))
			buf.WriteString("\"")
			if k < l-1 {
				buf.WriteString(",")
			}
			newLine(context)
		}
	case AT_VECTOR4_ARRAY:
		a := attribute.value.([]vector.Vector4[float32])
		l := len(a)
		for k, v := range a {
			writeTabs(context)
			buf.WriteString("\"")
			buf.WriteString(fmt.Sprintf("%g %g %g %g", v[0], v[1], v[2], v[3]))
			buf.WriteString("\"")
			if k < l-1 {
				buf.WriteString(",")
			}
			newLine(context)
		}
	case AT_QUATERNION_ARRAY:
		a := attribute.value.([]vector.Quaternion[float32])
		l := len(a)
		for k, v := range a {
			writeTabs(context)
			buf.WriteString("\"")
			buf.WriteString(fmt.Sprintf("%g %g %g %g", v[0], v[1], v[2], v[3]))
			buf.WriteString("\"")
			if k < l-1 {
				buf.WriteString(",")
			}
			newLine(context)
		}
	case AT_UINT64_ARRAY:
		a := attribute.value.([]uint64)
		l := len(a)
		for k, i := range a {
			writeTabs(context)
			buf.WriteString(strconv.FormatUint(i, 10))
			if k < l-1 {
				buf.WriteString(",")
			}
			newLine(context)
		}
	default:
		panic("Unknown attribute type in serializeArrayBinary " + type_to_string[attribute.attributeType])
	}
	/*
		log.Println(attribute.value)
		a, ok := attribute.value.([]*DmElement)
		if !ok {
			panic("Value is not an array")
		}

		for _, v := range a {
			log.Println(v)
		}
	*/
	return nil
}

func serializeAttributeBinary(context *serializerContext, attribute *DmAttribute) error {
	buf := context.buf
	attributeType := attribute.attributeType
	if attributeType >= AT_FIRST_ARRAY_TYPE {
		//panic("TODO")

		writeTabs(context)
		buf.WriteString("\"")
		buf.WriteString(attribute.name)
		buf.WriteString("\" \"")
		buf.WriteString(type_to_string[attribute.attributeType])
		buf.WriteString("\"")
		newLine(context)
		writeTabs(context)
		buf.WriteString("[")
		newLine(context)
		pushTab(context)

		err := serializeArrayBinary(context, attribute)
		if err != nil {
			return err
		}

		popTab(context)
		writeTabs(context)
		buf.WriteString("]")
		newLine(context)
	} else {
		if attributeType == AT_ELEMENT {
			element := attribute.value.(*DmElement)

			if shouldInlineElement(context, element) {
				writeTabs(context)
				buf.WriteString("\"")
				buf.WriteString(attribute.name)
				buf.WriteString("\" ")
				err := serializeElementText(context, attribute.value.(*DmElement))
				if err != nil {
					return err
				}
				newLine(context)
			} else {
				writeTabs(context)
				buf.WriteString("\"")
				buf.WriteString(attribute.name)
				buf.WriteString("\" \"element\" ")
				if element != nil {
					uuid := fmt.Sprintf("\"%x-%x-%x-%x-%x\"", element.id[0:4], element.id[4:6], element.id[6:8], element.id[8:10], element.id[10:])
					buf.WriteString(uuid)
				} else {
					buf.WriteString("\"\"")
				}
				//buf.WriteString("\"")
				newLine(context)
			}
		} else {
			writeTabs(context)
			buf.WriteString("\"")
			buf.WriteString(attribute.name)
			buf.WriteString("\" \"")
			buf.WriteString(type_to_string[attribute.attributeType])
			buf.WriteString("\" \"")
			buf.WriteString(attribute.StringValue())
			buf.WriteString("\"")
			newLine(context)
		}
	}

	return nil
}
