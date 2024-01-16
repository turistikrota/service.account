package platform

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	GetBySlug(ctx context.Context, slug string) (*Entity, *i18np.Error)
	Create(ctx context.Context, platform *Entity) *i18np.Error
	Update(ctx context.Context, slug string, platform *Entity) *i18np.Error
	Disable(ctx context.Context, slug string) *i18np.Error
	Delete(ctx context.Context, slug string) *i18np.Error
	Enable(ctx context.Context, slug string) *i18np.Error
	ListAll(ctx context.Context) ([]*Entity, *i18np.Error)
	TranslationCreate(ctx context.Context, platform string, locale Locale, translations Translations) *i18np.Error
	TranslationUpdate(ctx context.Context, platform string, locale Locale, translations Translations) *i18np.Error
	TranslationDelete(ctx context.Context, platform string, locale Locale) *i18np.Error
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

func (r *repo) Create(ctx context.Context, p *Entity) *i18np.Error {
	_, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		return r.factory.Errors.Failed("create")
	}
	return nil
}

func (r *repo) GetBySlug(ctx context.Context, slug string) (*Entity, *i18np.Error) {
	filter := bson.M{fields.Slug: slug}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound("platform")
	}
	return *o, nil
}

func (r *repo) Update(ctx context.Context, slug string, platform *Entity) *i18np.Error {
	filter := bson.M{fields.Slug: slug}
	setter := bson.M{"$set": bson.M{
		fields.Name:      platform.Name,
		fields.Slug:      platform.Slug,
		fields.Regexp:    platform.Regexp,
		fields.Prefix:    platform.Prefix,
		fields.UpdatedAt: platform.UpdatedAt,
	}}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Disable(ctx context.Context, slug string) *i18np.Error {
	filter := bson.M{fields.Slug: slug}
	setter := bson.M{"$set": bson.M{
		fields.IsActive:  false,
		fields.UpdatedAt: time.Now(),
	}}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Enable(ctx context.Context, slug string) *i18np.Error {
	filter := bson.M{fields.Slug: slug}
	setter := bson.M{"$set": bson.M{
		fields.IsActive:  true,
		fields.UpdatedAt: time.Now(),
	}}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, slug string) *i18np.Error {
	filter := bson.M{fields.Slug: slug}
	setter := bson.M{"$set": bson.M{
		fields.IsDeleted: true,
		fields.UpdatedAt: time.Now(),
	}}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) ListAll(ctx context.Context) ([]*Entity, *i18np.Error) {
	filter := bson.M{fields.IsDeleted: bson.M{
		"$ne": true,
	}, fields.IsActive: true}
	return r.helper.GetListFilter(ctx, filter)
}

func (r *repo) TranslationCreate(ctx context.Context, platform string, locale Locale, translations Translations) *i18np.Error {
	filter := bson.M{
		fields.Slug: platform,
	}
	setter := bson.M{
		"$set": bson.M{
			translationField(locale.String()): bson.M{
				translationFields.Name:        translations.Name,
				translationFields.Description: translations.Description,
				translationFields.Placeholder: translations.Placeholder,
			},
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) TranslationUpdate(ctx context.Context, platform string, locale Locale, translations Translations) *i18np.Error {
	filter := bson.M{
		fields.Slug: platform,
	}
	setter := bson.M{
		"$set": bson.M{
			translationField(locale.String()): bson.M{
				translationFields.Name:        translations.Name,
				translationFields.Description: translations.Description,
				translationFields.Placeholder: translations.Placeholder,
			},
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) TranslationDelete(ctx context.Context, platform string, locale Locale) *i18np.Error {
	filter := bson.M{fields.Slug: platform}
	setter := bson.M{
		"$unset": bson.M{
			translationField(locale.String()): "",
		},
		"$set": bson.M{
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}
