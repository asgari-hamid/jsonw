package jsonw_test

import (
	"testing"

	"github.com/asgari-hamid/jsonw"
)

func TestArrayWriter(t *testing.T) {
	arr := &jsonw.ArrayWriter{}
	arr.Open(nil)

	arr.StringValue("hello")
	arr.NumberValue("123.45")
	arr.IntegerValue(100)
	arr.FloatValue(1.23)
	arr.BoolValue(false)
	arr.NullValue()

	obj := arr.ObjectValue()
	obj.StringField("key", "value")
	obj.Close()

	arr.Close()

	got, err := arr.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes error: %v", err)
	}

	want := `["hello",123.45,100,1.23,false,null,{"key":"value"}]`
	if string(got) != want {
		t.Errorf("unexpected JSON:\n got: %s\nwant: %s", got, want)
	}
}
