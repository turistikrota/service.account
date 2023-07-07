package adapters

import (
	"github.com/turistikrota/service.account/src/adapters/memory"
	"github.com/turistikrota/service.account/src/adapters/mongo"
	"github.com/turistikrota/service.account/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)
