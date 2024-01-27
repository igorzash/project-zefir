IMAGE_NAME_API := api
IMAGE_NAME_MIGRATIONS := migrations
K8S_MANIFESTS_DIR_API := k8s/api
K8S_MANIFESTS_DIR_MIGRATIONS := k8s/migrations

.PHONY: dev delete-pods-api build-api apply-manifests-api migrate build-migrations apply-manifests-migrations

dev: build-api delete-pods-api apply-manifests-api

test:
	cd api && \
	GIN_MODE=release CGO_ENABLED=1 SECRET_KEY=$$(dd if=/dev/urandom bs=100 count=1 status=none | base64) go test

delete-pods-api:
	kubectl delete --all pods --namespace=api

build-api:
	docker build -f ./api/api.Dockerfile -t $(IMAGE_NAME_API) ./api

#load-to-kind-api:
#	kind load docker-image $(IMAGE_NAME_API)

apply-manifests-api:
	kubectl apply -k $(K8S_MANIFESTS_DIR_API)/overlays/dev

migrate: build-migrations delete-migration-job apply-manifests-migrations

delete-migration-job:
	kubectl delete -napi job.batch api-migration --ignore-not-found=true

build-migrations:
	docker build -f ./migrations/migrations.Dockerfile -t $(IMAGE_NAME_MIGRATIONS) ./migrations

#load-to-kind-migrations:
#	kind load docker-image $(IMAGE_NAME_MIGRATIONS)

apply-manifests-migrations:
	kubectl apply -k $(K8S_MANIFESTS_DIR_MIGRATIONS)/base
