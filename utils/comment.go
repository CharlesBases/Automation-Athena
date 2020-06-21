package utils

import (
	"encoding/json"
	"fmt"
)

// comment path
const (
	// tag numbers in FileDescriptorProto
	PATH_PACKAGE   = 2  // package comment
	PATH_MESSAGE   = 4  // message comment
	PATH_ENUM      = 5  // enum comment
	PATH_SERVICE   = 6  // service comment
	PATH_EXTENSION = 7  // extension comment
	PATH_SYNTAX    = 12 // syntax comment

	// tag numbers in DescriptorProto
	PATH_MESSAGE_FIELD     = 2 // message.field
	PATH_MESSAGE_MESSAGE   = 3 // message.nested
	PATH_MESSAGE_ENUM      = 4 // message.enum
	PATH_MESSAGE_EXTENSION = 6 // message.ectension

	// tag numbers in EnumDescriptorProto
	PATH_ENUM_VALUE = 2 // value

	// tag numbers in ServiceDescriptorProto
	PATH_SERVICE_METHOD = 2
)

type (
	comments map[string]*comment

	comment struct {
		Leading         string
		Trailing        string
		LeadingDetached []string
	}
)

// get new comments
func new_comments() comments {
	return make(map[string]*comment, 0)
}

// get get comment by path
func (cs comments) get(name string, path ...int) string {
	if comment, ok := cs[fmt.Sprintf("%v", path)]; ok {
		if comment.Leading != "" {
			return comment.Leading
		}
	}
	return name
}

// parse parse comment to desc and uri
func (cs comments) parse(name string, path ...int) *desc {
	source := ""
	desc := new(desc)
	if json.Unmarshal([]byte(source), desc) != nil {
		desc.path = name
		if leading, ok := cs[fmt.Sprintf("%v", path)]; ok {
			desc.desc = leading.Leading
		}
	}
	if desc.desc == "" {
		desc.desc = name
	}
	if desc.path == "" {
		desc.path = name
	}
	return desc
}
