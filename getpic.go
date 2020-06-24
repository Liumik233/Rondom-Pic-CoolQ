package main

import (
	"io/ioutil"
	"net/http"
)

func getpic() []byte {
	pic,_:=http.Get("https://pic.atori.top/imgs")
	picbody,_:=ioutil.ReadAll(pic.Body)
	return picbody
}
