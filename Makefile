start:
	docker-compose up
a-1-up:
	curl localhost:8080/set_gauge_a/proc/1/to_value/2
a-1-down:
	curl localhost:8080/set_gauge_a/proc/1/to_value/0
a-2-up:
	curl localhost:8080/set_gauge_a/proc/2/to_value/2
a-2-down:
	curl localhost:8080/set_gauge_a/proc/2/to_value/0
b-up:
	curl localhost:8080/set_gauge_b/to_value/2
b-down:
	curl localhost:8080/set_gauge_b/to_value/0
