# Package auth

---
## 导航 <a name="top"> </a>
+ [服务](#srv)
+ [结构](#msg)
+ [枚举](#enu)
---

## 服务 <a name="srv"> </a>

+ <font size=3>Sys  [Sys 系统设置]</font>
  + <font size=2>[Get](#Get)   [Get 获取系统设置]</font>
  + <font size=2>[Del](#Del)   [Del 删除系统设置]</font>
  
+ <font size=3>User  [User]</font>
  + <font size=2>[Add](#Add)   [Add 添加用户]</font>
  + <font size=2>[List](#List)   [List 获取用户列表]</font>
  

---

## 函数
#### Get <a name="Get"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: Get 获取系统设置
路径: [auth/Sys/Get]
```
+ <font size=3>请求</font>

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>name</font> | <font size=2>[String](#String)</font> | <font size=2>可选</font> | <font size=2>name 名称</font> |

**<font size=2>示例</font>**
``` json
{
  "name": ""
}
```
+ <font size=3>响应</font>

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |
| <font size=2>data</font> | <font size=2>[System](#System)</font> | <font size=2>重复</font> | <font size=2>data 内容</font> |

**<font size=2>示例</font>**
``` json
{
  "data": [
    {
      "name": "",
      "routers": {
        "name": "",
        "routers": null
      }
    }
  ],
  "resp": {
    "code": 0,
    "mess": ""
  }
}
```
---
#### Del <a name="Del"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: Del 删除系统设置
路径: [auth/Sys/Del]
```
+ <font size=3>请求</font>

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>names</font> | <font size=2>[String](#String)</font> | <font size=2>重复</font> | <font size=2>names 名称</font> |

**<font size=2>示例</font>**
``` json
{
  "names": [
    ""
  ]
}
```
+ <font size=3>响应</font>

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |

**<font size=2>示例</font>**
``` json
{
  "resp": {
    "code": 0,
    "mess": ""
  }
}
```
---

#### Add <a name="Add"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: Add 添加用户
路径: [auth/User/Add]
```
+ <font size=3>请求</font>

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>user</font> | <font size=2>[UserInfo](#UserInfo)</font> | <font size=2>可选</font> | <font size=2>user 用户信息</font> |

**<font size=2>示例</font>**
``` json
{
  "user": {
    "age": 0,
    "name": "",
    "status": 0
  }
}
```
+ <font size=3>响应</font>

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |

**<font size=2>示例</font>**
``` json
{
  "resp": {
    "code": 0,
    "mess": ""
  }
}
```
---
#### List <a name="List"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: List 获取用户列表
路径: [auth/User/List]
```
+ <font size=3>请求</font>

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>name</font> | <font size=2>[String](#String)</font> | <font size=2>可选</font> | <font size=2>name 性别</font> |

**<font size=2>示例</font>**
``` json
{
  "name": ""
}
```
+ <font size=3>响应</font>

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |
| <font size=2>data</font> | <font size=2>[UserInfo](#UserInfo)</font> | <font size=2>重复</font> | <font size=2>内容</font> |

**<font size=2>示例</font>**
``` json
{
  "data": [
    {
      "age": 0,
      "name": "",
      "status": 0
    }
  ],
  "resp": {
    "code": 0,
    "mess": ""
  }
}
```
---



## 结构 <a name="msg"> </a>

| <font size=2>类型</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: |
| <font size=2>[AddRequest](#AddRequest)</font> | <font size=2>AddRequest rpc-Add request message</font> |
| <font size=2>[AddResponse](#AddResponse)</font> | <font size=2>AddResponse rpc-Add response message</font> |
| <font size=2>[DelRequest](#DelRequest)</font> | <font size=2>DelRequest rpc-Del request message</font> |
| <font size=2>[DelResponse](#DelResponse)</font> | <font size=2>DelResponse rpc-Del response message</font> |
| <font size=2>[GetRequest](#GetRequest)</font> | <font size=2>GetRequest rpc-Get request message</font> |
| <font size=2>[GetResponse](#GetResponse)</font> | <font size=2>GetResponse rpc-Get response message</font> |
| <font size=2>[ListRequest](#ListRequest)</font> | <font size=2>ListRequest rpc-List request message</font> |
| <font size=2>[ListResponse](#ListResponse)</font> | <font size=2>ListResponse rpc-List response message</font> |
| <font size=2>[Response](#Response)</font> | <font size=2>Response 返回</font> |
| <font size=2>[System](#System)</font> | <font size=2>System 系统设置信息</font> |
| <font size=2>[UserInfo](#UserInfo)</font> | <font size=2>UserInfo 用户信息</font> |

---
+ AddRequest <a name="AddRequest"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: AddRequest rpc-Add request message
```

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>user</font> | <font size=2>[UserInfo](#UserInfo)</font> | <font size=2>可选</font> | <font size=2>user 用户信息</font> |

+ AddResponse <a name="AddResponse"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: AddResponse rpc-Add response message
```

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |

+ DelRequest <a name="DelRequest"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: DelRequest rpc-Del request message
```

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>names</font> | <font size=2>[String](#String)</font> | <font size=2>重复</font> | <font size=2>names 名称</font> |

+ DelResponse <a name="DelResponse"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: DelResponse rpc-Del response message
```

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |

+ GetRequest <a name="GetRequest"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: GetRequest rpc-Get request message
```

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>name</font> | <font size=2>[String](#String)</font> | <font size=2>可选</font> | <font size=2>name 名称</font> |

+ GetResponse <a name="GetResponse"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: GetResponse rpc-Get response message
```

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |
| <font size=2>data</font> | <font size=2>[System](#System)</font> | <font size=2>重复</font> | <font size=2>data 内容</font> |

+ ListRequest <a name="ListRequest"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: ListRequest rpc-List request message
```

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>name</font> | <font size=2>[String](#String)</font> | <font size=2>可选</font> | <font size=2>name 性别</font> |

+ ListResponse <a name="ListResponse"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: ListResponse rpc-List response message
```

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |
| <font size=2>data</font> | <font size=2>[UserInfo](#UserInfo)</font> | <font size=2>重复</font> | <font size=2>内容</font> |

+ Response <a name="Response"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: Response 返回
```

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>code</font> | <font size=2>[Number](#Number)</font> | <font size=2>可选</font> | <font size=2>code 状态</font> |
| <font size=2>mess</font> | <font size=2>[String](#String)</font> | <font size=2>可选</font> | <font size=2>mess 信息</font> |

+ System <a name="System"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: System 系统设置信息
```

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>name</font> | <font size=2>[String](#String)</font> | <font size=2>可选</font> | <font size=2>name 名称</font> |
| <font size=2>routers</font> | <font size=2>[System](#System)</font> | <font size=2>可选</font> | <font size=2>routers 子项</font> |

+ UserInfo <a name="UserInfo"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: UserInfo 用户信息
```

| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>name</font> | <font size=2>[String](#String)</font> | <font size=2>可选</font> | <font size=2>name 姓名</font> |
| <font size=2>age</font> | <font size=2>[Number](#Number)</font> | <font size=2>可选</font> | <font size=2>age 年龄</font> |
| <font size=2>status</font> | <font size=2>[Status](#Status)</font> | <font size=2>可选</font> | <font size=2>status 状态</font> |



---
## 枚举 <a name="enu"> </a>

+ Gender <a name="Gender"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
| <font size=2>键</font> | <font size=2>值</font> | <font size=2>描述</font> |
| :--------------------: | :--------------------: | :---------------------: |
| <font size=2>Man</font> | <font size=2>0</font> | <font size=2>Gender:    Man</font> |
| <font size=2>Woman</font> | <font size=2>1</font> | <font size=2>Gender:    Woman</font> |

+ Status <a name="Status"> </a> <font size=2>[服务](#srv)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
| <font size=2>键</font> | <font size=2>值</font> | <font size=2>描述</font> |
| :--------------------: | :--------------------: | :---------------------: |
| <font size=2>Unknown</font> | <font size=2>0</font> | <font size=2>Status:    Unknown</font> |
| <font size=2>Enabled</font> | <font size=2>1</font> | <font size=2>Status:    Enabled</font> |
| <font size=2>Disable</font> | <font size=2>2</font> | <font size=2>Status:    Disable</font> |


---
