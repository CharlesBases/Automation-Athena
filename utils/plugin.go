package utils

import (
	"fmt"
	"io"
	"os"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

var Information = new(infor)

type infor struct {
	Data map[string]*Package // 包名 ｜ 内容
}

func init() {
	Information.Data = make(map[string]*Package, 0)
}

// Plugin proto 文件解析器
func Plugin() {
	protofileModule.unmarshal()
	protofileModule.parse_proto()
}

// Push push response to response
func Push(filename string, data []byte) {
	rsp.File = append(
		rsp.File,
		&pluginpb.CodeGeneratorResponse_File{
			Name:    proto.String(filename),
			Content: proto.String(string(data)),
		},
	)
}

// WriteResponse return to os.Stdout
func WriteResponse(writer io.Writer) {
	data, err := proto.Marshal(rsp)
	ThrowCheck(err)
	_, err = writer.Write(data)
	ThrowCheck(err)
}

// push_package 包名不存在则新增
func (*infor) push_package(packagename string) {
	if _, ok := Information.Data[packagename]; ok {
		return
	}
	Information.Data[packagename] = &Package{
		PackageName: packagename,
		Services:    make(map[string]*Service, 0),
		Messages:    make(map[string]*Message, 0),
		Enums:       make(map[string]*Enum, 0),
	}
}

// push_servce new service
func (*infor) push_servce(serviceNamme string, desc *desc) *Service {
	return &Service{
		ServiceName: serviceNamme,
		Methods:     make([]*Method, 0),
		Description: desc.desc,
		Uri:         desc.path,
	}
}

// push_method new method
func (*infor) push_method(methodName string, desc *desc) *Method {
	return &Method{
		MethodName:  methodName,
		Description: desc.desc,
		Uri:         desc.path,
	}
}

// push_message new message
func (*infor) push_message(messageName string, description string) *Message {
	return &Message{
		MessageName: messageName,
		Fields:      make([]*Field, 0),
		Description: description,
	}
}

// push_enum new enum
func (*infor) push_enum(enumName string, description string) *Enum {
	return &Enum{
		EnumName:    enumName,
		Fields:      make([]*EnumField, 0),
		Description: description,
	}
}

// split_type split by "." and return package and type name
func split_type(typename string) [2]string {
	list := strings.Split(typename, ".")
	if len(list) != 3 {
		ThrowCheck(fmt.Errorf("split type error: [%s]", typename))
	}
	return [2]string{list[1], list[2]}
}

// tofile .
func tofile(datas []byte) {
	file, _ := os.OpenFile("protofile", os.O_CREATE|os.O_RDWR, 0755)
	defer file.Close()

	file.Write(datas)
}
