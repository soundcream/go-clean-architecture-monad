package global

type (
	Config struct {
		App          AppConfig
		DbConfig     DbConfig
		ReadDbConfig DbConfig
		Public       []string
	}
	AppConfig struct {
		AppName  string
		Domain   string
		HttpPort int
	}
	DbConfig struct {
		Host     string
		Port     int
		DbName   string
		Username string
		Password string
	}
	InvalidateField struct {
		Error       bool
		FailedField string
		Tag         string
		Msg         string
		Value       interface{}
	}
	ErrorHandlerResp struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Ext     interface{} `json:"ext"`
	}
	KeyValue[T, S any] struct {
		Key T
		Val S
	}
	PagingModel[T any] struct {
		Data  []T `json:"data"`
		Total int `json:"total"`
	}
)

func NewKeyValue[T, S any](key T, value S) KeyValue[T, S] {
	return KeyValue[T, S]{Key: key, Val: value}
}
