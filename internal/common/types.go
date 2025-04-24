package common

type EnumApiMethod string

const (
	HTTP_GET     EnumApiMethod = "GET"
	HTTP_POST    EnumApiMethod = "POST"
	HTTP_PUT     EnumApiMethod = "PUT"
	HTTP_DELETE  EnumApiMethod = "DELETE"
	HTTP_OPTIONS EnumApiMethod = "OPTIONS"
	HTTP_HEAD    EnumApiMethod = "HEAD"
)

func (e EnumApiMethod) String() string {
	return string(e)
}

type EnumDictContentType int

const (
	CTT_BOOLEAN  EnumDictContentType = 0 // boolean
	CTT_INT      EnumDictContentType = 1 // int
	CTT_FLOAT    EnumDictContentType = 2 // float
	CTT_JSON     EnumDictContentType = 3 // json
	CTT_TEXT     EnumDictContentType = 4 // text
	CTT_TEXT_XML EnumDictContentType = 5 // xml
	CTT_YAML     EnumDictContentType = 6 // yaml
	CTT_TOML     EnumDictContentType = 7 // toml
)
