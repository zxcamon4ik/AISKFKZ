myUID := $(shell id -u)


d-run:
	@export myUID=${myUID} &&\
	COMPOSE_PROFILES=full_dev \
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 \
		docker compose up \
			--build
d-purge:
	@export myUID=${myUID} &&\
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 \
		docker compose down --volumes --remove-orphans --rmi local --timeout 0 


.PHONY: d-run d-purge