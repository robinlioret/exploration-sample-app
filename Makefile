.PHONY: build
build:
	pack-cli build cr.sandbox.local/sample-app --builder gcr.io/buildpacks/builder:google-22

.PHONY: publish
publish: build
	docker push cr.sandbox.local/sample-app:latest

.PHONY: deploy
deploy: publish
	kubectl apply --server-side -f deploy.yaml

.PHONY: undeploy
undeploy:
	kubectl delete -f deploy.yaml