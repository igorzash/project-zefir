IMAGE_NAME_WEB := web
IMAGE_NAME_MIGRATIONS := migrations
K8S_MANIFESTS_DIR_WEB := k8s/web
K8S_MANIFESTS_DIR_MIGRATIONS := k8s/web_migrations

.PHONY: dev test delete-pods-web build-web apply-manifests-web migrate delete-migration-job build-migrations apply-manifests-migrations

dev: build-web delete-pods-web apply-manifests-web

test:
	cd web && \
	GIN_MODE=debug CGO_ENABLED=1 SECRET_KEY=$$(dd if=/dev/urandom bs=100 count=1 status=none | base64) go test ./...

delete-pods-web:
	kubectl delete --all pods --namespace=web

build-web:
	docker build -f ./images/web.Dockerfile -t $(IMAGE_NAME_WEB) ./web

#load-to-kind-api:
#	kind load docker-image $(IMAGE_NAME_API)

apply-manifests-web:
	kubectl apply -k $(K8S_MANIFESTS_DIR_WEB)/overlays/dev

migrate: build-migrations delete-migration-job apply-manifests-migrations

delete-migration-job:
	kubectl delete -napi job.batch api-migration --ignore-not-found=true

build-migrations:
	docker build -f ./images/migrations.Dockerfile -t $(IMAGE_NAME_MIGRATIONS) ./web_migrations

#load-to-kind-migrations:
#	kind load docker-image $(IMAGE_NAME_MIGRATIONS)

apply-manifests-migrations:
	kubectl apply -k $(K8S_MANIFESTS_DIR_MIGRATIONS)/base
