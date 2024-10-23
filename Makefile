cobra:
	cobra-cli add $(filter-out $@,$(MAKECMDGOALS))

run:
	go run main.go $(filter-out $@,$(MAKECMDGOALS))

normal:
	go run main.go --url=http://localhost:8080 --concurrency=5 --requests=150	

int: 
	go run main.go --url=http://localhost:8080/int --concurrency=5 --requests=150

notfound:
	go run main.go --url=http://localhost:8080/notfound --concurrency=5 --requests=150

random:
	go run main.go --url=http://localhost:8080/random --concurrency=5 --requests=150

.PHONY: cobra, run

%:
	@:
