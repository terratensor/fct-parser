build:
	GOOS=darwin GOARCH=amd64 go build -o ./dist/fct-parser.darwin.amd64
	GOOS=darwin GOARCH=arm64 go build -o ./dist/fct-parser.darwin.arm64
	GOOS=windows GOARCH=amd64 go build -o ./dist/fct-parser64.exe
	GOOS=windows GOARCH=386 go build -o ./dist/fct-parser.exe
	GOOS=linux GOARCH=amd64 go build -o ./dist/fct-parser.linux.amd64

docker-build:
	docker --log-level=debug build --pull --file=docker/Dockerfile --tag=${REGISTRY}/fct-parser:${IMAGE_TAG} .

push:
	docker push ${REGISTRY}/fct-parser:${IMAGE_TAG}
