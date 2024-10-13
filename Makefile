cobra:
	cobra-cli add $(filter-out $@,$(MAKECMDGOALS))

run:
	go run main.go $(filter-out $@,$(MAKECMDGOALS))

.PHONY: cobra, run

%:
	@:
