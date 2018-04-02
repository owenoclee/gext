package responses

import "net/http"

type Response interface {
	Write(http.ResponseWriter)
}
