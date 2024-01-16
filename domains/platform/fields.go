package platform

type fieldsType struct {
	UUID         string
	Name         string
	Slug         string
	Regexp       string
	Prefix       string
	Translations string
	IsActive     string
	IsDeleted    string
	UpdatedAt    string
	CreatedAt    string
}

type translationFieldType struct {
	Name        string
	Placeholder string
	Description string
}

var fields = fieldsType{
	UUID:         "_id",
	Name:         "name",
	Slug:         "slug",
	Regexp:       "regexp",
	Prefix:       "prefix",
	Translations: "translations",
	IsActive:     "is_active",
	IsDeleted:    "is_deleted",
	UpdatedAt:    "updated_at",
	CreatedAt:    "created_at",
}

var translationFields = translationFieldType{
	Name:        "name",
	Placeholder: "placeholder",
	Description: "description",
}

func translationFieldInArray(field string) string {
	return fields.Translations + ".$." + field
}

func translationField(field string) string {
	return fields.Translations + "." + field
}
