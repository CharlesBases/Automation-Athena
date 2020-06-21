package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

var protofileModule = new(protofile)

var (
	req = new(pluginpb.CodeGeneratorRequest)
	rsp = new(pluginpb.CodeGeneratorResponse)
)

type protofile struct{}

// unmarshal proto.Unmarshal
func (*protofile) unmarshal() {
	data, err := ioutil.ReadAll(os.Stdin)
	ThrowCheck(err)

	ThrowCheck(proto.Unmarshal(data, req))

	if len(req.FileToGenerate) == 0 {
		ThrowCheck(fmt.Errorf("no file to generate"))
	}
}

// parseproto parse proto file
func (*protofile) parse_proto() {
	for _, fileinfor := range req.GetProtoFile() {
		file := fileinfor

		packagename := file.GetPackage()
		Information.push_package(packagename)

		// parse comment
		comments := protofileModule.parse_comment(file.SourceCodeInfo)

		// parse message
		for messageIndex, messageinfor := range file.GetMessageType() {
			message := messageinfor

			Information.Data[packagename].Messages[message.GetName()] = comments.parse_message(messageIndex, message)
		}

		// parse enum
		for enumIndex, enuminfor := range file.GetEnumType() {
			enum := enuminfor

			Information.Data[packagename].Enums[enum.GetName()] = comments.parse_enum(enumIndex, enum)
		}

		// parse service
		for serviceIndex, serviceinfor := range file.GetService() {
			service := serviceinfor

			Information.Data[packagename].Services[service.GetName()] = comments.parse_service(serviceIndex, service)
		}
	}
}

// parse_comment paarse comment in proto
func (*protofile) parse_comment(infor *descriptorpb.SourceCodeInfo) comments {
	comments := new_comments()
	for _, location := range infor.GetLocation() {
		if location.GetLeadingComments() == "" && location.GetTrailingComments() == "" && len(location.GetLeadingDetachedComments()) == 0 {
			continue
		}

		detached := make([]string, 0)
		for _, val := range location.GetLeadingDetachedComments() {
			detached = append(detached, strings.TrimSpace(strings.Replace(val, "\n ", "\n", -1)))
		}

		comments[fmt.Sprintf("%v", location.GetPath())] = &comment{
			Leading:         strings.TrimSpace(strings.ReplaceAll(location.GetLeadingComments(), "\n ", "\n")),
			Trailing:        strings.TrimSpace(strings.ReplaceAll(location.GetTrailingComments(), "\n ", "\n")),
			LeadingDetached: detached,
		}
	}
	return comments
}

// parseservice parse service in proto
func (cs comments) parse_service(serviceIndex int, protoService *descriptorpb.ServiceDescriptorProto) *Service {
	serviceInfor := Information.push_servce(protoService.GetName(), cs.parse(protoService.GetName(), PATH_SERVICE, serviceIndex))
	for methodindex, methodinfor := range protoService.GetMethod() {
		method := methodinfor
		serviceInfor.Methods = append(serviceInfor.Methods, cs.parse_method(serviceIndex, methodindex, method))
	}
	return serviceInfor
}

// parsemethod parse method in service
func (cs comments) parse_method(serviceIndex int, methodindex int, protoMethod *descriptorpb.MethodDescriptorProto) *Method {
	methodInfor := Information.push_method(protoMethod.GetName(), cs.parse(protoMethod.GetName(), PATH_SERVICE, serviceIndex, PATH_SERVICE_METHOD, methodindex))
	methodInfor.RequestParam = split_type(protoMethod.GetInputType())[1]
	methodInfor.ResponseParam = split_type(protoMethod.GetOutputType())[1]
	return methodInfor
}

// parse_message parse message in proto
func (cs comments) parse_message(messageIndex int, protoMessage *descriptorpb.DescriptorProto) *Message {
	messageInfor := Information.push_message(protoMessage.GetName(), cs.get(PATH_MESSAGE, messageIndex))
	for fieldIndex, fieldinfor := range protoMessage.GetField() {
		field := fieldinfor
		messageInfor.Fields = append(messageInfor.Fields, cs.parse_field(messageIndex, fieldIndex, messageInfor.MessageName, field))
	}
	return messageInfor
}

// parseenum parse enum in proto
func (cs comments) parse_enum(enumIndex int, protoEnum *descriptorpb.EnumDescriptorProto) *Enum {
	enumInfor := Information.push_enum(protoEnum.GetName(), cs.get(PATH_ENUM, enumIndex))
	for enumFieldIndex, enumfieldinfor := range protoEnum.GetValue() {
		enumField := enumfieldinfor
		enumInfor.Fields = append(enumInfor.Fields, cs.parse_enum_field(enumIndex, enumFieldIndex, enumField))
	}
	return enumInfor
}

// parse_enum_field parse field in enum
func (cs comments) parse_enum_field(enumIndex int, enumFieldIndex int, protoEnumField *descriptorpb.EnumValueDescriptorProto) *EnumField {
	return &EnumField{
		EnumFieldName:  protoEnumField.GetName(),
		EnumFieldValue: protoEnumField.GetNumber(),
		Description:    cs.get(PATH_ENUM, enumIndex, PATH_ENUM_VALUE, enumFieldIndex),
	}
}

// parse_field parse field in message
func (cs comments) parse_field(messageIndex int, fieldIndex int, messageName string, protoMessageField *descriptorpb.FieldDescriptorProto) *Field {
	fieldInfor := new(Field)

	fieldInfor.MessageName = messageName

	fieldInfor.ProtoName = protoMessageField.GetName()
	fieldInfor.ProtoLaber = protoMessageField.GetLabel()
	fieldInfor.ProtoType = protoMessageField.GetType()
	fieldInfor.ProtoNumber = protoMessageField.GetNumber()

	fieldInfor.JsonName = protoMessageField.GetJsonName()
	fieldInfor.JsonLabel = protoLabel2JsonLabel[protoMessageField.GetLabel()]
	fieldInfor.JsonType = protoType2JsonType[protoMessageField.GetType()]
	fieldInfor.JsonDefaultValue = jsonTypeDefaultValue[fieldInfor.JsonType]

	switch fieldInfor.JsonType {
	case TYPE_OBJECT:
		typename := split_type(protoMessageField.GetTypeName())
		fieldInfor.ProtoTypeName = typename[1]
		fieldInfor.ProtoPackagePath = typename[0]
		fieldInfor.ProtoFullTypeName = protoMessageField.GetTypeName()

		fieldInfor.JsonType = typename[1]
	case TPYE_NUMBER, TYPE_STRING, TYPE_BOOLEAN:
		fieldInfor.ProtoTypeName = descriptorpb.FieldDescriptorProto_Type_name[int32(fieldInfor.ProtoType)]
	}

	fieldInfor.Description = cs.get(PATH_MESSAGE, messageIndex, PATH_MESSAGE_FIELD, fieldIndex)

	return fieldInfor
}
