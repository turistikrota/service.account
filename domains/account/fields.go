package account

type fieldsType struct {
	UUID          string
	UserUUID      string
	UserName      string
	FullName      string
	Description   string
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
	IsActive:      "is_active",
	CompletedRate: "completed_rate",
	IsDeleted:     "is_deleted",
	IsVerified:    "is_verified",
	BirthDate:     "birth_date",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}
