package responses

import (
	"net/http"

	proto "github.com/golang/protobuf/proto"
)

type protobuf struct {
	message proto.Message
	code    int
}

func (p protobuf) Write(w http.ResponseWriter) {
	data, err := proto.Marshal(p.message)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/vnd.google.protobuf")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(p.code)
	w.Write(data)
}

func Protobuf(message proto.Message, code int) Response {
	return protobuf{
		message: message,
		code:    code,
	}
}
