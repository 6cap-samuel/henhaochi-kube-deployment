package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-clean-arch/configurations"
	"golang-clean-arch/entities"
	"golang-clean-arch/exceptions"
	"golang-clean-arch/usecases/out"
)

type postRepository struct {
	Collection *mongo.Collection
}

func NewPostRepository(database *mongo.Database) out.PostDataSource {
	return &postRepository{
		Collection: database.Collection("posts"),
	}
}

func (p postRepository) GetAll() (response []entities.Post) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"date_created", -1}})

	cursor, err := p.Collection.Find(ctx, bson.M{}, findOptions)
	exceptions.PanicIfNeeded(err)

	for cursor.Next(ctx) {
		var post entities.Post

		err := cursor.Decode(&post)
		exceptions.PanicIfNeeded(err)

		response = append(response, post)
	}
	return response
}

func (p postRepository) Create(post entities.Post) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	_, err := p.Collection.InsertOne(ctx, post)

	exceptions.PanicIfNeeded(err)
}
