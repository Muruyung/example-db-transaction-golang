package connection_elastic

import (
	"os"

	"github.com/elastic/go-elasticsearch/v8"

	"coba/pkg/dotenv"
)

type Elastic struct {
	*elasticsearch.TypedClient
	index string
}

func Init() (*Elastic, error) {
	cert, err := os.ReadFile(dotenv.ELASTICPATHCRT())
	if err != nil {
		return nil, err
	}

	config := elasticsearch.Config{
		Addresses: []string{
			dotenv.ELASTICADDRESS(),
		},
		Username: dotenv.ELASTICUSERNAME(),
		Password: dotenv.ELASTICPASSWORD(),
		CACert:   cert,
	}

	client, err := elasticsearch.NewTypedClient(config)

	return &Elastic{
		TypedClient: client,
	}, err
}

func (es *Elastic) SetIndex(index string) {
	es.Indices.Create(index)
	es.index = index
}

func (es *Elastic) GetIndex() string {
	return es.index
}
