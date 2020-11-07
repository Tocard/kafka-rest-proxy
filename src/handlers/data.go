package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//	"github.com/Shopify/sarama"
	//	log "github.com/sirupsen/logrus"
)

func SendData(r *http.Request) (int, string) {
	d, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var result interface{}
	json.Unmarshal(d, &result)
	m := result.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case []interface{}:
			for _, u := range vv {
				switch zz := u.(type) {
				case map[string]interface{}:
					for h, o := range zz {
						if h == "meta" {
							meta := make(map[string]interface{})
							switch l := o.(type) {
							case map[string]interface{}:
								for k, v := range l {
									meta[k] = v
								}
								meta["kafka_rest_proxy"] = "hostname"
							}
							o = meta
						}
					}
				default:
					fmt.Printf("Type handler not implemented: %#v\n", zz)
				}
			}
		default:
			return http.StatusInternalServerError, k + "is of a type I don't know how to handle"
		}
	}
	json.Marshal(m)
	return http.StatusAccepted, string("OK")
}
