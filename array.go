package jsonw

import "github.com/mailru/easyjson/jwriter"

// ArrayWriter provides a low-level JSON array builder using easyjson's jwriter.Writer.
// It supports appending values of different types, including nested arrays and objects.
type ArrayWriter struct {
	writer     *jwriter.Writer
	needsComma bool
}

// Open initializes the ArrayWriter with a given jwriter.Writer.
// If writer is nil, a new writer will be created. It also writes the opening '['.
func (w *ArrayWriter) Open(writer *jwriter.Writer) {
	if writer == nil {
		w.writer = &jwriter.Writer{}
	} else {
		w.writer = writer
	}

	w.writer.RawByte(openBracket)
	w.needsComma = false
}

// ObjectValue appends a new object to the array and returns an ObjectWriter for it.
func (w *ArrayWriter) ObjectValue() *ObjectWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	obj := &ObjectWriter{}
	obj.Open(w.writer)
	return obj
}

// ArrayValue appends a new array to the array and returns an ArrayWriter for it.
func (w *ArrayWriter) ArrayValue() *ArrayWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	arr := &ArrayWriter{}
	arr.Open(w.writer)
	return arr
}

// StringValue appends a string value to the array.
func (w *ArrayWriter) StringValue(value string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.String(value)

	w.needsComma = true
}

// NumberValue appends a raw number (as string) to the array.
func (w *ArrayWriter) NumberValue(value string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawString(value)

	w.needsComma = true
}

// IntegerValue appends an int64 value to the array.
func (w *ArrayWriter) IntegerValue(value int64) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Int64(value)

	w.needsComma = true
}

// FloatValue appends a float64 value to the array.
func (w *ArrayWriter) FloatValue(value float64) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Float64(value)

	w.needsComma = true
}

// BoolValue appends a boolean value to the array.
func (w *ArrayWriter) BoolValue(value bool) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Bool(value)

	w.needsComma = true
}

// NullValue appends a JSON null to the array.
func (w *ArrayWriter) NullValue() {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Raw(nullValue, nil)

	w.needsComma = true
}

// Close writes the closing ']' for the array.
func (w *ArrayWriter) Close() {
	w.writer.RawByte(closeBracket)

	w.needsComma = false
}

// BuildBytes returns the JSON bytes written by the writer.
func (w *ArrayWriter) BuildBytes() ([]byte, error) {
	return w.writer.BuildBytes()
}
