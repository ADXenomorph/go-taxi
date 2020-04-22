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

	openMutex sync.RWMutex
	open []string
}

func NewStorage() *RequestStorage {
	return &RequestStorage{}
}

func (rs *RequestStorage) Save(r *Request) {
	rs.data.Store(r.ID, r)
	rs.updateOpenList()
}

func (rs *RequestStorage) Get(requestId string) (*Request, bool) {
	req, ok := rs.data.Load(requestId)

	if !ok {
		return nil, false
	}

	return req.(*Request), true
}

func (rs *RequestStorage) GetRandom() *Request {
	rs.openMutex.RLock()
	defer rs.openMutex.RUnlock()

	if len(rs.open) == 0 {
		return nil
	}

	rand.Seed(time.Now().Unix())
	id := rs.open[rand.Intn(len(rs.open))]

	req, ok := rs.Get(id)

	if !ok {
		return nil
	}

	return req
}

func (rs *RequestStorage) GetRandomAndCount() *Request {
	req := rs.GetRandom()
	rs.inc(req.ID)
	return req
}

func (rs *RequestStorage) GetCounters() []string {
	res := make([]string, 0)
	rs.counters.Range(func(key interface{}, val interface{}) bool {
		res = append(res, key.(string)+" - "+strconv.Itoa(val.(int)))
		return true
	})

	return res
}

func (rs *RequestStorage) inc(requestId string) {
	val, ok := rs.counters.Load(requestId)
	if !ok {
		val = 0
	}

	rs.counters.Store(requestId, val.(int)+1)
}

func (rs *RequestStorage) updateOpenList() {
	rs.openMutex.Lock()
	defer rs.openMutex.Unlock()

	open := make([]string, 0)
	rs.data.Range(func(key interface{}, val interface{}) bool {
		open = append(open, val.(*Request).ID)
		return true
	})

	rs.open = open
}
