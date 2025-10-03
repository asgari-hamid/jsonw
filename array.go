package jsonw

import "github.com/mailru/easyjson/jwriter"

// ArrayWriter provides a low-level JSON array builder using easyjson's jwriter.Writer.
// It supports appending values of different types, including nested arrays and objects.
type ArrayWriter struct {
	writer     *jwriter.Writer
	needsComma bool
}

// NewArrayWriter constructs a new ArrayWriter with optionally a writer from parent.
func NewArrayWriter(writer *jwriter.Writer) *ArrayWriter {
	if writer == nil {
		writer = &jwriter.Writer{}
	}
	return &ArrayWriter{
		writer:     writer,
		needsComma: false,
	}
}

// Open writes the opening '[' for the array.
func (w *ArrayWriter) Open() {
	w.writer.RawByte(openBracket)
	w.needsComma = false
}

// ObjectValue appends a new object to the array and returns an ObjectWriter for it.
func (w *ArrayWriter) ObjectValue() *ObjectWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.needsComma = true

	return NewObjectWriter(w.writer)
}

// ArrayValue appends a new array to the array and returns an ArrayWriter for it.
func (w *ArrayWriter) ArrayValue() *ArrayWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.needsComma = true

	return NewArrayWriter(w.writer)
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

// BooleanValue appends a boolean value to the array.
func (w *ArrayWriter) BooleanValue(value bool) {
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
