package jsonw

import (
	"encoding/json"
	"testing"
)

// Helper function to validate JSON produced by the writer.
func assertValidJSON(t *testing.T, data []byte) {
	t.Helper()
	var tmp interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		t.Fatalf("Invalid JSON: %s, error: %v", string(data), err)
	}
}

func TestArrayWriter_Primitives(t *testing.T) {
	w := NewArrayWriter(nil)

	w.Open()
	w.StringValue("hello")
	w.IntegerValue(42)
	w.FloatValue(3.14)
	w.BoolValue(true)
	w.NullValue()
	w.Close()

	data, err := w.BuildBytes()
	if err != nil {
		t.Fatal(err)
	}

	assertValidJSON(t, data)
	expected := `["hello",42,3.14,true,null]`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

func TestArrayWriter_Nested(t *testing.T) {
	w := NewArrayWriter(nil)

	w.Open()

	arr := w.ArrayValue()
	arr.Open()
	arr.IntegerValue(1)
	arr.IntegerValue(2)
	arr.Close()

	obj := w.ObjectValue()
	obj.Open()
	obj.StringField("name", "test")
	obj.IntegerField("value", 100)
	obj.Close()

	w.Close()

	data, err := w.BuildBytes()
	if err != nil {
		t.Fatal(err)
	}

	assertValidJSON(t, data)
	expected := `[[1,2],{"name":"test","value":100}]`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

func TestObjectWriter_Primitives(t *testing.T) {
	w := NewObjectWriter(nil)

	w.Open()
	w.StringField("str", "hello")
	w.IntegerField("int", 42)
	w.FloatField("float", 3.14)
	w.BoolField("bool", true)
	w.NullField("null")
	w.Close()

	data, err := w.BuildBytes()
	if err != nil {
		t.Fatal(err)
	}

	assertValidJSON(t, data)
	expected := `{"str":"hello","int":42,"float":3.14,"bool":true,"null":null}`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

func TestObjectWriter_Nested(t *testing.T) {
	w := NewObjectWriter(nil)

	w.Open()
	arr := w.ArrayField("arr")
	arr.Open()
	arr.IntegerValue(1)
	arr.IntegerValue(2)
	arr.Close()

	obj := w.ObjectField("obj")
	obj.Open()
	obj.StringField("key", "value")
	obj.IntegerField("num", 10)
	obj.Close()

	w.Close()

	data, err := w.BuildBytes()
	if err != nil {
		t.Fatal(err)
	}

	assertValidJSON(t, data)
	expected := `{"arr":[1,2],"obj":{"key":"value","num":10}}`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

func TestStringEscaping(t *testing.T) {
	w := NewObjectWriter(nil)

	w.Open()
	w.StringField("special", "line\nquote\"back\\slash\t")
	w.Close()

	data, err := w.BuildBytes()
	if err != nil {
		t.Fatal(err)
	}

	assertValidJSON(t, data)
	expected := `{"special":"line\nquote\"back\\slash\t"}`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

func TestEmptyArrayAndObject(t *testing.T) {
	arr := NewArrayWriter(nil)
	arr.Open()
	arr.Close()

	data, err := arr.BuildBytes()
	if err != nil {
		t.Fatal(err)
	}
	assertValidJSON(t, data)
	if string(data) != "[]" {
		t.Errorf("Expected [], got %s", string(data))
	}

	obj := NewObjectWriter(nil)
	obj.Open()
	obj.Close()

	data, err = obj.BuildBytes()
	if err != nil {
		t.Fatal(err)
	}
	assertValidJSON(t, data)
	if string(data) != "{}" {
		t.Errorf("Expected {}, got %s", string(data))
	}
}
