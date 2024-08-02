package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/myselfBZ/Blog/v2/elasticsearch"
	storeage "github.com/myselfBZ/Blog/v2/storage"
	"github.com/myselfBZ/Blog/v2/types"
	"github.com/olivere/elastic/v7"
)

//Common http errors
var (
    MethodNotAllowed = map[string]interface{}{
        "error":"method not allowed",
    }
    InvalidJSON =  map[string]interface{}{
        "error":"invalid json",
    }
    NotFound = map[string]interface{}{
        "error":"not found",
    }
    InternaleServerErr = map[string]string{
        "error":"server error",
    }
)


type Handler struct{
    Store       storeage.Store 
    Elastic     *elasticsearch.ElasticSearch
}

func NewHandler(Store storeage.Store, Elastic *elasticsearch.ElasticSearch) *Handler {
	return &Handler{
		Store: Store,
        Elastic: Elastic,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
    CheckMethod(w, r, http.MethodPost) 
    var user types.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        WriteJSONErr(w) 
        return 
    }
    err := h.Store.InsertUser(r.Context(),&user)
    if err != nil{
        json.NewEncoder(w).Encode(map[string]interface{}{"error":err})
        return 
    }    
    w.WriteHeader(http.StatusOK)
}

func(h *Handler) SearchBlog(w http.ResponseWriter, r *http.Request){
    CheckMethod(w, r, http.MethodGet)
    log.Println("Method is get") 
    var query types.SearchQuery
    if err := json.NewDecoder(r.Body).Decode(&query); err != nil{
        WriteJSONErr(w)
        return 
    }
    
    log.Println(query.Query) 
    results, err := h.Elastic.Search(r.Context(), query.Query)
    if err != nil{
        if elastic.IsNotFound(err){
            json.NewEncoder(w).Encode(NotFound)
        }
        log.Println(err)
        json.NewEncoder(w).Encode(InternaleServerErr)
        return 
    }
    log.Println(results)
    var ids []map[string]interface{}
    for _, h := range results{
        var id map[string]interface{}
        _ = json.Unmarshal(h.Source, &id)
        ids = append(ids, id)
    }
    json.NewEncoder(w).Encode(ids)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request){
    CheckMethod(w, r, http.MethodGet)
    id := r.PathValue("id")
    blog, err := h.Store.GetByID(r.Context(),id) 
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(NotFound)
        return  
    }
    json.NewEncoder(w).Encode(blog)
}

func (h *Handler) DeleteBlog(w http.ResponseWriter, r *http.Request){
    CheckMethod(w, r, http.MethodDelete)
    return  
}

func (h *Handler) CreateBlog(w http.ResponseWriter, r *http.Request){
    CheckMethod(w, r, http.MethodPost)
    var blog types.Blog
    if err := json.NewDecoder(r.Body).Decode(&blog); err != nil{
        WriteJSONErr(w)
        return 
    }
    log.Println(blog)
    err := h.Store.InsertBlog(r.Context(),&blog)  
    if err != nil{ 
        log.Print(err)
        json.NewEncoder(w).Encode(InternaleServerErr)
        return 
    }
    err = h.Elastic.AddIndex(r.Context(), blog.Title, blog.ID)
    if err != nil{
        log.Print(err)
    }
    w.WriteHeader(http.StatusOK)
}

// Error handling functions, I know this looks awful 
func WriteJSONErr(w http.ResponseWriter)  {
    json.NewEncoder(w).Encode(InvalidJSON) 
}

func CheckMethod(w http.ResponseWriter, r *http.Request, method string)  {
    if r.Method != method{
        json.NewEncoder(w).Encode(MethodNotAllowed)
            
    }
}
