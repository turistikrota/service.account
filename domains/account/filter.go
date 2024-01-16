package account

import "go.mongodb.org/mongo-driver/bson"

type FilterEntity struct {
	Query string `query:"q,omitempty" validate:"omitempty,max=1000"`
}

func (r *repo) filterToBson(filter FilterEntity) bson.M {
	list := make([]bson.M, 0)
	list = r.filterByQuery(list, filter)
	listLen := len(list)
	if listLen == 0 {
		return bson.M{}
	}
	if listLen == 1 {
		return list[0]
	}
	return bson.M{
		"$and": list,
	}
}

func (r *repo) filterByQuery(list []bson.M, filter FilterEntity) []bson.M {
	if filter.Query != "" {
		list = append(list, bson.M{
			"$or": []bson.M{
				{
					fields.FullName: bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					fields.Description: bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					fields.UserName: bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					fields.UUID: bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					fields.UserUUID: bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
			},
		})
	}
	return list
}
