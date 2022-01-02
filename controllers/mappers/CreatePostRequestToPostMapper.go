package mappers

import (
	"golang-clean-arch/configurations"
	"golang-clean-arch/controllers/requests"
	"golang-clean-arch/entities"
	"time"
)

func CreatePostRequestToPostMapper(request requests.CreatePostRequest) (*entities.Post, []entities.Hashtag) {
	var hashtags []entities.Hashtag

	for _, hashtag := range request.HashTags {
		hashtags = append(hashtags, entities.Hashtag{
			Id:    configurations.NewIdentity(),
			Name:  hashtag,
			Count: 1,
		})
	}
	return &entities.Post{
		Id:          configurations.NewIdentity(),
		Description: request.Description,
		Positives:   request.Positives,
		Neutrals:    request.Neutrals,
		Negatives:   request.Negatives,
		Store: entities.Store{
			Id:       configurations.NewIdentity(),
			Name:     request.StoreName,
			Image:    request.StoreImage,
			Location: request.StoreLocation,
			Lat:      request.StoreLat,
			Long:     request.StoreLong,
		},
		Rating:      request.Rating,
		HashTags:    request.HashTags,
		DateCreated: time.Now(),
	}, hashtags
}
