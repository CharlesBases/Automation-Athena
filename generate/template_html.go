package generate

const HTML = `{{$packagename := .PackageName -}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>{{$packagename}}</title>
    <meta charset="UTF-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1"/>
    <style type="text/css">
      html{
        font-family: sans-serif;
        -ms-text-size-adjust: 100%;
        -webkit-text-size-adjust: 100%;
      }
      body{
        margin: 10px;
      }
      table{
        border-spacing: 0;
        border-padding: 0;
        border-collapse: collapse;
      }
      td, th{
        padding: 0;
        spacing: 0;
      }
      .pure-table{
        border-collapse: collapse;
        border-spacing: 0;
        empty-cells: show;
        border: 1px solid #CBCBCB;
      }
      .pure-table td{
        background-color: transparent;
        border-left: 1px solid #cbcbcb;
        border-width: 0 0 0 1px;
        font-size: inherit;
        margin: 0;
        overflow: visible;
        padding: .5em 1em;
      }
      .pure-table th{
        border-left: 1px solid #cbcbcb;
        border-width: 0 0 0 1px;
        font-size: inherit;
        margin: 0;
        overflow: visible;
        padding: .5em 1em;
      }
      .pure-table thead{
        background-color: #E0E0E0;
        color: #000;
        text-align: center;
        vertical-align: bottom;
      }
      .pure-table-odd td{
        background-color: #F2F2F2;
      }
      .arrow{
        border: 9px solid transparent;
        border-bottom-color: #3DA0DB;
        width: 0px;
        height: 0px;
        top:0px
      }
      .stick{
        width: 8px;
        height: 14px;
        border-radius: 1px;
        background-color: #3DA0DB;
        top:15px;
      }
      #back_top div{
        position: absolute;
        margin: auto;
        right: 0px;
        left: 0px;
      }
      #back_top{
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
    <code>
    <h1>Package {{$packagename}}</h1>

    <h2>导航</h2>
    <ul>
      <li><a href="#srv">服务</a></li>
      <li><a href="#msg">结构</a></li>
      <li><a href="#enu">枚举</a></li>
    </ul>

    <h2><a id="srv">服务</a></h2>
    <ul>
    {{range $serviceName, $service := .Services -}}
      <li>{{$service.ServiceName}}
        <ul>
        {{range $methodindex, $method := $service.Methods -}}
        <li><a href="#{{$service.ServiceName}}.{{$method.MethodName}}">{{$method.MethodName}}</a>[{{$method.Description}}]</li>
        {{end}}
        </ul>
      </li>
    {{end}}
    </ul>
    
    <h2>函数</h2>
    {{range $serviceName, $service := .Services -}}
    {{range $methodindex, $method := $service.Methods -}}
    <h4><a id="{{$service.ServiceName}}.{{$method.MethodName}}">{{$method.MethodName}}</a></h4>
    服务: {{$service.Uri}}</br>
    描述: {{$method.Description}}</br>
    路径: [ {{$packagename}}/{{$service.Uri}}/{{$method.Uri}} ]</br>
    <h5>请求</h5>
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
    <h6>示例</h6>
    <pre>
{{jsonparse $request.MessageName}}
    </pre>
    <h5>响应</h5>
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
    <h6>示例</h6>
    <pre>
{{jsonparse $response.MessageName}}
    </pre>
    {{end}}
    {{end}}

    <h2><a id="msg">结构</a></h2>
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
      <li><a id="{{$message.MessageName}}">{{$message.MessageName}}</a></li>
      描述: {{$message.Description}}
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

    <h2><a id="enu">枚举</a></h2>
    {{range $enumname, $enum := .Enums -}}
    <ul>
      <li><a id="{{$enum.EnumName}}">{{$enum.EnumName}}</a></li>
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
    <code>
  </body>
</html>
`
