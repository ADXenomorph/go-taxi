package taxi_request

import (
	"strconv"
	"sync"

	"github.com/valyala/fastrand"
)

// RequestStorage stores taxi requests and their statistic counters
type RequestStorage struct {
	// data stores taxi requests
	data sync.Map
	// counters store taxi requests statistic counters
	counters sync.Map

	// open serves as a cached list of current open taxi requests
	sync.RWMutex
	open []string
}

// NewStorage returns new initialized taxi RequestStorage
func NewStorage() *RequestStorage {
	return &RequestStorage{}
}

// Save stores the taxi request and updates a list of open requests
func (rs *RequestStorage) Save(r *Request) {
	rs.data.Store(r.ID, r)
	rs.updateOpenList()
}

// Get returns the request by id, returns (nil, false) tuple if its not found
func (rs *RequestStorage) Get(requestId string) (*Request, bool) {
	req, ok := rs.data.Load(requestId)

	if !ok {
		return nil, false
	}

	return req.(*Request), true
}

// GetRandom takes a random id from list of open requests and returns it from storage
func (rs *RequestStorage) GetRandom() *Request {
	id := rs.getRandomId()

	if id == "" {
		return nil
	}

	req, ok := rs.Get(id)

	if !ok {
		return nil
	}

	return req
}

func (rs *RequestStorage) getRandomId() string {
	rs.RLock()
	defer rs.RUnlock()

	if len(rs.open) == 0 {
		return ""
	}

	// fastrand usage allowed to reduce operation time from 8000ns to 80ns
	id := rs.open[fastrand.Uint32n(uint32(len(rs.open)))]

	return id
}

// GetRandomAndCount returns a random request and increases its statistic counter
func (rs *RequestStorage) GetRandomAndCount() *Request {
	req := rs.GetRandom()

	if req != nil {
		rs.inc(req.ID)
	}

	return req
}

// GetCounters returns a slice of statistic lines for all requests
// Example: ["aa - 456", "bb - 123"]
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
	rs.Lock()
	defer rs.Unlock()

	open := make([]string, 0)
	rs.data.Range(func(key interface{}, val interface{}) bool {
		open = append(open, val.(*Request).ID)
		return true
	})

	rs.open = open
}
