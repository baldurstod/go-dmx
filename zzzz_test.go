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
	root.CreateIntAttribute("test", 1234)
	elem := dmx.NewDmElement("DmElement")
	root.CreateElementAttribute("Sub", dmx.NewDmElement("DmElement"))
	root.CreateElementAttribute("Dup1", elem)
	root.CreateElementAttribute("Dup2", elem)

	buf := new(bytes.Buffer)
	dmx.SerializeText(buf, root)

	log.Println(buf)
}
