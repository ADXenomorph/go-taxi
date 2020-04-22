package taxi_request

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type RequestStorage struct {
	data     sync.Map
	counters sync.Map
}

func NewStorage() *RequestStorage {
	return &RequestStorage{}
}

func (rs *RequestStorage) Save(r *Request) {
	rs.data.Store(r.ID, r)
}

func (rs *RequestStorage) Get(requestId string) (*Request, bool) {
	req, ok := rs.data.Load(requestId)

	if !ok {
		return nil, false
	}

	return req.(*Request), true
}

func (rs *RequestStorage) GetRandom() *Request {
	open := rs.All(Open)

	if len(open) == 0 {
		return nil
	}

	rand.Seed(time.Now().Unix())
	req := open[rand.Intn(len(open))]

	return req
}

func (rs *RequestStorage) GetRandomAndCount() *Request {
	req := rs.GetRandom()
	rs.inc(req.ID)
	return req
}

func (rs *RequestStorage) inc(requestId string) {
	val, ok := rs.counters.Load(requestId)
	if !ok {
		val = 0
	}

	rs.counters.Store(requestId, val.(int)+1)
}

func (rs *RequestStorage) All(status Status) []*Request {
	res := make([]*Request, 0)
	rs.data.Range(func(key interface{}, val interface{}) bool {
		res = append(res, val.(*Request))
		return true
	})

	return res
}

func (rs *RequestStorage) GetCounters() []string {
	res := make([]string, 0)
	rs.counters.Range(func(key interface{}, val interface{}) bool {
		res = append(res, key.(string)+" - "+strconv.Itoa(val.(int)))
		return true
	})

	return res
}
