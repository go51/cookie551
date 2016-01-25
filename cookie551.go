package cookie551

import "net/http"

type Cookie struct {
	w http.ResponseWriter
	r *http.Request
}

func New(w http.ResponseWriter, r *http.Request) *Cookie {
	cookie := Cookie{
		w: w,
		r: r,
	}

	return &cookie
}
