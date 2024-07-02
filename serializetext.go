package dmx

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/baldurstod/go-vector"
)

type DmToken int32

const (
	TOKEN_INVALID       = iota - 1 // A bogus token
	TOKEN_OPEN_BRACE               // {
	TOKEN_CLOSE_BRACE              // }
	TOKEN_OPEN_BRACKET             // [
	TOKEN_CLOSE_BRACKET            // ]
	TOKEN_COMMA                    // ,
	//		TOKEN_STRING,			// Any non-quoted string
	TOKEN_DELIMITED_STRING // Any quoted string
	TOKEN_INCLUDE          // #include
	TOKEN_EOF              // End of buffer
)

type serializerContext struct {
	buf        *bytes.Buffer
	dictionary map[*DmElement]uint
	tabs       int
}

func newSerializerContext(buf *bytes.Buffer) *serializerContext {
	return &serializerContext{
		buf:        buf,
		dictionary: make(map[*DmElement]uint),
		tabs:       0,
	}
}

func SerializeText(buf *bytes.Buffer, root *DmElement) error {
	context := newSerializerContext(buf)

	buf.WriteString("<!-- dmx encoding keyvalues2 4 format sfm_session 22 -->\n")

	buildElementList(context, root)

	err := serializeElementText(context, root)
	if err != nil {
		return err
	}

	newLine(context)
	return serializeDictText(context)
}

func buildElementList(context *serializerContext, element *DmElement) error {
	if element == nil {
		return nil
	}

	v, exist := context.dictionary[element]
	if exist {
		context.dictionary[element] = v + 1
		return nil // This element was previously processed
	} else {
		context.dictionary[element] = 1
	}

	for _, v := range element.attributes {
		switch v.attributeType {
		case AT_ELEMENT:
			e, ok := v.value.(*DmElement)
			if ok {
				buildElementList(context, e)
			}
		case AT_ELEMENT_ARRAY:
			a, ok := v.value.([]*DmElement)
			if ok {
				for _, e := range a {
					buildElementList(context, e)
				}
			}
		}
	}
	return nil
}

func shouldInlineElement(context *serializerContext, element *DmElement) bool {
	if element == nil {
		return false
	}
	v, exist := context.dictionary[element]
	if exist {
		return v < 2
	}
	return true
}

func serializeDictText(context *serializerContext) error {
	for e, i := range context.dictionary {
		if i > 1 {
			err := serializeElementText(context, e)
			if err != nil {
				return err
			}
			newLine(context)
		}
	}
	return nil
}

func serializeElementText(context *serializerContext, element *DmElement) error {
	if element == nil {
		return nil
	}
	buf := context.buf

	//writeTabs(context)
	buf.WriteString("\"")
	buf.WriteString(element.elementType)
	buf.WriteString("\"")
	newLine(context)
	writeTabs(context)
	buf.WriteString("{")
	newLine(context)
	pushTab(context)

	if element.Name != "" {
		writeTabs(context)
		buf.WriteString("\"name\" \"string\" \"")
		buf.WriteString(element.Name)
		buf.WriteString("\"")
		newLine(context)
	}

	writeTabs(context)
	buf.WriteString("\"id\" \"elementid\" ")
	uuid := fmt.Sprintf("\"%x-%x-%x-%x-%x\"", element.id[0:4], element.id[4:6], element.id[6:8], element.id[8:10], element.id[10:])
	buf.WriteString(uuid)
	newLine(context)

	err := serializeAttributesText(context, element)
	if err != nil {
		return err
	}

	popTab(context)
	writeTabs(context)
	buf.WriteString("}")
	//newLine(context)

	return nil
}

func serializeAttributesText(context *serializerContext, element *DmElement) error {
	for _, a := range element.attributes {
		err := serializeAttributeText(context, a)
		if err != nil {
			return err
		}
	}
	return nil
}

func serializeArrayText(context *serializerContext, attribute *DmAttribute) error {
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
				buf.WriteString("1")
			} else {
				buf.WriteString("0")
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
		panic("Unknown attribute type in serializeArrayText " + type_to_string[attribute.attributeType])
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

func serializeAttributeText(context *serializerContext, attribute *DmAttribute) error {
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

		err := serializeArrayText(context, attribute)
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

func newLine(context *serializerContext) {
	context.buf.WriteByte('\n')
}
func pushTab(context *serializerContext) {
	context.tabs++
}
func popTab(context *serializerContext) {
	if context.tabs > 0 {
		context.tabs--
	}
}
func writeTabs(context *serializerContext) {

	for i := 0; i < context.tabs; i++ {
		context.buf.WriteByte('\t')
	}
}
