package jsonw

var (
	openBrace      = byte('{')
	closeBrace     = byte('}')
	openBracket    = byte('[')
	closeBracket   = byte(']')
	quote          = byte('"')
	comma          = byte(',')
	quoteColon     = []byte(`":`)
	nullValue      = []byte("null")
	quoteColonNull = []byte(`":null`)
)
