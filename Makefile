start:
	docker-compose up
u:
	curl localhost:2112/modify/2
d:
	curl localhost:2112/modify/0
c:
	curl localhost:2112/modify/$v
