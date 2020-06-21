package utils

import "google.golang.org/protobuf/types/descriptorpb"

// json 字段标签
const (
	LABEL_OPTIONAL = "可选" // 可选的
	LABEL_REQUIRED = "必须" // 必须的
	LABEL_REPEATED = "重复" // 重复的
)

// json 字段类型
const (
	TPYE_NUMBER  = `Number`
	TYPE_STRING  = `String`
	TYPE_BOOLEAN = `Boolean`
	TYPE_OBJECT  = `Object`
)

var (
	jsonTypeDefaultValue = map[string]interface{}{
		TPYE_NUMBER:  0,
		TYPE_STRING:  "",
		TYPE_BOOLEAN: false,
		TYPE_OBJECT:  nil,
	}

	protoType2JsonType = map[descriptorpb.FieldDescriptorProto_Type]string{
		descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:   TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_FLOAT:    TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_INT32:    TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_UINT32:   TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_INT64:    TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_UINT64:   TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32:  TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64:  TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_SINT32:   TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_SINT64:   TPYE_NUMBER,
		descriptorpb.FieldDescriptorProto_TYPE_BYTES:    TYPE_STRING,
		descriptorpb.FieldDescriptorProto_TYPE_STRING:   TYPE_STRING,
		descriptorpb.FieldDescriptorProto_TYPE_BOOL:     TYPE_BOOLEAN,
		descriptorpb.FieldDescriptorProto_TYPE_ENUM:     TYPE_OBJECT,
		descriptorpb.FieldDescriptorProto_TYPE_GROUP:    TYPE_OBJECT,
		descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:  TYPE_OBJECT,
	}

	protoLabel2JsonLabel = map[descriptorpb.FieldDescriptorProto_Label]string{
		descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL: LABEL_OPTIONAL,
		descriptorpb.FieldDescriptorProto_LABEL_REQUIRED: LABEL_REQUIRED,
		descriptorpb.FieldDescriptorProto_LABEL_REPEATED: LABEL_REPEATED,
	}
)

type (
	Package struct {
		PackageName string
		Services    map[string]*Service // service infor
		Messages    map[string]*Message // message infor
		Enums       map[string]*Enum    // enum infor
	}

	Service struct {
		ServiceName string    // service name
		Methods     []*Method // method infor
		Description string    // description
		Uri         string    // uri
	}

	Method struct {
		MethodName    string // method name
		RequestParam  string // method request
		ResponseParam string // method response
		Description   string // description
		Uri           string
	}

	Message struct {
		MessageName string   // message name
		Fields      []*Field // message fields
		Description string   // description
	}

	Enum struct {
		EnumName    string       // enum name
		Fields      []*EnumField // enum fields
		Description string       // description
	}

	Field struct {
		MessageName string // 字段所在 message

		ProtoName         string                                  // proto 名称
		ProtoLaber        descriptorpb.FieldDescriptorProto_Label // proto 标签
		ProtoType         descriptorpb.FieldDescriptorProto_Type  // 隐式类型
		ProtoTypeName     string                                  // 显示类型
		ProtoPackagePath  string                                  // 包路径
		ProtoFullTypeName string                                  // 类型全称
		ProtoNumber       int32                                   // 排序

		JsonName         string      // json 名称
		JsonLabel        string      // json 标签
		JsonType         string      // json 类型
		JsonDefaultValue interface{} // json 数据默认值

		Description string // description
	}

	EnumField struct {
		EnumFieldName  string // enum field name
		EnumFieldValue int32  // enum field value
		Description    string // description
	}

	desc struct {
		desc string `json:"desc"`
		path string `json:"path"`
	}
)
