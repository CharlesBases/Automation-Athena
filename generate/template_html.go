package generate

const HTML = `{{$packagename := .PackageName -}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>{{$packagename}}</title>
    <meta charset="UTF-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1"/>
    <style type="text/css">
      html {
        line-height: 140%;
      }
      body {
        margin: 10px;
      }
      table {
        border-spacing: 0;
        border-padding: 0;
        border-collapse: collapse;
      }
      td, th {
        padding: 0;
        spacing: 0;
      }
      .pure-table {
        border-collapse: collapse;
        border-spacing: 0;
        empty-cells: show;
        border: 1px solid #CBCBCB;
      }
      .pure-table td {
        background-color: transparent;
        border-left: 1px solid #cbcbcb;
        border-width: 0 0 0 1px;
        font-size: inherit;
        margin: 0;
        overflow: visible;
        padding: .5em 1em;
      }
      .pure-table th {
        border-left: 1px solid #CBCBCB;
        border-width: 0 0 0 1px;
        font-size: inherit;
        margin: 0;
        overflow: visible;
        padding: .5em 1em;
      }
      .pure-table thead {
        background-color: #DCDCDC;
        color: #000000;
        text-align: center;
        vertical-align: bottom;
      }
      .pure-table-odd td {
        background-color: #F1F1F1;
      }
      .arrow {
        border: 9px solid transparent;
        border-bottom-color: #3DA0DB;
        width: 0px;
        height: 0px;
        top:0px
      }
      .stick {
        width: 8px;
        height: 14px;
        border-radius: 1px;
        background-color: #3DA0DB;
        top:15px;
      }
      #back_top div {
        position: absolute;
        margin: auto;
        right: 0px;
        left: 0px;
      }
      #back_top {
        background-color: #DDDDDD;
        height: 38px;
        width: 38px;
        border-radius: 3px;
        display: block;
        cursor: pointer;
        position: fixed;
        right: 50px;
        bottom: 100px;
        display: none;
      }
      #header {
        background-color: #000000;
        color: #DCDCDC;
        text-align: center;
        padding: 5px;
      }
      #codeblock {
        background-color: #F5F5F5;
        padding-left: 5px;
        padding-bottom: 1px;
        text-align: left;
        border: 1px solid #CBCBCB;
      }
    </style>
  </head>

  <body>
    <div id="article"></div>
    <div id="back_top">
      <div class="arrow"></div>
      <div class="stick"></div>
    </div>
    <script src="http://cdn.staticfile.org/jquery/1.11.1-rc2/jquery.min.js"></script>
    <script>
        $(function(){
          $(window).scroll(function(){
            var scrollt = document.documentElement.scrollTop + document.body.scrollTop;
            if( scrollt > 200 ){
              $("#back_top").fadeIn(400);
            }else{
              $("#back_top").stop().fadeOut(400);
            }
          });
          $("#back_top").click(function(){
            $("html,body").animate({scrollTop:"0px"},200);
          });
        });
    </script>

    <!-- 接口文档 -->
    <div id="header"><h1>{{$packagename}}</h1></div>
    <code>
    <h1>导航</h1>
    <ul>
      <li><a href="#srv">服务</a></li>
      <li><a href="#msg">结构</a></li>
      <li><a href="#enu">枚举</a></li>
    </ul>
    <h1><a id="srv">服务</a></h1>
    <ul>
    {{range $serviceName, $service := .Services -}}
      <li>{{$service.ServiceName}}
        <ul>
        {{range $methodindex, $method := $service.Methods -}}
        <li><a href="#{{$service.ServiceName}}.{{$method.MethodName}}">{{$method.MethodName}}</a>{{dynamic $method.MethodName}}[{{$method.Description}}]</li>
        {{end}}
        </ul>
      </li>
    {{end}}
    </ul>
    <HR>
    <h1>函数</h1>
    {{range $serviceName, $service := .Services -}}
    {{range $methodindex, $method := $service.Methods -}}
    <h2><a id="{{$service.ServiceName}}.{{$method.MethodName}}">{{$method.MethodName}}</a></h2>
    <div id="codeblock"><font color="#696969">
    描述: {{$method.Description}}</br>
    服务: {{$service.Uri}}</br>
    路径: [{{$packagename}}/{{$service.Uri}}/{{$method.Uri}}]</br>
    </font></div>
    <h3>请求</h3>
    {{$request := parsemessage $method.RequestParam -}}
    <table class="pure-table">
      <thead>
        <tr>
          <td>字段</td>
          <td>类型</td>
          <td>标签</td>
          <td>描述</td>
        </tr>
      </thead>
      <tbody>
        {{$index := 1}}{{range $fieldindex, $field := $request.Fields -}}
        <tr {{if polling $index}}class="pure-table-odd"{{end}}{{$index = increasing $index}}>
          <td>{{$field.JsonName}}</td>
          <td><a href="#{{$field.JsonType}}">{{$field.JsonType}}</a></td>
          <td>{{$field.JsonLabel}}</td>
          <td>{{$field.Description}}</td>
        </tr>
        {{end}}
      <tbody>
    </table>
    <h4>示例</h4>
    <pre><div id="codeblock">{{jsonparse $request.MessageName}}</pre></pre>
    <h3>响应</h3>
    {{$response := parsemessage $method.ResponseParam -}}
    <table class="pure-table">
      <thead>
        <tr>
          <td>字段</td>
          <td>类型</td>
          <td>标签</td>
          <td>描述</td>
        </tr>
      </thead>
      <tbody>
      {{$index := 1}}{{range $fieldindex, $field := $response.Fields -}}
        <tr {{if polling $index}}class="pure-table-odd"{{end}}{{$index = increasing $index}}>
          <td>{{$field.JsonName}}</td>
          <td><a href="#{{$field.JsonType}}">{{$field.JsonType}}</a></td>
          <td>{{$field.JsonLabel}}</td>
          <td>{{$field.Description}}</td>
        </tr>
      {{end}}
      <tbody>
    </table>
    <h4>示例</h4>
    <pre><div id="codeblock">{{jsonparse $response.MessageName}}</div></pre>
    {{end}}
    {{end}}

    <h1><a id="msg">结构</a></h1>
    <table class="pure-table">
      <thead>
        <tr>
          <td>类型</td>
          <td>描述</td>
        </tr>
      </thead>
      <tbody>
        {{$index := 1}}{{range $messagename, $message := .Messages -}}
        <tr {{if polling $index}}class="pure-table-odd"{{end}}{{$index = increasing $index}}>
          <td><a href="#{{$message.MessageName}}">{{$message.MessageName}}</a></td>
          <td>{{$message.Description}}</td>
        </tr>
        {{end}}
      </tbody>
    </table>

    <!-- 结构列表 -->
    {{range $messagename, $message := .Messages -}}
    <ul>
      <li><h3><a id="{{$message.MessageName}}">{{$message.MessageName}}</a></h3></li>
      <p><font color="#696969">说明: {{$message.Description}}</font></p>
      <table class="pure-table">
        <thead>
          <tr>
            <td>字段</td>
            <td>类型</td>
            <td>标签</td>
            <td>描述</td>
          </tr>
        </thead>
        <tbody>
          {{$index := 1}}{{range $fieldindex, $field := $message.Fields -}}
          <tr {{if polling $index}}class="pure-table-odd"{{end}}{{$index = increasing $index}}>
            <td>{{$field.JsonName}}</td>
            <td><a href="#{{$field.JsonType}}">{{$field.JsonType}}</a></td>
            <td>{{$field.JsonLabel}}</td>
            <td>{{$field.Description}}</td>
          </tr>
          {{end}}
        </tbody>
      </table>
    </ul>
    {{end}}

    <h1><a id="enu">枚举</a></h1>
    {{range $enumname, $enum := .Enums -}}
    <ul>
      <li><h4><a id="{{$enum.EnumName}}">{{$enum.EnumName}}</a></h4></li>
      <table class="pure-table">
        <thead>
          <tr>
            <td>键</td>
            <td>值</td>
            <td>描述</td>
          </tr>
        </thead>
        <tbody>
          {{$index := 1}}{{range $fieldindex, $field := $enum.Fields -}}
          <tr {{if polling $index}}class="pure-table-odd"{{end}}{{$index = increasing $index}}>
            <td>{{$field.EnumFieldName}}</td>
            <td>{{$field.EnumFieldValue}}</td>
            <td>{{$enum.Description}}:  {{$field.Description}}</td>
          </tr>
          {{end}}
        </tbody>
      </table>
    </ul>
    {{end}}
    </code>
  </body>
</html>
`
