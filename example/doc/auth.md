# Package auth

---
## 导航 <a name="top"> </a>
+ [服务](#srv)
+ [函数](#mtd)
+ [结构](#msg)
+ [枚举](#enu)
---

## 服务 <a name="srv"> </a>

+ Sys 系统设置
  + [Get](#Get)   [Get 获取系统设置]
  + [Del](#Del)   [Del 删除系统设置]
  
+ User
  + [Add](#Add)   [Add 添加用户]
  + [List](#List)   [List 获取用户列表]
  

---

## 函数 <a name="mtd"> </a>
#### Get <a name="Get"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
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
#### Del <a name="Del"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
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

#### Add <a name="Add"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
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
#### List <a name="List"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
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
| [AddRequest](#AddRequest) | AddRequest rpc-Add request message |
| [AddResponse](#AddResponse) | AddResponse rpc-Add response message |
| [DelRequest](#DelRequest) | DelRequest rpc-Del request message |
| [DelResponse](#DelResponse) | DelResponse rpc-Del response message |
| [GetRequest](#GetRequest) | GetRequest rpc-Get request message |
| [GetResponse](#GetResponse) | GetResponse rpc-Get response message |
| [ListRequest](#ListRequest) | ListRequest rpc-List request message |
| [ListResponse](#ListResponse) | ListResponse rpc-List response message |
| [Response](#Response) | Response 返回 |
| [System](#System) | System 系统设置信息 |
| [UserInfo](#UserInfo) | UserInfo 用户信息 |

---
+ AddRequest <a name="AddRequest"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: AddRequest rpc-Add request message
```
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>user</font> | <font size=2>[UserInfo](#UserInfo)</font> | <font size=2>可选</font> | <font size=2>user 用户信息</font> |

+ AddResponse <a name="AddResponse"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: AddResponse rpc-Add response message
```
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |

+ DelRequest <a name="DelRequest"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: DelRequest rpc-Del request message
```
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>names</font> | <font size=2>[String](#String)</font> | <font size=2>重复</font> | <font size=2>names 名称</font> |

+ DelResponse <a name="DelResponse"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: DelResponse rpc-Del response message
```
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |

+ GetRequest <a name="GetRequest"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: GetRequest rpc-Get request message
```
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>name</font> | <font size=2>[String](#String)</font> | <font size=2>可选</font> | <font size=2>name 名称</font> |

+ GetResponse <a name="GetResponse"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: GetResponse rpc-Get response message
```
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |
| <font size=2>data</font> | <font size=2>[System](#System)</font> | <font size=2>重复</font> | <font size=2>data 内容</font> |

+ ListRequest <a name="ListRequest"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: ListRequest rpc-List request message
```
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>name</font> | <font size=2>[String](#String)</font> | <font size=2>可选</font> | <font size=2>name 性别</font> |

+ ListResponse <a name="ListResponse"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: ListResponse rpc-List response message
```
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>resp</font> | <font size=2>[Response](#Response)</font> | <font size=2>可选</font> | <font size=2>resp 返回</font> |
| <font size=2>data</font> | <font size=2>[UserInfo](#UserInfo)</font> | <font size=2>重复</font> | <font size=2>内容</font> |

+ Response <a name="Response"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: Response 返回
```
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>code</font> | <font size=2>[Number](#Number)</font> | <font size=2>可选</font> | <font size=2>code 状态</font> |
| <font size=2>mess</font> | <font size=2>[String](#String)</font> | <font size=2>可选</font> | <font size=2>mess 信息</font> |

+ System <a name="System"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
```
描述: System 系统设置信息
```
| <font size=2>字段</font> | <font size=2>类型</font> | <font size=2>标签</font> | <font size=2>描述</font> |
| :----------------------: | :---------------------: | :----------------------: | :----------------------: |
| <font size=2>name</font> | <font size=2>[String](#String)</font> | <font size=2>可选</font> | <font size=2>name 名称</font> |
| <font size=2>routers</font> | <font size=2>[System](#System)</font> | <font size=2>可选</font> | <font size=2>routers 子项</font> |

+ UserInfo <a name="UserInfo"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
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
+ Gender <a name="Gender"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
| <font size=2>键</font> | <font size=2>值</font> | <font size=2>描述</font> |
| :--------------------: | :--------------------: | :---------------------: |
| Man | 0 | :  |
| Woman | 1 | :  |

+ Status <a name="Status"> </a> <font size=2>[服务](#srv)</font> <font size=2>[函数](#mtd)</font> <font size=2>[结构](#msg)</font> <font size=2>[枚举](#enu)</font>
| <font size=2>键</font> | <font size=2>值</font> | <font size=2>描述</font> |
| :--------------------: | :--------------------: | :---------------------: |
| Unknown | 0 | :  |
| Enabled | 1 | :  |
| Disable | 2 | :  |


---
