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

type EnumMenuType uint

const (
	EMT_MENU      EnumMenuType = 0
	EMT_DIRECTORY EnumMenuType = 1
)

type EnumSexType uint

const (
	EST_UNKNOWN EnumSexType = 0
	EST_MALE    EnumSexType = 1
	EST_FEMALE  EnumSexType = 2
)

// String
//
//	@Description: 转换为字符串
//	@receiver e EnumSexType
//	@return string
func (e EnumSexType) String() string {
	switch e {
	case EST_UNKNOWN:
		return "其他"
	case EST_MALE:
		return "男"
	case EST_FEMALE:
		return "女"
	default:
		return "男"
	}
}
