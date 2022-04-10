package subscription

import (
	"context"
	"polunzh/my-feed/dal"
	"polunzh/my-feed/model"
)

var client = dal.InitClient("subscription")

func Add(ctx context.Context, entity model.Subscription) (*model.Subscription, error) {
	data, err := client.Subscription.Create().SetName(entity.Name).SetURL(entity.URL).Save(ctx)
	if err != nil {
		return nil, err
	}

	return model.ToModel(data), nil
}
