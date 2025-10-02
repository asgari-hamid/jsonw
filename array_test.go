package jsonw_test

import (
	"encoding/json"
	"testing"

	"github.com/asgari-hamid/jsonw"
)

func TestArrayWriter(t *testing.T) {
	w := &jsonw.ArrayWriter{}
	w.Open(nil)

	w.StringValue("first")
	w.IntegerValue(2)
	w.FloatValue(3.14)
	w.BoolValue(true)
	w.NullValue()

	// Nested array
	arr := w.ArrayValue()
	arr.StringValue("nested")
	arr.IntegerValue(99)
	arr.Close()

	// Nested object
	obj := w.ObjectValue()
	obj.StringField("key", "value")
	obj.Close()

	w.Close()

	bytes, err := w.BuildBytes()
	if err != nil {
		t.Fatal(err)
	}

	var data []interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if data[0] != "first" || data[1].(float64) != 2 || data[2].(float64) != 3.14 || !data[3].(bool) || data[4] != nil {
		t.Errorf("primitive array values mismatch")
	}

	nestedArr := data[5].([]interface{})
	if nestedArr[0] != "nested" || nestedArr[1].(float64) != 99 {
		t.Errorf("nested array mismatch")
	}

	nestedObj := data[6].(map[string]interface{})
	if nestedObj["key"] != "value" {
		t.Errorf("nested object mismatch")
	}
}
