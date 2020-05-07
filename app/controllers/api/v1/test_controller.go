package v1

import (
	"encoding/json"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	dataMap := map[string]interface{}{
		"msg": "Hello, World",
	}

	data, _ := json.Marshal(dataMap)

	w.Write(data)
}
