envoy:
	envoy -c ./envoy.yaml -l debug

run:
	go run cmd/cloudscale/*.go controller

.PHONY: ui
ui:
	cd ui; PUBLIC_URL=/ui npm run build

operator:
	go run cmd/cloudscale/*.go operator
