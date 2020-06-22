package generate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"

	"protoc-gen-doc/utils"
)

// Markdown generate markdown
func (g *generate) Markdown() []byte {
	temp := template.New("Markdown Template")

	temp.Funcs(template.FuncMap{
		"codeblock": func(languages ...string) template.HTML {
			for x := range languages {
				return template.HTML(fmt.Sprintf("``` %s", languages[x]))
			}
			return template.HTML("```")
		},
		"parsemessage": g.parse_message,
		"jsonparse":    g.json_parse,
	})
	markdown, err := temp.Parse(markdownTemplate)
	utils.ThrowCheck(err)

	var buffer bytes.Buffer
	utils.ThrowCheck(markdown.Execute(&buffer, g.Data))

	return buffer.Bytes()
}

// parse_message get message information by message name
func (g *generate) parse_message(name string) *utils.Message {
	return g.Data.Messages[name]
}

// json_parse json parse for message
func (g *generate) json_parse(messageName string) template.HTML {
	statistics := make(map[string]bool, 0)
	data, _ := json.MarshalIndent(g.encoder.encode(g, g.Data.Messages[messageName].Fields, statistics), "", "  ")
	return template.HTML(string(data))
}

// encode message fields to json demo. statistics: 预防同名结构体嵌套导致 goroutine 堆栈字节溢出
func (encoder *encode) encode(g *generate, fields []*utils.Field, statistics map[string]bool) interface{} {
	data := make(map[string]interface{}, 0)
	for _, field := range fields {
		data[field.JsonName] = encoder.encode_field(g, field, statistics)
	}
	return data
}

// encode_field parse field
func (encoder *encode) encode_field(g *generate, field *utils.Field, statistics map[string]bool) interface{} {
	if statistics[field.JsonType] {
		return nil
	}

	var data interface{}

	switch field.JsonType {
	case utils.TPYE_NUMBER, utils.TYPE_STRING, utils.TYPE_BOOLEAN:
		data = field.JsonDefaultValue
	default:
		if field.JsonType == field.MessageName {
			statistics[field.MessageName] = true
		}
		if _, ok := g.Data.Enums[field.JsonType]; ok {
			data = 0
		} else {
			data = encoder.encode(g, g.Data.Messages[field.ProtoTypeName].Fields, statistics)
		}
	}

	switch field.JsonLabel {
	case utils.LABEL_REPEATED:
		return [1]interface{}{data}
	default:
		return data
	}
}
