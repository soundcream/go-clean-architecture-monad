package msg

import "fmt"

const (
	InvalidDataFormat    = "%s"
	MissingRequiredField = ""
)

func GetInvalidDataFormatSF(msg string) string {
	return fmt.Sprintf(InvalidDataFormat, msg)
}
