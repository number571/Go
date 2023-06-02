.PHONY: default build run 
default: build run 
build:
	docker build . --tag simple_service
run: 
	docker run -it --rm -v $$(pwd)/mounted:/mounted -p 8080:80 simple_service
clean:
	docker rm $$(docker ps -a -q)
	docker rmi -f $$(docker images -aq)
