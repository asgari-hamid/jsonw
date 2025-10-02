package jsonw

import "github.com/mailru/easyjson/jwriter"

// ObjectWriter provides a low-level JSON object builder using easyjson's jwriter.Writer.
// It supports writing fields of different types, including nested objects and arrays.
type ObjectWriter struct {
	writer     *jwriter.Writer
	needsComma bool
}

// Open initializes the ObjectWriter with a given jwriter.Writer.
// If writer is nil, a new writer will be created. It also writes the opening '{'.
func (w *ObjectWriter) Open(writer *jwriter.Writer) {
	if writer == nil {
		w.writer = &jwriter.Writer{}
	} else {
		w.writer = writer
	}

	w.writer.RawByte(openBrace)
	w.needsComma = false
}

// ObjectField starts a new nested object field with the given name and returns an ObjectWriter.
func (w *ObjectWriter) ObjectField(name string) *ObjectWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)

	obj := &ObjectWriter{}
	obj.Open(w.writer)
	return obj
}

// ArrayField starts a new nested array field with the given name and returns an ArrayWriter.
func (w *ObjectWriter) ArrayField(name string) *ArrayWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)

	arr := &ArrayWriter{}
	arr.Open(w.writer)
	return arr
}

// StringField writes a string field to the object.
func (w *ObjectWriter) StringField(name, value string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)
	w.writer.String(value)

	w.needsComma = true
}

// NumberField writes a raw number (as string) field to the object.
func (w *ObjectWriter) NumberField(name, value string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)
	w.writer.RawString(value)

	w.needsComma = true
}

// IntegerField writes an int64 field to the object.
func (w *ObjectWriter) IntegerField(name string, value int64) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)
	w.writer.Int64(value)

	w.needsComma = true
}

// FloatField writes a float64 field to the object.
func (w *ObjectWriter) FloatField(name string, value float64) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)
	w.writer.Float64(value)

	w.needsComma = true
}

// BoolField writes a boolean field to the object.
func (w *ObjectWriter) BoolField(name string, b bool) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColon, nil)

	w.writer.Bool(b)

	w.needsComma = true
}

// NullField writes a JSON null field to the object.
func (w *ObjectWriter) NullField(name string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColonNull, nil)

	w.needsComma = true
}

// Close writes the closing '}' for the object.
func (w *ObjectWriter) Close() {
	w.writer.RawByte(closeBrace)

	w.needsComma = false
}

// BuildBytes returns the JSON bytes written by the writer.
func (w *ObjectWriter) BuildBytes() ([]byte, error) {
	return w.writer.BuildBytes()
}
