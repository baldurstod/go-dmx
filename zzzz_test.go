package dmx_test

import (
	"bytes"
	"github.com/baldurstod/go-dmx"
	"log"
	"testing"
)

func TestAttributesTypes(t *testing.T) {
	log.Println(dmx.AT_FIRST_VALUE_TYPE, dmx.AT_VMATRIX, dmx.AT_FIRST_ARRAY_TYPE, dmx.AT_TYPE_COUNT)

	if dmx.AT_FIRST_VALUE_TYPE != 1 {
		t.Error("dmx.AT_FIRST_VALUE_TYPE != 1")
	}

	if dmx.AT_UINT64 != 15 {
		t.Error("dmx.AT_UINT64 != 15")
	}

	if dmx.AT_FIRST_ARRAY_TYPE != 16 {
		t.Error("dmx.AT_FIRST_ARRAY_TYPE != 16")
	}

	if dmx.AT_TYPE_COUNT != 31 {
		t.Error("dmx.AT_TYPE_COUNT != 31")
	}
}

func TestAttributes(t *testing.T) {
	element := dmx.NewDmElement("DmElement")

	attribute := element.CreateAttribute("test", dmx.AT_INT)

	log.Println(attribute)
}

func TestTokens(t *testing.T) {
	log.Println(dmx.TOKEN_INVALID, dmx.TOKEN_OPEN_BRACE, dmx.TOKEN_CLOSE_BRACE, dmx.TOKEN_OPEN_BRACKET, dmx.TOKEN_EOF)

	if dmx.TOKEN_INVALID != -1 {
		t.Error("dmx.TOKEN_INVALID != -1")
	}

	if dmx.TOKEN_OPEN_BRACE != 0 {
		t.Error("dmx.TOKEN_OPEN_BRACE != 0")
	}

	if dmx.TOKEN_CLOSE_BRACE != 1 {
		t.Error("dmx.TOKEN_CLOSE_BRACE != 1")
	}

	if dmx.TOKEN_EOF != 7 {
		t.Error("dmx.TOKEN_EOF != 7")
	}
}

func TestSerializeText(t *testing.T) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	root := dmx.NewDmElement("DmElement")
	root.CreateIntAttribute("int_attrib", 1234)
	root.CreateFloatAttribute("float_attrib", 123.456)
	root.CreateBoolAttribute("bool_attrib_false", false)
	root.CreateBoolAttribute("bool_attrib_true", true)
	root.CreateStringAttribute("string_attrib", "this is a string")
	root.CreateTimeAttribute("time_attrib", 123)
	root.CreateColorAttribute("color_attrib", [...]byte{1, 2, 3, 4})
	root.CreateVector2Attribute("vec2_attrib", [...]float64{1.414, 3.14})
	root.CreateVector3Attribute("vec3_attrib", [...]float64{1.23, 4.56, 7.89})
	root.CreateVector4Attribute("vec4_attrib", [...]float64{-1.414, -3.14, -2.718, 10000.000123})
	root.CreateQAngleAttribute("qangle_attrib", [...]float64{0, 90, 270})
	root.CreateQuaternionAttribute("quaternion_attrib", [...]float64{0, 0.7071, 0, 0.7071})
	root.CreateMatrixAttribute("matrix_attrib", [...]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
	root.CreateUint64Attribute("uint64_attrib", 18446744073709551)
	root.CreateElementAttribute("inline_element", dmx.NewDmElement("DmElement"))
	elem := dmx.NewDmElement("DmElement")
	root.CreateElementAttribute("non_inline_element_1", elem)
	root.CreateElementAttribute("non_inline_element_2", elem)
	elem2 := dmx.NewDmElement("DmElement")
	root.CreateElementAttribute("non_inline_element_3", elem2)
	root.CreateElementAttribute("non_inline_element_4", elem2)

	elemArray := root.CreateAttribute("element_array_attrib", dmx.AT_ELEMENT_ARRAY)
	elemArray.PushElement(dmx.NewDmElement("DmElement"))
	elemArray.PushElement(dmx.NewDmElement("DmElement"))
	elemArray.PushElement(dmx.NewDmElement("DmElement"))
	elemArray.PushElement(elem)

	intArray := root.CreateAttribute("int_array_attrib", dmx.AT_INT_ARRAY)
	intArray.PushInt(1)
	intArray.PushInt(2)
	intArray.PushInt(3)

	floatArray := root.CreateAttribute("float_array_attrib", dmx.AT_FLOAT_ARRAY)
	floatArray.PushFloat(1.414)
	floatArray.PushFloat(2.718)
	floatArray.PushFloat(3.14)

	boolArray := root.CreateAttribute("bool_array_attrib", dmx.AT_BOOL_ARRAY)
	boolArray.PushBool(true)
	boolArray.PushBool(true)
	boolArray.PushBool(false)

	stringArray := root.CreateAttribute("string_array_attrib", dmx.AT_STRING_ARRAY)
	stringArray.PushString("this is string 1")
	stringArray.PushString("this is string 2")
	stringArray.PushString("this is string 3")

	timeArray := root.CreateAttribute("time_array_attrib", dmx.AT_TIME_ARRAY)
	timeArray.PushTime(1)
	timeArray.PushTime(2)
	timeArray.PushTime(3)

	colorArray := root.CreateAttribute("color_array_attrib", dmx.AT_COLOR_ARRAY)
	colorArray.PushColor([...]byte{1, 2, 3, 4})
	colorArray.PushColor([...]byte{5, 6, 7, 8})
	colorArray.PushColor([...]byte{9, 10, 11, 12})

	vec2Array := root.CreateAttribute("vec2_array_attrib", dmx.AT_VECTOR2_ARRAY)
	vec2Array.PushVector2([...]float64{1.414, 3.14})
	vec2Array.PushVector2([...]float64{1.414, 3.14})
	vec2Array.PushVector2([...]float64{1.414, 3.14})

	vec3Array := root.CreateAttribute("vec3_array_attrib", dmx.AT_VECTOR3_ARRAY)
	vec3Array.PushVector3([...]float64{1.23, 4.56, 7.89})
	vec3Array.PushVector3([...]float64{1.23, 4.56, 7.89})
	vec3Array.PushVector3([...]float64{1.23, 4.56, 7.89})

	vec4Array := root.CreateAttribute("vec4_array_attrib", dmx.AT_VECTOR4_ARRAY)
	vec4Array.PushVector4([...]float64{-1.414, -3.14, -2.718, 10000.000123})
	vec4Array.PushVector4([...]float64{-1.414, -3.14, -2.718, 10000.000123})
	vec4Array.PushVector4([...]float64{-1.414, -3.14, -2.718, 10000.000123})

	uint64Array := root.CreateAttribute("uint64_array_attrib", dmx.AT_UINT64_ARRAY)
	uint64Array.PushUint64(18446744073709551)
	uint64Array.PushUint64(18446744073709551)
	uint64Array.PushUint64(18446744073709551)

	root.CreateElementAttribute("nil_element", nil)

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, root)

	log.Println(buf)
}

func TestSerializeInlineText(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	root := dmx.NewDmElement("DmElement")
	clip := dmx.NewDmElement("DmeFilmClip")
	timeFrame := dmx.NewDmElement("DmeTimeFrame")

	root.CreateElementAttribute("activeClip", clip)
	elemArray := root.CreateAttribute("clipBin", dmx.AT_ELEMENT_ARRAY)
	elemArray.PushElement(clip)

	root.CreateElementAttribute("activeClip", clip)
	clip.CreateElementAttribute("timeFrame", timeFrame)

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, root)

	log.Println(buf)
}
