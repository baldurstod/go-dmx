package dmx_test

import (
	"bytes"
	"log"
	"os"
	"path"
	"testing"

	"github.com/baldurstod/go-dmx"
)

func TestSerializeBinary(t *testing.T) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	root := dmx.NewDmElement("test_DmElement", "DmElement")
	root.CreateIntAttribute("int_attrib", 1234)
	root.CreateFloatAttribute("float_attrib", 123.456)
	root.CreateBoolAttribute("bool_attrib_false", false)
	root.CreateBoolAttribute("bool_attrib_true", true)
	root.CreateStringAttribute("string_attrib", "this is a string")
	root.CreateTimeAttribute("time_attrib", 123)
	root.CreateColorAttribute("color_attrib", [...]byte{1, 2, 3, 4})
	root.CreateVector2Attribute("vec2_attrib", [...]float32{1.414, 3.14})
	root.CreateVector3Attribute("vec3_attrib", [...]float32{1.23, 4.56, 7.89})
	root.CreateVector4Attribute("vec4_attrib", [...]float32{-1.414, -3.14, -2.718, 10000.000123})
	root.CreateQAngleAttribute("qangle_attrib", [...]float32{0, 90, 270})
	root.CreateQuaternionAttribute("quaternion_attrib", [...]float32{0, 0.7071, 0, 0.7071})
	root.CreateMatrixAttribute("matrix_attrib", [...]float32{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
	root.CreateUint64Attribute("uint64_attrib", 18446744073709551)
	root.CreateElementAttribute("inline_element", dmx.NewDmElement("test_DmElement", "DmElement"))
	elem := dmx.NewDmElement("test_DmElement", "DmElement")
	root.CreateElementAttribute("non_inline_element_1", elem)
	root.CreateElementAttribute("non_inline_element_2", elem)
	elem2 := dmx.NewDmElement("test_DmElement", "DmElement")
	root.CreateElementAttribute("non_inline_element_3", elem2)
	root.CreateElementAttribute("non_inline_element_4", elem2)

	elemArray := root.CreateAttribute("element_array_attrib", dmx.AT_ELEMENT_ARRAY)
	elemArray.PushElement(dmx.NewDmElement("test_DmElement", "DmElement"))
	elemArray.PushElement(dmx.NewDmElement("test_DmElement", "DmElement"))
	elemArray.PushElement(dmx.NewDmElement("test_DmElement", "DmElement"))
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
	vec2Array.PushVector2([...]float32{1.414, 3.14})
	vec2Array.PushVector2([...]float32{1.414, 3.14})
	vec2Array.PushVector2([...]float32{1.414, 3.14})

	vec3Array := root.CreateAttribute("vec3_array_attrib", dmx.AT_VECTOR3_ARRAY)
	vec3Array.PushVector3([...]float32{1.23, 4.56, 7.89})
	vec3Array.PushVector3([...]float32{1.23, 4.56, 7.89})
	vec3Array.PushVector3([...]float32{1.23, 4.56, 7.89})

	vec4Array := root.CreateAttribute("vec4_array_attrib", dmx.AT_VECTOR4_ARRAY)
	vec4Array.PushVector4([...]float32{-1.414, -3.14, -2.718, 10000.000123})
	vec4Array.PushVector4([...]float32{-1.414, -3.14, -2.718, 10000.000123})
	vec4Array.PushVector4([...]float32{-1.414, -3.14, -2.718, 10000.000123})

	quatArray := root.CreateAttribute("quat_array_attrib", dmx.AT_QUATERNION_ARRAY)
	quatArray.PushQuaternion([...]float32{-1.414, -3.14, -2.718, 10000.000123})
	quatArray.PushQuaternion([...]float32{-1.414, -3.14, -2.718, 10000.000123})
	quatArray.PushQuaternion([...]float32{-1.414, -3.14, -2.718, 10000.000123})

	uint64Array := root.CreateAttribute("uint64_array_attrib", dmx.AT_UINT64_ARRAY)
	uint64Array.PushUint64(18446744073709551)
	uint64Array.PushUint64(18446744073709551)
	uint64Array.PushUint64(18446744073709551)

	root.CreateElementAttribute("nil_element", nil)

	buf := new(bytes.Buffer)
	if err := dmx.SerializeBinary(buf, root, "sfm_session", 22); err != nil {
		t.Error(err)
	}

	log.Println(buf)
	os.WriteFile(path.Join("./var/", "test_session.dmx"), buf.Bytes(), 0666)
}

func TestSerializeBinary2(t *testing.T) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	root := dmx.NewDmElement("test_DmElement", "DmElement")

	buf := new(bytes.Buffer)
	if err := dmx.SerializeBinary(buf, root, "sfm_session", 22); err != nil {
		t.Error(err)
	}

	log.Println(buf)
	os.WriteFile(path.Join("./var/", "test_session.dmx"), buf.Bytes(), 0666)
}
