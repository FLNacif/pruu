package models

//Request is a type to map all toml request files
type Request struct {
	Name     string
	Method   string
	Headers  map[string]string
	Body     Body
	Response Response
	Output   Output
}
