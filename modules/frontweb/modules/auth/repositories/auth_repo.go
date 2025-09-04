package auth_repo

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

var ErrExampleNotFound = errors.New("example not found")

func cloneAndInjectNotDeleted(filter interface{}) interface{} {
	m, ok := filter.(bson.M)
	if !ok {
		return filter
	}
	cp := make(bson.M, len(m)+1)
	for k, v := range m {
		cp[k] = v
	}
	if _, exists := cp["deleted_at"]; !exists {
		if _, hasOr := cp["$or"]; !hasOr {
			cp["$or"] = []bson.M{{"deleted_at": bson.M{"$exists": false}}, {"deleted_at": nil}}
		}
	}
	return cp
}
