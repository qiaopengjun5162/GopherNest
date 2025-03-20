package tests

import (
	"github.com/qiaopengjun5162/GopherNest/experiments"
	"strings"
	"testing"
)

// TestToJSON_ValidInput_ReturnsJSON tests that ToJSON correctly converts a valid input to JSON.
func TestToJSON_ValidInput_ReturnsJSON(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "John Doe", Age: 30}
	expectedJSON := `{"Name":"John Doe","Age":30}`

	actualJSON := experiments.ToJSON(person)

	if string(actualJSON) != expectedJSON {
		t.Errorf("ToJSON(%v) = %s; want %s", person, actualJSON, expectedJSON)
	}
}

// TestToJSON_InvalidInput_ReturnsEmpty tests that ToJSON returns an empty byte slice for invalid input.
func TestToJSON_InvalidInput_ReturnsEmpty(t *testing.T) {
	type Invalid struct {
		Foo func() // Functions are not JSON serializable
	}

	invalid := Invalid{Foo: func() {}}

	actualJSON := experiments.ToJSON(invalid)

	if len(actualJSON) != 0 {
		t.Errorf("ToJSON(%v) = %s; want empty byte slice", invalid, actualJSON)
	}
}

// TestToJSONString_NilInput_ReturnsNull tests the behavior of ToJSONString when the input is nil.
func TestToJSONString_NilInput_ReturnsNull(t *testing.T) {
	result := experiments.ToJSONString(nil)
	if result != "null" {
		t.Errorf("Expected 'null' for nil input, got '%s'", result)
	}
}

// TestToJSONString_ValidInput_ReturnsJSONString tests the behavior of ToJSONString with a valid input.
func TestToJSONString_ValidInput_ReturnsJSONString(t *testing.T) {
	input := map[string]string{"key": "value"}
	expected := `{"key":"value"}`
	result := experiments.ToJSONString(input)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

// TestToJSONString_UnserializableInput_ReturnsEmptyString tests the behavior of ToJSONString with an unserializable input.
func TestToJSONString_UnserializableInput_ReturnsEmptyString(t *testing.T) {
	type Unserializable struct {
		Chan chan int
	}
	input := Unserializable{Chan: make(chan int)}
	result := experiments.ToJSONString(input)
	if result != "" {
		t.Errorf("Expected empty string for unserializable input, got '%s'", result)
	}
}

func TestToPrettyJSON_NilInput_ReturnsNull(t *testing.T) {
	result := experiments.ToPrettyJSON(nil)
	if result != "null" {
		t.Errorf("Expected 'null' for nil input, got %s", result)
	}
}

func TestToPrettyJSON_ValidStruct_ReturnsFormattedJSON(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	person := Person{Name: "John Doe", Age: 30}
	expected := `{
  "name": "John Doe",
  "age": 30
}`
	result := experiments.ToPrettyJSON(person)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestToPrettyJSON_CircularReference_ReturnsErrorJSON(t *testing.T) {
	type Node struct {
		Next *Node `json:"next"`
	}
	node := &Node{}
	node.Next = node // 创建循环引用

	result := experiments.ToPrettyJSON(node)
	if !strings.Contains(result, "error") {
		t.Errorf("Expected error JSON for circular reference, got %s", result)
	}
}
