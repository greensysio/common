package firebase

import (
	cmContext "bitbucket.org/greensys-tech/common/context"
	"bitbucket.org/greensys-tech/common/log"
	locationModel "bitbucket.org/greensys-tech/common/model/location"
	"fmt"
)

// UpdateNode will update node by path
func UpdateNode(c *cmContext.CustomContext, path string, data map[string]interface{}) bool {
	ctx, cncl := cmContext.InitNewCtxFromCustomCtx(c)
	defer cncl()

	log.InfofCtx(c, "Update firebase node. Data: %+v", data)
	nodeRef := GetDBClient().NewRef(path)
	if err := nodeRef.Update(ctx, data); err != nil {
		log.ErrorfCtx(c, "Update node fail! Path: %s. Params: %+v. Error: %s", path, data, err.Error())
	}
	return true
}

// GetLastVehicleLocationOfOrder will return last location of vehicle of order
func GetLastVehicleLocationOfOrder(c *cmContext.CustomContext, odID string) (location *locationModel.Location, ok bool) {
	ctx, cncl := cmContext.InitNewCtxFromCustomCtx(c)
	defer cncl()

	path := fmt.Sprintf("location_history/%s", odID)
	nodeRef := GetDBClient().NewRef(path)
	nodes, err := nodeRef.OrderByKey().LimitToLast(1).GetOrdered(ctx)
	log.InfofCtx(c, "Get last location of driver on order id %s at Firebase. Path: %s", odID, path)
	if err != nil {
		log.ErrorfCtx(c, "GetLastVehicleLocationOfOrder fail! Path: %s. Error: %s", path, err.Error())
		return nil, false
	}

	if len(nodes) > 0 {
		l := locationModel.Location{}
		if err := nodes[0].Unmarshal(&l); err != nil {
			log.ErrorfCtx(c, "Error Unmarshal result: %+v . Response: %+v", err, nodes[0])
			return nil, false
		}
		return &l, true
	}
	return nil, true
}

// GetVehicleLocationOfOrderAsQuery will return vehicle locations of order as query.
func GetVehicleLocationOfOrderAsQuery(c *cmContext.CustomContext, odID string, query *DBQuery) (locations []*locationModel.Location, ok bool) {
	ctx, cncl := cmContext.InitNewCtxFromCustomCtx(c)
	defer cncl()

	path := fmt.Sprintf("location_history/%s", odID)
	nodeRef := GetDBClient().NewRef(path)
	q, err := buildQuery(nodeRef, query)
	if err != nil {
		log.ErrorCtx(c, "build query error:", err)
		return nil, false
	}
	nodes, err := q.GetOrdered(ctx)
	log.InfofCtx(c, "GetVehicleLocationOfOrderAsQuery: order id %s at Firebase. Path: %s", odID, path)
	if err != nil {
		log.ErrorfCtx(c, "GetVehicleLocationOfOrderAsQuery fail! Path: %s. Error: %s", path, err.Error())
		return nil, false
	}

	locations = make([]*locationModel.Location, len(nodes))
	for i, node := range nodes {
		l := locationModel.Location{}
		if err := node.Unmarshal(&l); err != nil {
			log.ErrorfCtx(c, "Error Unmarshal result: %+v . Response: %+v", err, nodes[0])
			return nil, false
		}
		locations[i] = &l
	}
	return locations, true
}
