package global

// Config
type (
	Config struct {
		App          AppConfig
		DbConfig     DbConfig
		ReadDbConfig DbConfig
		Http         HttpConfig
		Public       []string
		Service      IntegrationService
	}
	AppConfig struct {
		AppName  string
		Domain   string
		HttpPort int
	}
	HttpConfig struct {
	}
	DbConfig struct {
		Host     string
		Port     int
		DbName   string
		Username string
		Password string
	}
	IntegrationService struct {
		PgwUrl   string
		QrPayUrl string
	}
)

// Error
type (
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
)

// Common
type (
	KeyValue[T, S any] struct {
		Key T
		Val S
	}
	PagingRequest struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
	PagingModel[T any] struct {
		Data   []T `json:"data"`
		Total  int `json:"total"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Page   int `json:"page"`
	}
)

func NewKeyValue[T, S any](key T, value S) KeyValue[T, S] {
	return KeyValue[T, S]{Key: key, Val: value}
}

func NewPagingModel[T any](data []T, total int, limit int, offset int) PagingModel[T] {
	page := 1
	if total > 0 && limit > 0 && offset > total {
		page = (offset / (total / limit)) + 1
	}
	return PagingModel[T]{
		Data:   data,
		Total:  total,
		Limit:  limit,
		Offset: offset,
		Page:   page,
	}
}
