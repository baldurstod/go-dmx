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

	if dmx.AT_VMATRIX != 14 {
		t.Error("dmx.AT_VMATRIX != 14")
	}

	if dmx.AT_FIRST_ARRAY_TYPE != 15 {
		t.Error("dmx.AT_FIRST_ARRAY_TYPE != 15")
	}

	if dmx.AT_TYPE_COUNT != 29 {
		t.Error("dmx.AT_TYPE_COUNT != 29")
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
	root.CreateColorAttribute("color_attrib", [...]float64{1, 2, 3, 4})
	root.CreateVector2Attribute("vec2_attrib", [...]float64{1.414, 3.14})
	root.CreateVector3Attribute("vec3_attrib", [...]float64{1.23, 4.56, 7.89})
	root.CreateVector4Attribute("vec4_attrib", [...]float64{-1.414, -3.14, -2.718, 10000.000123})
	root.CreateQAngleAttribute("qangle_attrib", [...]float64{0, 90, 270})
	root.CreateQuaternionAttribute("quaternion_attrib", [...]float64{0, 0.7071, 0, 0.7071})
	root.CreateMatrixAttribute("matrix_attrib", [...]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
	elem := dmx.NewDmElement("DmElement")
	root.CreateElementAttribute("inline_element", dmx.NewDmElement("DmElement"))
	root.CreateElementAttribute("non_inline_element_1", elem)
	root.CreateElementAttribute("non_inline_element_2", elem)

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, root)

	log.Println(buf)
}
