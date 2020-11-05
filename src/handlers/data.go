package handlers

import (
	"data"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendData(r *http.Request) (int, string) {
	d := json.NewDecoder(r.Body)
	metric := data.NewMetric()
	d.Decode(metric)

	fmt.Println(d)
	fmt.Println(metric)
	c, _ := json.Marshal(metric)
	return http.StatusAccepted, string(c)
}
