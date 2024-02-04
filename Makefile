DOCKER := docker
KUBECTL := kubectl
IMAGE_NAME_WEB := web
IMAGE_NAME_MIGRATIONS := migrations
DOCKERFILE_WEB := images/web.Dockerfile
DOCKERFILE_MIGRATIONS := images/migrations.Dockerfile
K8S_MANIFESTS_DIR_WEB := k8s/web
K8S_MANIFESTS_DIR_MIGRATIONS := k8s/web_migrations

.PHONY: dev test delete-pods-web build-web apply-manifests-web migrate delete-migration-job build-migrations apply-manifests-migrations help

# Build and deploy the web application
dev: build-web delete-pods-web apply-manifests-web

MIGRATIONS_URL=file://$(shell pwd)/web_migrations/migrations
SECRET_KEY=$$(dd if=/dev/urandom bs=100 count=1 status=none | base64)
TEST_GO_PARAMS=GIN_MODE=debug CGO_ENABLED=1 ALLOW_IN_MEMORY_DB=1 MIGRATIONS_URL=$(MIGRATIONS_URL) SECRET_KEY=$(SECRET_KEY)

# Run tests on the web application
test:
	@cd web && \
	$(TEST_GO_PARAMS) go test ./...

# Delete all pods in the 'web' namespace
delete-pods-web:
	@$(KUBECTL) delete --all pods --namespace=web

# Build the Docker image for the web application
build-web:
	@$(DOCKER) build -f $(DOCKERFILE_WEB) -t $(IMAGE_NAME_WEB) ./web

# Apply the Kubernetes manifests for the web application
apply-manifests-web:
	@$(KUBECTL) apply -k $(K8S_MANIFESTS_DIR_WEB)/overlays/dev

# Build the Docker image for the database migrations, delete any existing migration jobs, and apply the Kubernetes manifests for the database migrations
migrate: build-migrations delete-migration-job apply-manifests-migrations

# Delete the migration job in the 'api' namespace
delete-migration-job:
	@$(KUBECTL) delete -napi job.batch api-migration --ignore-not-found=true

# Build the Docker image for the database migrations
build-migrations:
	@$(DOCKER) build -f $(DOCKERFILE_MIGRATIONS) -t $(IMAGE_NAME_MIGRATIONS) ./web_migrations

# Apply the Kubernetes manifests for the database migrations
apply-manifests-migrations:
	@$(KUBECTL) apply -k $(K8S_MANIFESTS_DIR_MIGRATIONS)/base

# Print out a description of each target
help:
	@echo "dev: Build and deploy the web application"
	@echo "test: Run tests on the web application"
	@echo "delete-pods-web: Delete all pods in the 'web' namespace"
	@echo "build-web: Build the Docker image for the web application"
	@echo "apply-manifests-web: Apply the Kubernetes manifests for the web application"
	@echo "migrate: Build the Docker image for the database migrations, delete any existing migration jobs, and apply the Kubernetes manifests for the database migrations"
	@echo "delete-migration-job: Delete the migration job in the 'api' namespace"
	@echo "build-migrations: Build the Docker image for the database migrations"
	@echo "apply-manifests-migrations: Apply the Kubernetes manifests for the database migrations"
