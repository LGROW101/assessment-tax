test:
	 go test -covermode=atomic ./... -coverprofile=coverage.out 
testout: 
	go tool cover -html=coverage.out 
testv:
	 go test -v ./tests/service  
	 go test -v ./tests/repository 
	go test -v ./tests/handler 