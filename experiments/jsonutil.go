package experiments

import (
	"encoding/json"
	"fmt"
)

// ToJSON 将一个 Go 语言值转换为 JSON 格式的字节切片。
// 该函数接受一个 interface{} 类型的参数 v，这意味着它可以接受任何类型的值。
// 使用 json.Marshal 函数将 v 转换为 JSON 格式。如果转换过程中发生错误，这个错误将被忽略。
// 返回值是转换后的 JSON 格式的字节切片。如果 v 无法被转换为 JSON 格式，则返回空的字节切片。
func ToJSON(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

// ToJSONString 将给定的值转换为JSON格式的字符串。
// 如果输入值为nil，则根据需求决定返回"null"或空字符串。
// 参数:
//
//	v - 要转换为JSON字符串的值。
//
// 返回值:
//
//	JSON格式的字符串表示，如果发生错误则返回空字符串。
func ToJSONString(v interface{}) string {
	// 检查输入值是否为nil，如果是，则根据需求返回"null"或空字符串。
	if v == nil {
		// 根据需求决定是否返回 "null" 或空字符串
		return "null"
	}

	// 尝试将输入值转换为JSON格式的字节切片。
	b, err := json.Marshal(v)
	// 如果转换过程中发生错误，则记录错误信息并返回空字符串。
	if err != nil {
		// 记录错误日志（如果需要）
		fmt.Println("Error during JSON marshaling:", err)
		// 返回空字符串或自定义错误信息
		return ""
	}
	// 如果转换成功，则将字节切片转换为字符串并返回。
	return string(b)
}

// ToPrettyJSON 将给定的接口类型数据转换为格式化的 JSON 字符串。
// 如果输入数据为 nil，返回 "null" 的 JSON 表示。
// 如果输入数据无法序列化为 JSON，返回包含错误信息的 JSON 字符串。
func ToPrettyJSON(v interface{}) string {
	// 如果 v 为 nil，直接返回 "null" 的 JSON 表示
	if v == nil {
		return "null"
	}

	// 使用 json.MarshalIndent 将数据格式化为带缩进的 JSON 字符串
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		// 如果序列化失败，返回一个明确的错误提示
		return fmt.Sprintf(`{"error": "%v"}`, err)
	}

	// 返回格式化后的 JSON 字符串
	return string(b)
}
