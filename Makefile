envoy:
	envoy -c ./envoy.yaml -l debug

run:
	go run cmd/cloudscale/*.go controller
