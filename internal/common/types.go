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
