build-docker:
		docker build -t dgparker/ascent-wow:latest -f ./build/package/Dockerfile .

push-docker:
		docker push dgparker/ascent-wow:latest