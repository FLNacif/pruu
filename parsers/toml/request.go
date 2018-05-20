package toml

import (
	"github.com/BurntSushi/toml"
	"github.com/PMoneda/pruu/core/models"
)

//Parser parse toml file into go struct data
type Parser struct {
}

//Parse interprets toml file into model data structure
func (p *Parser) Parse(content string) (*models.Request, error) {
	request := new(models.Request)
	if _, err := toml.Decode(content, request); err != nil {
		return nil, err
	}
	return request, nil
}
