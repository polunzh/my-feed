package feed

import (
	"context"
	"polunzh/my-feed/dal"
	"polunzh/my-feed/model"
)

var client = dal.InitClient("feed")

func Add(ctx context.Context, entity model.Feed) (*model.Feed, error) {
	data, err := client.Feed.Create().SetName(entity.Name).SetURL(entity.URL).Save(ctx)
	if err != nil {
		return nil, err
	}

	return model.ToModel(data), nil
}
