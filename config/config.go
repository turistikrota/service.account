package config

type MongoAccount struct {
	Host       string `env:"MONGO_ACCOUNT_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_ACCOUNT_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_ACCOUNT_USERNAME" envDefault:""`
	Password   string `env:"MONGO_ACCOUNT_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_ACCOUNT_DATABASE" envDefault:"account"`
	Collection string `env:"MONGO_ACCOUNT_COLLECTION" envDefault:"accounts"`
	Query      string `env:"MONGO_ACCOUNT_QUERY" envDefault:""`
}

type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Http struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"account"`
}

type Grpc struct {
	Port int `env:"GRPC_PORT" envDefault:"3001"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
	Pw   string `env:"REDIS_PASSWORD"`
	Db   int    `env:"REDIS_DB"`
}

type CacheRedis struct {
	Host string `env:"REDIS_CACHE_HOST"`
	Port string `env:"REDIS_CACHE_PORT"`
	Pw   string `env:"REDIS_CACHE_PASSWORD"`
	Db   int    `env:"REDIS_CACHE_DB"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type Session struct {
	Topic string `env:"SESSION_TOPIC"`
}

type Topics struct {
	Account AccountEvents
}

type AccountEvents struct {
	Deleted  string `env:"STREAMING_TOPIC_ACCOUNT_DELETED"`
	Created  string `env:"STREAMING_TOPIC_ACCOUNT_CREATED"`
	Updated  string `env:"STREAMING_TOPIC_ACCOUNT_UPDATED"`
	Disabled string `env:"STREAMING_TOPIC_ACCOUNT_DISABLED"`
	Restored string `env:"STREAMING_TOPIC_ACCOUNT_RESTORED"`
	Enabled  string `env:"STREAMING_TOPIC_ACCOUNT_ENABLED"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type App struct {
	Protocol string `env:"PROTOCOL" envDefault:"http"`
	DB       struct {
		Account MongoAccount
	}
	RSA         RSA
	Grpc        Grpc
	HttpHeaders HttpHeaders
	Http        Http
	Session     Session
	I18n        I18n
	Topics      Topics
	Nats        Nats
	Redis       Redis
	CacheRedis  CacheRedis
	TokenSrv    TokenSrv
}
