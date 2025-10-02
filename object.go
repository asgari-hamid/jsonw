package jsonw

import "github.com/mailru/easyjson/jwriter"

type ObjectWriter struct {
	writer     *jwriter.Writer
	needsComma bool
}

func (w *ObjectWriter) Open(writer *jwriter.Writer) {
	if writer == nil {
		w.writer = &jwriter.Writer{}
	} else {
		w.writer = writer
	}

	w.writer.RawByte(openBrace)
	w.needsComma = false
}

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

func (w *ObjectWriter) NullField(name string) {
	if w.needsComma {
		w.writer.RawByte(comma)
	}

	w.writer.RawByte(quote)
	w.writer.RawString(name)
	w.writer.Raw(quoteColonNull, nil)

	w.needsComma = true
}

func (w *ObjectWriter) Close() {
	w.writer.RawByte(closeBrace)
}

func (w *ObjectWriter) BuildBytes() ([]byte, error) {
	return w.writer.BuildBytes()
}
