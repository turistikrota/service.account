package mongo

import (
	mongo_account "github.com/turistikrota/service.account/src/adapters/mongo/account"
	mongo_platform "github.com/turistikrota/service.account/src/adapters/mongo/platform"
	"github.com/turistikrota/service.account/src/domain/account"
	"github.com/turistikrota/service.account/src/domain/platform"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewAccount(accountFactory account.Factory, collection *mongo.Collection) account.Repository
	NewPlatform(platformFactory platform.Factory, collection *mongo.Collection) platform.Repository
}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}

func (m *mongodb) NewAccount(accountFactory account.Factory, collection *mongo.Collection) account.Repository {
	return mongo_account.New(accountFactory, collection)
}

func (m *mongodb) NewPlatform(platformFactory platform.Factory, collection *mongo.Collection) platform.Repository {
	return mongo_platform.New(platformFactory, collection)
}
