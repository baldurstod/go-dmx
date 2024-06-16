package dmx_test

import (
	"github.com/baldurstod/go-dmx"
	"log"
	"testing"
)

func TestAttributesTypes(t *testing.T) {
	log.Println(dmx.AT_FIRST_VALUE_TYPE, dmx.AT_VMATRIX, dmx.AT_FIRST_ARRAY_TYPE, dmx.AT_TYPE_COUNT)

	if dmx.AT_FIRST_VALUE_TYPE != 1 {
		t.Error("dmx.AT_FIRST_VALUE_TYPE != 1")
	}

	if dmx.AT_VMATRIX != 15 {
		t.Error("dmx.AT_VMATRIX != 15")
	}

	if dmx.AT_FIRST_ARRAY_TYPE != 16 {
		t.Error("dmx.AT_FIRST_ARRAY_TYPE != 16")
	}

	if dmx.AT_TYPE_COUNT != 31 {
		t.Error("dmx.AT_TYPE_COUNT != 31")
	}
}

func TestAttributes(t *testing.T) {
	element := dmx.NewDmElement()

	attribute := element.CreateAttribute("test", dmx.AT_INT)

	log.Println(attribute)
}
