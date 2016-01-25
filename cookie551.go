package cookie551

import (
	"net/http"
	"time"
)

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

func (c *Cookie) Set(name, value string, second time.Duration) {

	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		Raw:      value,
	}

	if second == 0 {
		cookie.MaxAge = 0
	} else {
		expire := time.Now().Add(second * time.Second)
		cookie.Expires = expire
	}

	http.SetCookie(c.w, &cookie)

}

func (c *Cookie) Get(name string) (string, error) {

	cookie, err := c.r.Cookie(name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (c *Cookie) Delete(name string) {

	cookie := http.Cookie{
		Name:   name,
		MaxAge: -1,
	}

	http.SetCookie(c.w, &cookie)
}
