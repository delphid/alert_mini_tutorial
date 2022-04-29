start:
	docker-compose up
ua1:
	curl localhost:2112/modify_a/1/2
da1:
	curl localhost:2112/modify_a/1/0
ua2:
	curl localhost:2112/modify_a/2/2
da2:
	curl localhost:2112/modify_a/2/0
