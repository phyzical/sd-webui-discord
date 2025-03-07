build:
	docker build . -t phyzical/sd-webui-discord 
run:
	docker run --rm --env-file=.env --mount type=bind,source="./user_center.db",target=/usr/local/bin/user_center.db --mount type=bind,source="./config.json",target=/usr/local/bin/config.json phyzical/sd-webui-discord	
dev:
	make build
	make run