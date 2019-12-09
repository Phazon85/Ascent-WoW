build-docker:
		docker build -t phazon85/ascent-wow:latest -f ./build/package/Dockerfile .

push-docker:
		docker push phazon85/ascent-wow:latest