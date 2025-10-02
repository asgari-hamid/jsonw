package jsonw

import "github.com/mailru/easyjson/jwriter"

type ArrayWriter struct {
	writer     *jwriter.Writer
	needsComma bool
}

func (w *ArrayWriter) Open(writer *jwriter.Writer) {
	if writer == nil {
		w.writer = &jwriter.Writer{}
	} else {
		w.writer = writer
	}

	w.writer.RawByte(openBracket)
	w.needsComma = false
}

func (w *ArrayWriter) ObjectValue() *ObjectWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	obj := &ObjectWriter{}
	obj.Open(w.writer)
	return obj
}

func (w *ArrayWriter) ArrayValue() *ArrayWriter {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	arr := &ArrayWriter{}
	arr.Open(w.writer)
	return arr
}

func (w *ArrayWriter) StringValue(value string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.String(value)

	w.needsComma = true
}

func (w *ArrayWriter) NumberValue(value string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawString(value)

	w.needsComma = true
}

func (w *ArrayWriter) IntegerValue(value int64) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Int64(value)

	w.needsComma = true
}

func (w *ArrayWriter) FloatValue(value float64) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Float64(value)

	w.needsComma = true
}

func (w *ArrayWriter) BoolValue(b bool) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Bool(b)

	w.needsComma = true
}

func (w *ArrayWriter) NullValue() {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.Raw(nullValue, nil)

	w.needsComma = true
}

func (w *ArrayWriter) Close() {
	w.writer.RawByte(closeBracket)
}

func (w *ArrayWriter) BuildBytes() ([]byte, error) {
	return w.writer.BuildBytes()
}
