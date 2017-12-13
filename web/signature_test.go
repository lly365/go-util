package web

import (
	"testing"
	//"net/http"
)

func TestSignature(t *testing.T) {
	serverSig := new(Signature)
	serverSig.Token = "abc123"
	serverSig.Timestamp =1513150241
	serverSig.Nonce = "ivix"
	serverSig.Echostr = "ivix.me"

	clientSig := new(Signature)
	clientSig.Signature = serverSig.MakeSignature()
	clientSig.Token = "abc123"
	clientSig.Timestamp = serverSig.Timestamp
	clientSig.Nonce = serverSig.Nonce
	clientSig.Echostr = serverSig.Echostr

	if clientSig.Signature != clientSig.MakeSignature() {
		t.Fatal("invalide signature")
	}

	//http.HandleFunc("/sign", clientSig.HttpHandle)
	//http.ListenAndServe(":12345", nil)
}