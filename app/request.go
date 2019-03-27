package app

import (
	"crypto/sha256"
	"encoding/base64"
	"time"

	"io/ioutil"

	"net/http/httputil"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//Dump request structure
type Dump struct {
	CreatedAt time.Time `json:"created_at"`
	Tag       string    `json:"tag"`
	Checksum  string    `json:"body_checksum"`
	UID       string    `json:"id"`
	Opened    bool      `json:"opened"`
	Value     string    `json:"value"`
	Method    string    `json:"method"`
	BodySize  int64     `json:"body_size"`
	URI       string    `json:"uri"`
}

//NewDump from request
func NewDump(c *gin.Context) Dump {
	full, _ := httputil.DumpRequest(c.Request, true)
	b := ioutil.NopCloser(c.Request.Body)
	body, _ := ioutil.ReadAll(b)

	dump := Dump{
		CreatedAt: time.Now(),
		Tag:       c.Param("tag"),
		Opened:    false,
		Checksum:  Sha256(body),
		UID:       uuid.NewV4().String(),
		Value:     string(full),
		Method:    c.Request.Method,
		BodySize:  c.Request.ContentLength,
		URI:       c.Request.RequestURI,
	}
	return dump
}

func Sha256(s []byte) string {
	h := sha256.New()
	h.Write(s)
	sEnc := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return sEnc
}
