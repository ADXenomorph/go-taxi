# go-taxi
Test task to implement a taxi service using Golang.

Imagine you're writing backend for a taxi service and your simplified REST API will
return current taxi requests or request analytics.

# Requirements
Implement a small service, that stores and handles taxi requests.
A taxi request is 2 latin letters (az, yu, br, qq etc.).
At the start of the application 50 random requests are generated.
Every 200 ms 1 random request is cancelled and 1 new appears.

REST API consists of 2 api calls:
1. /request, that returns a random request from the current open ones;
2. /admin/requests, that returns a list of all created and cancelled requests and statistics for each request of how many times it was returned. Zero returned requests can be skipped.

A synthetic [apache ab|https://en.wikipedia.org/wiki/ApacheBench] load test should be implemented to test the rps for implmented service and to check how many simultaneous taxi drivers the service can handle.

## Additional notes
This task is estimated to 1-2 hours. The real time spent wont affect the score, but it is expected, that you will be able to solve similar problem in this time in future.

The task is graded by following criterias:
1. Architecture. Your chosen file structure of the project, objects in code, their dependencies and inheritance.
2. Instuments and libraries. What 3rd party dependencies was chosen for the solution, build process etc.
3. Asynchrony. Do you understand pros and cons of the language you're using? Did you solve the problem of async request? Which method did you use to solve async problems (mutex, channels or other)
4. Performance. Do you know about [C10k|https://en.wikipedia.org/wiki/C10k_problem]? Do you understand what the results of your load testing mean?

Optional points that will increase the grade:
- Tests
- Function/variable naming
- Understanding of godoc and comments usage
- Github repo you've commited your solution to

Points that will decrease the grade:
- Using an existing database. (sqlite, postgres, boltdb etc)
- Complex architectural patterns (interfaces usage, reflection etc)
- [Race detectors|https://golang.org/doc/articles/race_detector.html]
- Missing gofmt formatting

## API Request examples
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
