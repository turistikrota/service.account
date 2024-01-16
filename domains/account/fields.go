package account

type fieldsType struct {
	UUID          string
	UserUUID      string
	UserName      string
	FullName      string
	Description   string
	Social        string
	IsActive      string
	CompletedRate string
	IsDeleted     string
	IsVerified    string
	BirthDate     string
	CreatedAt     string
	UpdatedAt     string
}

type socialFieldsType struct {
	Platform   string
	Value      string
	FixedValue string
}

var fields = fieldsType{
	UUID:          "_id",
	UserUUID:      "user_uuid",
	UserName:      "user_name",
	FullName:      "full_name",
	Description:   "description",
	Social:        "social",
	IsActive:      "is_active",
	CompletedRate: "completed_rate",
	IsDeleted:     "is_deleted",
	IsVerified:    "is_verified",
	BirthDate:     "birth_date",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var socialFields = socialFieldsType{
	Platform:   "platform",
	Value:      "value",
	FixedValue: "fixed_value",
}

func socialField(field string) string {
	return fields.Social + "." + field
}

func socialFieldInArray(field string) string {
	return fields.Social + ".$." + field
}
