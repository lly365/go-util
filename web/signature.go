package web

import (
	"fmt"
	"sort"
	"strconv"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"io"
)

type Signature struct {
	Signature string
	Timestamp int
	Nonce string
	Echostr string
	Token string
}

func (s Signature) String() string {
	return fmt.Sprintf("Signature: %s, Timestamp: %d, Nonce: %s, Echostr: %s, Token: %s", s.Signature, s.Timestamp, s.Nonce, s.Echostr, s.Token);
}

func NewSignature(signature string, timestamp int, nonce string, echostr string, token string) *Signature {
	s := &Signature{signature, timestamp, nonce, echostr, token};
	return s
}

func (s *Signature) MakeSignature() string{
	params:=sort.StringSlice{s.Token, strconv.Itoa(s.Timestamp), s.Nonce}
	params.Sort()

	buf := make([]byte, len(s.Token) + len(strconv.Itoa(s.Timestamp)) + len(s.Nonce))
	buf = append(buf, params[0]...)
	buf = append(buf, params[1]...)
	buf = append(buf, params[2]...)

	hashSum := sha1.Sum(buf)
	return hex.EncodeToString(hashSum[:])
}

func (s *Signature) HttpHandle(w http.ResponseWriter, r *http.Request){
	if s.MakeSignature() == s.Signature {
		io.WriteString(w, s.Echostr)
	} else {
		io.WriteString(w, "INVALIDE SIGNATURE")
	}
}