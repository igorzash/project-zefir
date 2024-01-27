IMAGE_NAME_API := api
IMAGE_NAME_MIGRATIONS := migrations
K8S_MANIFESTS_DIR_API := k8s/api
K8S_MANIFESTS_DIR_MIGRATIONS := k8s/migrations

.PHONY: dev delete-pods-api build-api load-to-kind-api apply-manifests-api migrate build-migrations load-to-kind-migrations apply-manifests-migrations

dev: build-api load-to-kind-api delete-pods-api apply-manifests-api

delete-pods-api:
	kubectl delete --all pods --namespace=api

build-api:
	docker build -f ./api/api.Dockerfile -t $(IMAGE_NAME_API) ./api

load-to-kind-api:
	kind load docker-image $(IMAGE_NAME_API)

apply-manifests-api:
	kubectl apply -f $(K8S_MANIFESTS_DIR_API)

migrate: build-migrations load-to-kind-migrations apply-manifests-migrations

build-migrations:
	docker build -f ./migrations/migrations.Dockerfile -t $(IMAGE_NAME_MIGRATIONS) ./migrations

load-to-kind-migrations:
	kind load docker-image $(IMAGE_NAME_MIGRATIONS)

apply-manifests-migrations:
	kubectl apply -f $(K8S_MANIFESTS_DIR_MIGRATIONS)
