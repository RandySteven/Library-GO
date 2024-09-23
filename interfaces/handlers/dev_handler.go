package handlers_interfaces

import "net/http"

type DevHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
	CreateBucket(w http.ResponseWriter, r *http.Request)
	GetListBuckets(w http.ResponseWriter, r *http.Request)
}
