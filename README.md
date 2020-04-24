# go-taxi
Test task to implement a taxi service using Golang.

Imagine you're writing backend for a taxi service and your simplified REST API will
return current taxi requests or request analytics.

- [Requirements](#requirements)
  - [Additional notes](#additional-notes)
  - [API Request examples](#api-request-examples)
- [Usage](#usage)
- [Benchmarks](#benchmarks)
  - [Golang benchmarks](#golang-benchmarks)
  - [Apache bench](#apache-bench)
  - [Wrk](#wrk)
- [Test coverage](#test-coverage)

## Requirements
Implement a small service, that stores and handles taxi requests.
A taxi request is 2 latin letters (az, yu, br, qq etc.).
At the start of the application 50 random requests are generated.
Every 200 ms 1 random request is cancelled and 1 new appears.

REST API consists of 2 api calls:
1. /request, that returns a random request from the current open ones;
2. /admin/requests, that returns a list of all created and cancelled requests and statistics for each request of how many times it was returned. Zero returned requests can be skipped.

A synthetic [apache ab](https://en.wikipedia.org/wiki/ApacheBench) load test should be implemented to test the rps for implmented service and to check how many simultaneous taxi drivers the service can handle.

### Additional notes
This task is estimated to 1-2 hours. The real time spent wont affect the score, but it is expected, that you will be able to solve similar problem in this time in future.

The task is graded by following criterias:
1. Architecture. Your chosen file structure of the project, objects in code, their dependencies and inheritance.
2. Instuments and libraries. What 3rd party dependencies was chosen for the solution, build process etc.
3. Asynchrony. Do you understand pros and cons of the language you're using? Did you solve the problem of async request? Which method did you use to solve async problems (mutex, channels or other)
4. Performance. Do you know about [C10k](https://en.wikipedia.org/wiki/C10k_problem)? Do you understand what the results of your load testing mean?

Optional points that will increase the grade:
- Tests
- Function/variable naming
- Understanding of godoc and comments usage
- Github repo you've commited your solution to

Points that will decrease the grade:
- Using an existing database. (sqlite, postgres, boltdb etc)
- Complex architectural patterns (interfaces usage, reflection etc)
- [Race detectors](https://golang.org/doc/articles/race_detector.html)
- Missing gofmt formatting

### API Request examples
```
    [GET] /request
    az
    
    [GET] /request
    br
    
    [GET] /request
    br
    
    [GET] /admin/requests
    az - 1
    br - 2
```

## Usage

Run `make` to get a list of all make targets

Build:
```
make build
```

Run:
```
./taxid
```

Test:
```
make test
```

Show test coverage:
```
make cover
make cover-html
```

Benchmarks:
```
make bench
make bench-race
make apache-bench
make wrk
```

Open godoc:
```
make godoc
```
Dont forget to kill godoc process after it

## Benchmarks

All benchmarks were run on Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz with 16 GB RAM

### Golang benchmarks

```
go test -benchmem -bench=. -cpu=1,2,3,4 ./...
```

Initial implementation of main request for comparison:
```
BenchmarkGetRandomRequest     	  112736	      9785 ns/op	    1056 B/op	      10 allocs/op
BenchmarkGetRandomRequest-2   	  123648	      9660 ns/op	    1056 B/op	      10 allocs/op
BenchmarkGetRandomRequest-3   	  119479	     10081 ns/op	    1056 B/op	      10 allocs/op
BenchmarkGetRandomRequest-4   	   96632	     10541 ns/op	    1056 B/op	      10 allocs/op
```

Main request that returns random taxi request:
```
BenchmarkGetRandomRequest     	 4112386	       289 ns/op	      40 B/op	       3 allocs/op
BenchmarkGetRandomRequest-2   	 4248294	       268 ns/op	      40 B/op	       3 allocs/op
BenchmarkGetRandomRequest-3   	 4282032	       270 ns/op	      40 B/op	       3 allocs/op
BenchmarkGetRandomRequest-4   	 4326327	       274 ns/op	      40 B/op	       3 allocs/op
```

Random ID generation:
```
BenchmarkGenerateRequestId     	14458981	        79.5 ns/op	       5 B/op	       1 allocs/op
BenchmarkGenerateRequestId-2   	16287727	        76.0 ns/op	       5 B/op	       1 allocs/op
BenchmarkGenerateRequestId-3   	15980731	        75.9 ns/op	       5 B/op	       1 allocs/op
BenchmarkGenerateRequestId-4   	16114341	        77.2 ns/op	       5 B/op	       1 allocs/op
```

Getting 1 known taxi request:
```
BenchmarkGet                   	38049488	        30.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet-2                 	36363507	        32.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet-3                 	39277065	        30.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGet-4                 	37253802	        32.6 ns/op	       0 B/op	       0 allocs/op
```

Getting 1 random taxi request:
```
BenchmarkGetRandom             	15146282	        77.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetRandom-2           	15262102	        77.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetRandom-3           	15453110	        81.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetRandom-4           	14706447	        79.8 ns/op	       0 B/op	       0 allocs/op
```

Getting 1 random taxi request and increasing it's count
I think counters can be quicker if I use a different map with atomic counters
```
BenchmarkGetRandomAndCount     	 4378852	       264 ns/op	      40 B/op	       3 allocs/op
BenchmarkGetRandomAndCount-2   	 4495218	       268 ns/op	      40 B/op	       3 allocs/op
BenchmarkGetRandomAndCount-3   	 4258623	       270 ns/op	      40 B/op	       3 allocs/op
BenchmarkGetRandomAndCount-4   	 4556630	       273 ns/op	      40 B/op	       3 allocs/op
```

### Apache bench

Initial implementation: 
```
ab -n 30000 -c 500 localhost:8080/request

Concurrency Level:      500
Time taken for tests:   14.477 seconds
Complete requests:      30000
Failed requests:        0
Total transferred:      3540000 bytes
HTML transferred:       60000 bytes
Requests per second:    2072.30 [#/sec] (mean)
Time per request:       241.278 [ms] (mean)
Time per request:       0.483 [ms] (mean, across all concurrent requests)
Transfer rate:          238.80 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   19 146.1      0    3051
Processing:     2  114 973.5     25   13459
Waiting:        2  114 973.5     25   13459
Total:         10  133 1070.2     25   14470

Percentage of the requests served within a certain time (ms)
  50%     25
  66%     26
  75%     26
  80%     27
  90%     28
  95%     29
  98%     34
  99%   1468
 100%  14470 (longest request)
```


Latest: 
```
ab -n 50000 -c 1000 localhost:8080/request

Concurrency Level:      1000
Time taken for tests:   2.343 seconds
Complete requests:      50000
Failed requests:        0
Total transferred:      7750000 bytes
HTML transferred:       100000 bytes
Requests per second:    21343.78 [#/sec] (mean)
Time per request:       46.852 [ms] (mean)
Time per request:       0.047 [ms] (mean, across all concurrent requests)
Transfer rate:          3230.75 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   38 183.1      3    1049
Processing:     1    6  12.4      4     219
Waiting:        1    4  12.0      2     213
Total:          5   44 190.0      8    1228

Percentage of the requests served within a certain time (ms)
  50%      8
  66%     10
  75%     11
  80%     11
  90%     12
  95%     29
  98%   1053
  99%   1064
 100%   1228 (longest request)
```

### Wrk

```
wrk -t 4 -c 16 -d 10 http://localhost:8080/request
```

Initial implementation: 
```
Running 10s test @ http://localhost:8080/request
  4 threads and 16 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.62ms    1.56ms  18.90ms   87.40%
    Req/Sec     3.04k   333.44     6.17k    83.33%
  121429 requests in 10.10s, 13.66MB read
Requests/sec:  12022.63
Transfer/sec:      1.35MB

```

```
Running 10s test @ http://localhost:8080/request
  4 threads and 16 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   282.54us    0.93ms  20.52ms   94.14%
    Req/Sec    60.51k     6.14k   77.90k    71.75%
  2408609 requests in 10.01s, 312.40MB read
Requests/sec: 240613.28
Transfer/sec:     31.21MB

```

## Test coverage

```
github.com/ADXenomorph/go-taxi/cmd/taxid/taxid.go:16:				main			0.0%
github.com/ADXenomorph/go-taxi/cmd/taxid/taxid.go:40:				CreateRouter		66.7%
github.com/ADXenomorph/go-taxi/internal/taxi/taxi.go:25:			NewApp			100.0%
github.com/ADXenomorph/go-taxi/internal/taxi/taxi.go:31:			CreateRequest		100.0%
github.com/ADXenomorph/go-taxi/internal/taxi/taxi.go:40:			CancelRequest		100.0%
github.com/ADXenomorph/go-taxi/internal/taxi/taxi.go:55:			GetRandomRequest	100.0%
github.com/ADXenomorph/go-taxi/internal/taxi/taxi.go:62:			GetRequestStatistics	100.0%
github.com/ADXenomorph/go-taxi/internal/taxi/taxi.go:68:			CreateRandomRequest	100.0%
github.com/ADXenomorph/go-taxi/internal/taxi/taxi.go:74:			CreateInitialRequests	100.0%
github.com/ADXenomorph/go-taxi/internal/taxi/taxi.go:81:			CancelRandomRequest	100.0%
github.com/ADXenomorph/go-taxi/internal/taxi/taxi.go:93:			SimulateChanges		0.0%
github.com/ADXenomorph/go-taxi/internal/taxi_request/id_generator.go:10:	GenerateRequestId	100.0%
github.com/ADXenomorph/go-taxi/internal/taxi_request/request.go:18:		NewRequest		100.0%
github.com/ADXenomorph/go-taxi/internal/taxi_request/storage.go:23:		NewStorage		100.0%
github.com/ADXenomorph/go-taxi/internal/taxi_request/storage.go:28:		Save			100.0%
github.com/ADXenomorph/go-taxi/internal/taxi_request/storage.go:34:		Get			100.0%
github.com/ADXenomorph/go-taxi/internal/taxi_request/storage.go:45:		GetRandom		85.7%
github.com/ADXenomorph/go-taxi/internal/taxi_request/storage.go:61:		getRandomId		100.0%
github.com/ADXenomorph/go-taxi/internal/taxi_request/storage.go:76:		GetRandomAndCount	100.0%
github.com/ADXenomorph/go-taxi/internal/taxi_request/storage.go:88:		GetCounters		100.0%
github.com/ADXenomorph/go-taxi/internal/taxi_request/storage.go:98:		inc			100.0%
github.com/ADXenomorph/go-taxi/internal/taxi_request/storage.go:107:		updateOpenList		100.0%
total:										(statements)		80.7%
```