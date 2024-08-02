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


func (d *MongoStore) InsertBlog(ctx *context.Context, blog *types.Blog) error {
    _, err := d.db.Collection("blogs").InsertOne(*ctx, blog)
	if err != nil {
		return err
	} 
    return err
}

func (d *MongoStore) GetByID(ctx *context.Context, id string) (*types.Blog, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		res      = d.db.Collection("blogs").FindOne(*ctx, bson.M{"_id": objID})
		p        = &types.Blog{}
		err      = res.Decode(p)
	)
	return p, err
}

func (d *MongoStore) DeleteBlog(ctx *context.Context, id string )error  {
    objID, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id":objID}
    _, err := d.db.Collection("blogs").DeleteOne(*ctx, filter)
    return err
}

