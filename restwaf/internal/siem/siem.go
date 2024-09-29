package siem

import (
	"context"
	"crypto/tls"
	"net/http"
	"restwaf/internal/config"
	"strings"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

type Siem struct {
	opensearch *opensearch.Client
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
