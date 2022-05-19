start:
	docker-compose up
a-up:
	curl localhost:8080/set_gauge_a/proc/$(proc)/to_value/2
a-down:
	curl localhost:8080/set_gauge_a/proc/$(proc)/to_value/0
b-up:
	curl localhost:8080/set_gauge_b/to_value/2
b-down:
	curl localhost:8080/set_gauge_b/to_value/0
