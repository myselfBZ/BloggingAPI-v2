package storeage

import (
	"context"
	"github.com/myselfBZ/Blog/v2/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type MongoStore struct{
    db      *mongo.Database
}

type Store interface{
    InsertBlog(context.Context, *types.Blog) error 
    DeleteBlog(context.Context, string) error
    GetByID (context.Context, string) (*types.Blog, error)
    InsertUser (context.Context, *types.User) error 
}


func NewMongoStore(db *mongo.Database) *MongoStore {
	return &MongoStore{
		db:   db,
	}
}




func (d *MongoStore) InsertBlog(ctx context.Context, blog *types.Blog) error {
    r, err := d.db.Collection("blogs").InsertOne(ctx, *blog)
	if err != nil {
		return err
	} 
    blog.ID = r.InsertedID.(primitive.ObjectID).Hex()
    return err
}

func (d *MongoStore) GetByID(ctx context.Context, id string) (*types.Blog, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		res      = d.db.Collection("blogs").FindOne(ctx, bson.M{"_id": objID})
		p        = &types.Blog{}
		err      = res.Decode(p)
	)
	return p, err
}

func (d *MongoStore) DeleteBlog(ctx context.Context, id string )error  {
    objID, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id":objID}
    _, err := d.db.Collection("blogs").DeleteOne(ctx, filter)
    return err
}


func (d *MongoStore) InsertUser(ctx context.Context, user *types.User) error {
    _, err := d.db.Collection("users").InsertOne(ctx, user) 
    if err != nil {
        return err 
    }
    return err
}




