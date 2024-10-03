package siem

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"restwaf/internal/config"
	"restwaf/internal/validator"
	"strings"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

type Siem struct {
	opensearch *opensearch.Client
	indexName  string
}

func CreateSiem() *Siem {
	return new(Siem)
}

func (siem *Siem) Initialization(configuration *config.OpenSearchConfiguration) error {
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: configuration.Insecureskipverify},
		},
		Addresses: configuration.Urls,
		Username:  configuration.Username, // For testing only. Don't store credentials in code.
		Password:  configuration.Password,
	})
	if err != nil {
		return err
	}
	siem.opensearch = client
	siem.indexName = configuration.Index
	response, error := opensearchapi.IndicesExistsRequest{
		Index: []string{configuration.Index},
	}.Do(context.Background(), client)
	if error != nil {
		return error
	}
	if response.StatusCode == 404 {
		settings := strings.NewReader(`{
			'settings': {
				'index': {
					'number_of_shards': 1,
					'number_of_replicas': 0
					}
				}
			}`)
		response, err = opensearchapi.IndicesCreateRequest{
			Index: configuration.Index,
			Body:  settings,
		}.Do(context.Background(), client)
		if err != nil {
			return err
		}
		println("%v", response)
	}
	return nil
}

func (siem *Siem) Publish(response *validator.ValidatorResponse) {
	jsonStr, _ := (json.Marshal(response))

	req := opensearchapi.IndexRequest{
		Index: siem.indexName,
		Body:  strings.NewReader(string(jsonStr)),
	}
	req.Do(context.Background(), siem.opensearch)

}
