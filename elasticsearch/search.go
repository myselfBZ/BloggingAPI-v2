package elasticsearch

import (
	"context"

	"github.com/olivere/elastic/v7"
)


type ElasticSearch struct{
    Client *elastic.Client
}


func(e *ElasticSearch) Search(ctx context.Context, title string) ([]*elastic.SearchHit, error) {
    r, err := e.Client.Search().Index("myblogs").Do(ctx)
    if err != nil{
        return nil, err
    }
    return r.Hits.Hits, nil
}

func(e *ElasticSearch) AddIndex(ctx context.Context, title string, id string) error {
    doc := map[string]string{
        "id":id,
    }
    _, err := e.Client.Index().Index("myblogs").Id(title).BodyJson(doc).Do(ctx)
    return err
}

