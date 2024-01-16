package account

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserUnique struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type Repo interface {
	Create(ctx context.Context, account *Entity) (*Entity, *i18np.Error)
	ProfileView(ctx context.Context, name string) (*Entity, *i18np.Error)
	Get(ctx context.Context, u UserUnique) (*Entity, *i18np.Error)
	Exist(ctx context.Context, u UserUnique) (bool, *i18np.Error)
	Update(ctx context.Context, u UserUnique, account *Entity) *i18np.Error
	Disable(ctx context.Context, u UserUnique) *i18np.Error
	Enable(ctx context.Context, u UserUnique) *i18np.Error
	Delete(ctx context.Context, name string) *i18np.Error
	Restore(ctx context.Context, name string) *i18np.Error
	GetByName(ctx context.Context, name string) (*Entity, *i18np.Error)
	ListMy(ctx context.Context, userUUID string) ([]*Entity, *i18np.Error)
	ListByUniques(ctx context.Context, ids []UserUnique) ([]*Entity, *i18np.Error)
	ListByUserId(ctx context.Context, userUUID string) ([]*Entity, *i18np.Error)
	Filter(ctx context.Context, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
}

type repo struct {
	factory    Factory
	collection *mongo.Collection
	helper     mongo2.Helper[*Entity, *Entity]
}

func NewRepo(collection *mongo.Collection, factory Factory) Repo {
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[*Entity, *Entity](collection, createEntity),
	}
}

func createEntity() **Entity {
	return new(*Entity)
}

func (r *repo) Create(ctx context.Context, e *Entity) (*Entity, *i18np.Error) {
	res, err := r.collection.InsertOne(ctx, e)
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	e.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return e, nil
}

func (r *repo) ProfileView(ctx context.Context, name string) (*Entity, *i18np.Error) {
	filter := bson.M{
		fields.UserName: name,
		fields.IsActive: true,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	opts := options.FindOne().SetProjection(bson.M{
		fields.UserName:    1,
		fields.FullName:    1,
		fields.Description: 1,
		fields.IsVerified:  1,
		fields.CreatedAt:   1,
	})
	o, exist, err := r.helper.GetFilter(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *o, nil
}

func (r *repo) Get(ctx context.Context, u UserUnique) (*Entity, *i18np.Error) {
	filter := bson.M{
		fields.UserUUID: u.UUID,
		fields.UserName: u.Name,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *o, nil
}

func (r *repo) GetByName(ctx context.Context, name string) (*Entity, *i18np.Error) {
	filter := bson.M{
		fields.UserName: name,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *o, nil
}

func (r *repo) Exist(ctx context.Context, u UserUnique) (bool, *i18np.Error) {
	filter := bson.M{
		fields.UserUUID: u.UUID,
		fields.UserName: u.Name,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	return o != nil, nil
}

func (r *repo) Update(ctx context.Context, u UserUnique, account *Entity) *i18np.Error {
	filter := bson.M{
		fields.UserUUID: u.UUID,
		fields.UserName: u.Name,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.UserName:      account.UserName,
			fields.FullName:      account.FullName,
			fields.Description:   account.Description,
			fields.BirthDate:     account.BirthDate,
			fields.CompletedRate: account.CompletedRate,
			fields.UpdatedAt:     account.UpdatedAt,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Disable(ctx context.Context, u UserUnique) *i18np.Error {
	filter := bson.M{
		fields.UserUUID: u.UUID,
		fields.UserName: u.Name,
		fields.IsActive: true,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsActive:  false,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Enable(ctx context.Context, u UserUnique) *i18np.Error {
	filter := bson.M{
		fields.UserUUID: u.UUID,
		fields.UserName: u.Name,
		fields.IsActive: false,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsActive:  true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, name string) *i18np.Error {
	filter := bson.M{
		fields.UserName: name,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsDeleted: true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Restore(ctx context.Context, name string) *i18np.Error {
	filter := bson.M{
		fields.UserName:  name,
		fields.IsDeleted: true,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsDeleted: false,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) ListMy(ctx context.Context, userUUID string) ([]*Entity, *i18np.Error) {
	filter := bson.M{
		fields.UserUUID: userUUID,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	return r.helper.GetListFilter(ctx, filter)
}

func (r *repo) ListByUniques(ctx context.Context, users []UserUnique) ([]*Entity, *i18np.Error) {
	ids := make([]string, len(users))
	names := make([]string, len(users))
	for i, u := range users {
		ids[i] = u.UUID
		names[i] = u.Name
	}
	filter := bson.M{
		fields.UserUUID: bson.M{
			"$in": ids,
		},
		fields.UserName: bson.M{
			"$in": names,
		},
	}
	return r.helper.GetListFilter(ctx, filter)
}

func (r *repo) ListByUserId(ctx context.Context, userUUID string) ([]*Entity, *i18np.Error) {
	filter := bson.M{
		fields.UserUUID: userUUID,
	}
	return r.helper.GetListFilter(ctx, filter)
}

func (r *repo) Filter(ctx context.Context, filter FilterEntity, listConf list.Config) (*list.Result[*Entity], *i18np.Error) {
	filters := r.filterToBson(filter)
	l, err := r.helper.GetListFilter(ctx, filters, r.listOptions(listConf))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filters)
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*Entity]{
		IsNext:        filtered > listConf.Offset+listConf.Limit,
		IsPrev:        listConf.Offset > 0,
		FilteredTotal: filtered,
		Total:         filtered,
		Page:          listConf.Offset/listConf.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) listOptions(listConfig list.Config) *options.FindOptions {
	return options.Find().SetLimit(listConfig.Limit).SetSkip(listConfig.Offset).SetSort(bson.M{
		fields.CreatedAt: -1,
	})
}
