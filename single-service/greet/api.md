# API 规范

* https://go-zero.dev/docs/tutorials




# 类型声明

* https://go-zero.dev/docs/tutorials/api/types

> 暂不支持 泛型, 弱类型, 比如 any 等


在 API 描述语言中，类型声明需要满足如下规则:

* 类型声明必须以 type 开头
* 不需要声明 struct关键字
* 不支持嵌套结构体声明
* 不支持别名

## 示例1

```api
type Request {
	Name string `path:"name,options=you|me"`
}

type Response {
	Message string `json:"message"`
}

service greet-api {
	@handler GreetHandler
	get /from/:name(Request) returns (Response)
}
```

## 示例2

```api

type StructureExample {
    // 基本数据类型示例
    BaseInt     int     `json:"base_int"`
    BaseBool    bool    `json:"base_bool"`
    BaseString  string  `json:"base_string"`
    BaseByte    byte    `json:"base_byte"`
    BaseFloat32 float32 `json:"base_float32"`
    BaseFloat64 float64 `json:"base_float64"`
    // 切片示例
    BaseIntSlice     []int     `json:"base_int_slice"`
    BaseBoolSlice    []bool    `json:"base_bool_slice"`
    BaseStringSlice  []string  `json:"base_string_slice"`
    BaseByteSlice    []byte    `json:"base_byte_slice"`
    BaseFloat32Slice []float32 `json:"base_float32_slice"`
    BaseFloat64Slice []float64 `json:"base_float64_slice"`
    // map 示例
    BaseMapIntString      map[int]string               `json:"base_map_int_string"`
    BaseMapStringInt      map[string]int               `json:"base_map_string_int"`
    BaseMapStringStruct   map[string]*StructureExample `json:"base_map_string_struct"`
    BaseMapStringIntArray map[string][]int             `json:"base_map_string_int_array"`
    // 匿名示例
    *Base
    // 指针示例
    Base4 *Base `json:"base4"`
    
    // 新的特性（ goctl >= 1.5.1 版本支持 ）
    // 标签忽略示例
    TagOmit string
}
```