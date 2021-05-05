build:
	echo "Start build..." && \
	go build && \
	echo "Build is finished!!"
	
test_all: test_s test_c

test_s:
	go clean -testcache && go test -cover ./test/services -coverpkg=./app/services
	
test_c:
	go clean -testcache && go test -cover ./test/controllers -coverpkg=./app/controllers