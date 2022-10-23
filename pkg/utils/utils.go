package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func GetIdConverted(id string) (int64, bool) {
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return 0, false
	}

	return ID, true
}
