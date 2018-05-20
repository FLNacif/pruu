package parser

import "github.com/PMoneda/pruu/core/models"

//RequestParser is a interface to parse some file format to struct data
type RequestParser interface {
	Parse(fileContent string) (*models.Request, error)
}
