package services

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func GetErrosMessage(err error) string {
	var concatenedErros string = ""
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			concatenedErros += GetErrosMessageWithField(fe.Field(), fe.Tag(), fe.Param()) + ". "
		}
	} else {
		concatenedErros = err.Error()
	}
	return concatenedErros
}

func GetErrosMessageWithField(field string, tag string, param string) string {
	var errorMessage string = ""

	switch tag {
	case "required":
		errorMessage = "O campo " + field + " é obrigatório"
	case "exists":
		errorMessage = "O campo " + field + " deve existir em " + param
	case "not_exists":
		errorMessage = "O campo " + field + " não deve existir em " + param
	case "exists_or_null":
		errorMessage = "O campo " + field + " deve existir em " + param + " ou ser nulo"
	case "oneof":
		errorMessage = "O campo " + field + " deve ser " + param
	case "time_format":
		errorMessage = "O campo " + field + " deve ser uma data válida"
	case "email":
		errorMessage = "O campo " + field + " deve ser um email válido"
	case "alpha":
		errorMessage = "O campo " + field + " deve conter apenas letras"
	case "alphanum":
		errorMessage = "O campo " + field + " deve conter apenas letras e números"
	case "alphanumunicode":
		errorMessage = "O campo " + field + " deve conter apenas letras e números"
	case "alphaunicode":
		errorMessage = "O campo " + field + " deve conter apenas letras"
	case "ascii":
		errorMessage = "O campo " + field + " deve conter apenas caracteres ASCII"
	case "boolean":
		errorMessage = "O campo " + field + " deve ser um booleano"
	case "contains":
		errorMessage = "O campo " + field + " deve conter " + param
	case "containsany":
		errorMessage = "O campo " + field + " deve conter " + param
	case "containsrune":
		errorMessage = "O campo " + field + " deve conter " + param
	case "endsnotwith":
		errorMessage = "O campo " + field + " não deve terminar com " + param
	case "endswith":
		errorMessage = "O campo " + field + " deve terminar com " + param
	case "excludes":
		errorMessage = "O campo " + field + " não deve conter " + param
	case "excludesall":
		errorMessage = "O campo " + field + " não deve conter " + param
	case "excludesrune":
		errorMessage = "O campo " + field + " não deve conter " + param
	case "lowercase":
		errorMessage = "O campo " + field + " deve ser minúsculo"
	case "multibyte":
		errorMessage = "O campo " + field + " deve conter caracteres multibyte"
	case "number":
		errorMessage = "O campo " + field + " deve ser um número"
	case "numeric":
		errorMessage = "O campo " + field + " deve ser numérico"
	case "printascii":
		errorMessage = "O campo " + field + " deve conter apenas caracteres ASCII imprimíveis"
	case "startsnotwith":
		errorMessage = "O campo " + field + " não deve começar com " + param
	case "startswith":
		errorMessage = "O campo " + field + " deve começar com " + param
	case "uppercase":
		errorMessage = "O campo " + field + " deve ser maiúsculo"
	case "base64":
		errorMessage = "O campo " + field + " deve ser uma string base64"
	case "datetime":
		errorMessage = "O campo " + field + " deve ser uma data válida"
	case "jwt":
		errorMessage = "O campo " + field + " deve ser um token JWT"
	case "latitude":
		errorMessage = "O campo " + field + " deve ser uma latitude válida"
	case "longitude":
		errorMessage = "O campo " + field + " deve ser uma longitude válida"
	case "rgb":
		errorMessage = "O campo " + field + " deve ser uma string RGB"
	case "rgba":
		errorMessage = "O campo " + field + " deve ser uma string RGBA"
	case "md5":
		errorMessage = "O campo " + field + " deve ser um hash MD5"
	case "dir":
		errorMessage = "O campo " + field + " deve ser um diretório existente"
	case "dirpath":
		errorMessage = "O campo " + field + " deve ser um caminho de diretório"
	case "file":
		errorMessage = "O campo " + field + " deve ser um arquivo existente"
	case "len":
		errorMessage = "O campo " + field + " deve ter " + param + " caracteres"
	case "max":
		errorMessage = "O campo " + field + " deve ser máximo " + param
	case "min":
		errorMessage = "O campo " + field + " deve ser mínimo " + param
	case "required_if":
		errorMessage = "O campo " + field + " é obrigatório se " + param
	case "required_unless":
		errorMessage = "O campo " + field + " é obrigatório a menos que " + param
	case "required_with":
		errorMessage = "O campo " + field + " é obrigatório com " + param
	case "required_with_all":
		errorMessage = "O campo " + field + " é obrigatório com todos " + param
	case "required_without":
		errorMessage = "O campo " + field + " é obrigatório sem " + param
	case "required_without_all":
		errorMessage = "O campo " + field + " é obrigatório sem todos " + param
	case "excluded_if":
		errorMessage = "O campo " + field + " é excluído se " + param
	case "excluded_unless":
		errorMessage = "O campo " + field + " é excluído a menos que " + param
	case "excluded_with":
		errorMessage = "O campo " + field + " é excluído com " + param
	case "excluded_with_all":
		errorMessage = "O campo " + field + " é excluído com todos " + param
	case "excluded_without":
		errorMessage = "O campo " + field + " é excluído sem " + param
	case "excluded_without_all":
		errorMessage = "O campo " + field + " é excluído sem todos " + param
	case "phone_number":
		errorMessage = "O campo " + field + " deve ser um número de telefone válido no formato +5511988776655 ou +551188776655"
	default:
		errorMessage = "Erro no campo " + field + " validação " + tag + " parâmetro " + param
	}

	return errorMessage
}
