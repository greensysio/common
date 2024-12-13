package firebase

import (
	"errors"

	"firebase.google.com/go/v4/db"
)

// OrderBy enum for query firebase db, order nodes by
type OrderBy int

const (
	// UnknownOrderBy is default status
	UnknownOrderBy OrderBy = iota // null
	OrderByChild
	OrderByValue
	OrderByKey
)

type QueryDBFunc func(query *db.Query) *db.Query

type DBQuery struct {
	OrderBy OrderBy
	// If OrderBy = OrderByuChild
	OrderByChild string
	QueryFunc    QueryDBFunc
}

func buildQuery(dbRef *db.Ref, query *DBQuery) (*db.Query, error) {
	if query == nil {
		return dbRef.OrderByKey(), nil
	}
	var q *db.Query
	var err error
	switch query.OrderBy {
	case OrderByKey:
		q = dbRef.OrderByKey()
	case OrderByValue:
		q = dbRef.OrderByValue()
	case OrderByChild:
		if len(query.OrderByChild) > 0 {
			q = dbRef.OrderByChild(query.OrderByChild)
		} else {
			err = errors.New("missing child for order by child")
		}
	default:
		q = dbRef.OrderByKey()
	}
	if err != nil {
		return nil, err
	}
	return query.QueryFunc(q), nil
}
