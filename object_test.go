package jsonw_test

import (
	"testing"

	"github.com/asgari-hamid/jsonw"
)

func TestObjectWriter(t *testing.T) {
	w := &jsonw.ObjectWriter{}
	w.Open(nil)

	w.StringField("name", `Alice "Admin" \ Example`)
	w.IntegerField("age", 30)
	w.BoolField("active", true)
	w.NullField("extra")

	arr := w.ArrayField("tags")
	arr.StringValue("go")
	arr.StringValue("json")
	arr.Close()

	w.Close()

	got, err := w.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes error: %v", err)
	}

	want := `{"name":"Alice \"Admin\" \\ Example","age":30,"active":true,"extra":null,"tags":["go","json"]}`
	if string(got) != want {
		t.Errorf("unexpected JSON:\n got: %s\nwant: %s", got, want)
	}
}

func TestEscaping(t *testing.T) {
	obj := &jsonw.ObjectWriter{}
	obj.Open(nil)
	obj.StringField("text", "Line1\nLine2\t\u2028\u2029\"\\")
	obj.Close()

	got, err := obj.BuildBytes()
	if err != nil {
		t.Fatalf("BuildBytes error: %v", err)
	}

	want := `{"text":"Line1\nLine2\t\u2028\u2029\"\\"}`
	if string(got) == "" {
		t.Errorf("JSON is empty")
	}

	if string(got) != want {
		t.Errorf("unexpected JSON:\n got: %s\nwant: %s", got, want)
	}
}
