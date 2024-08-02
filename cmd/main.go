package main

import (
	"context"
	"log"
	"net/http"

	"github.com/myselfBZ/Blog/v2/api"
	"github.com/myselfBZ/Blog/v2/elasticsearch"
	storeage "github.com/myselfBZ/Blog/v2/storage"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
    connMongo, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    FailOnErr(err)
    mongoDB := storeage.NewMongoStore(connMongo.Database("blogging"))
    connElcs, err  := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
    FailOnErr(err)
    els := elasticsearch.ElasticSearch{
        Client: connElcs, 
    }
    h := api.NewHandler(mongoDB, &els) 
    mux := http.NewServeMux()
    mux.HandleFunc("GET /blogs/{id}", h.GetById)
    mux.HandleFunc("POST /users", h.CreateUser)
    mux.HandleFunc("POST /blogs", h.CreateBlog)
    mux.HandleFunc("GET /search-blog", h.SearchBlog)
    log.Println("Listening...")
    http.ListenAndServe(":8080", mux)
}

func FailOnErr(err error, msg ...string)  {
    if err != nil{
        log.Fatal(err, msg)
    } 
}
