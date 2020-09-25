package generate

const Markdown = `# {{$packagename := .PackageName -}}
Package {{$packagename}}

---
## 导航 <a name="top"> </a>
+ [服务](#srv)
+ [结构](#msg)
+ [枚举](#enu)
---

## 服务 <a name="srv"> </a>

{{range $serviceName, $service := .Services -}}
+ <font size=3>{{$service.ServiceName}}  [{{$service.Description}}]</font>
  {{range $methodindex, $method := $service.Methods -}}
  + <font size=2>[{{$method.MethodName}}](#{{$service.ServiceName}}.{{$method.MethodName}})   [{{$method.Description}}]</font>
  {{end}}
{{end}}
---

## 函数
{{range $serviceName, $service := .Services -}}
{{range $methodindex, $method := $service.Methods -}}
#### {{$method.MethodName}} <a name="{{$service.ServiceName}}.{{$method.MethodName}}"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
{{codeblock}}
描述: {{$method.Description}}
路径: [{{$packagename}}/{{$service.Uri}}/{{$method.Uri}}]
{{codeblock}}
+ <font size=3>请求</font>

{{$message := parsemessage $method.RequestParam -}}
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
{{range $fieldindex, $field := $message.Fields -}}
| <font size=2>{{$field.JsonName}}</font> | <font size=2>[{{$field.JsonType}}](#{{$field.JsonType}})</font> | <font size=2>{{$field.JsonLabel}}</font> | <font size=2>{{$field.Description}}</font> |
{{end}}
**<font size=2>示例</font>**
{{codeblock "json"}}
{{jsonparse $message.MessageName}}
{{codeblock}}
+ <font size=3>响应</font>

{{$message := parsemessage $method.ResponseParam -}}
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
{{range $fieldindex, $field := $message.Fields -}}
| <font size=2>{{$field.JsonName}}</font> | <font size=2>[{{$field.JsonType}}](#{{$field.JsonType}})</font> | <font size=2>{{$field.JsonLabel}}</font> | <font size=2>{{$field.Description}}</font> |
{{end}}
**<font size=2>示例</font>**
{{codeblock "json"}}
{{jsonparse $message.MessageName}}
{{codeblock}}
---
{{end}}
{{end}}

## 结构 <a name="msg"> </a>

| <font size=2>类型</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: |
{{range $messagename, $message := .Messages -}}
| <font size=2>[{{$message.MessageName}}](#{{$message.MessageName}})</font> | <font size=2>{{$message.Description}}</font> |
{{end}}
---
{{range $messagename, $message := .Messages -}}
+ {{$message.MessageName}} <a name="{{$message.MessageName}}"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
{{codeblock}}
描述: {{$message.Description}}
{{codeblock}}

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
{{range $fieldindex, $field := $message.Fields -}}
| <font size=2>{{$field.JsonName}}</font> | <font size=2>[{{$field.JsonType}}](#{{$field.JsonType}})</font> | <font size=2>{{$field.JsonLabel}}</font> | <font size=2>{{$field.Description}}</font> |
{{end}}
{{end}}

---
## 枚举 <a name="enu"> </a>

{{range $enumname, $enum := .Enums -}}
+ {{$enum.EnumName}} <a name="{{$enum.EnumName}}"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
| <font size=2>键</font> | <font size=2>值</font> | <font size=2>描述</font> |
| :--------------------: | :--------------------: | :---------------------: |
{{range $fieldindex, $field := $enum.Fields -}}
| <font size=2>{{$field.EnumFieldName}}</font> | <font size=2>{{$field.EnumFieldValue}}</font> | <font size=2>{{$enum.Description}}:    {{$field.Description}}</font> |
{{end}}
{{end}}
---
`
