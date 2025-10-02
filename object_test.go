package jsonw_test

import (
	"encoding/json"
	"testing"

	"github.com/asgari-hamid/jsonw"
)

func TestObjectWriter(t *testing.T) {
	w := &jsonw.ObjectWriter{}
	w.Open(nil)

	// Primitive fields
	w.StringField("name", "Hamid")
	w.IntegerField("age", 30)
	w.FloatField("score", 95.5)
	w.BoolField("active", true)
	w.NullField("nickname")

	// Nested object
	obj := w.ObjectField("address")
	obj.StringField("city", "Tehran")
	obj.IntegerField("zip", 12345)
	obj.Close()

	// Nested array
	arr := w.ArrayField("tags")
	arr.StringValue("go")
	arr.StringValue("json")
	arr.Close()

	w.Close()

	bytes, err := w.BuildBytes()
	if err != nil {
		t.Fatal(err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if data["name"] != "Hamid" || data["age"].(float64) != 30 || data["score"].(float64) != 95.5 {
		t.Errorf("primitive fields mismatch")
	}
	if !data["active"].(bool) || data["nickname"] != nil {
		t.Errorf("bool/null fields mismatch")
	}
	address := data["address"].(map[string]interface{})
	if address["city"] != "Tehran" || address["zip"].(float64) != 12345 {
		t.Errorf("nested object mismatch")
	}
	tags := data["tags"].([]interface{})
	if tags[0] != "go" || tags[1] != "json" {
		t.Errorf("nested array mismatch")
	}
}

func TestStringEscaping(t *testing.T) {
	w := &jsonw.ObjectWriter{}
	w.Open(nil)

	special := "Line1\nLine2\t\"Quote\"\\Backslash\u2028\u2029"
	w.StringField("text", special)

	w.Close()
	bytes, err := w.BuildBytes()
	if err != nil {
		t.Fatal(err)
	}

	var data map[string]string
	if err := json.Unmarshal(bytes, &data); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if data["text"] != special {
		t.Errorf("escaped string mismatch: got %v", data["text"])
	}
}

func TestEmptyStructures(t *testing.T) {
	obj := &jsonw.ObjectWriter{}
	obj.Open(nil)
	obj.Close()

	arr := &jsonw.ArrayWriter{}
	arr.Open(nil)
	arr.Close()

	objBytes, _ := obj.BuildBytes()
	arrBytes, _ := arr.BuildBytes()

	if string(objBytes) != "{}" {
		t.Errorf("empty object mismatch: %s", objBytes)
	}
	if string(arrBytes) != "[]" {
		t.Errorf("empty array mismatch: %s", arrBytes)
	}
}
