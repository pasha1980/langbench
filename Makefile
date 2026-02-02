THREADS ?= 16
DURATION ?= 30
CONNS ?= 100
URL ?= http://127.0.0.1:8080

test: test-wrk

test-wrk:
	wrk -s test.lua -t${THREADS} -d${DURATION} -c${CONNS} ${URL}
