package dmx

type DmAttributeType = uint8

const (
	AT_UNKNOWN = iota

	AT_ELEMENT
	AT_INT
	AT_FLOAT
	AT_BOOL
	AT_STRING
	AT_VOID
	AT_TIME
	AT_COLOR //rgba
	AT_VECTOR2
	AT_VECTOR3
	AT_VECTOR4
	AT_QANGLE
	AT_QUATERNION
	AT_VMATRIX
	AT_UINT64
	// Reserved
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	AT_ELEMENT_ARRAY
	AT_INT_ARRAY
	AT_FLOAT_ARRAY
	AT_BOOL_ARRAY
	AT_STRING_ARRAY
	AT_VOID_ARRAY
	AT_TIME_ARRAY
	AT_COLOR_ARRAY
	AT_VECTOR2_ARRAY
	AT_VECTOR3_ARRAY
	AT_VECTOR4_ARRAY
	AT_QANGLE_ARRAY
	AT_QUATERNION_ARRAY
	AT_VMATRIX_ARRAY
	AT_UINT64_ARRAY
	AT_TYPE_COUNT

	AT_FIRST_VALUE_TYPE = AT_ELEMENT
	AT_FIRST_ARRAY_TYPE = AT_ELEMENT_ARRAY
)

var type_to_string = [...]string{
	"",
	"element",
	"int",
	"float",
	"bool",
	"string",
	"binary",
	"time",
	"color",
	"vector2",
	"vector3",
	"vector4",
	"qangle",
	"quaternion",
	"matrix",
	"uint64",
	// Reserved
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"element_array",
	"int_array",
	"float_array",
	"bool_array",
	"string_array",
	"binary_array",
	"time_array",
	"color_array",
	"vector2_array",
	"vector3_array",
	"vector4_array",
	"qangle_array",
	"quaternion_array",
	"matrix_array",
	"uint64_array",
}
