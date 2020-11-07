package handlers

import (
	"data"
	"encoding/json"
	"github.com/go-martini/martini"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Shopify/sarama"
)

func SendData(r *http.Request) (int, string) {
	d, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var result interface{}
	json.Unmarshal(d, &result)
	m := result.(map[string]interface{})

	for _, v := range m {
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
				}
			}
		}
	}
	json.Marshal(m)
	return http.StatusAccepted, string("OK")
}

func SendToKafka(params martini.Params, r *http.Request) (int, string) {
	defer r.Body.Close()
	broker := data.YamlConfig.KafkaBroker
	producer := data.NewProducer(broker)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithField("error", err.Error()).Error("failed to read body")
	}
	topic := params["topic"]
	topic = strings.ReplaceAll(topic, ":", "") // need to understand why
	log.Info(topic)
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(body),
	})
	if err != nil {
		log.WithField("error", err.Error()).Error("failed to stored")
	}

	return http.StatusAccepted, string("message delivered\n")
}
