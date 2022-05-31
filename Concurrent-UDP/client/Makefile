build:
	GOOS=linux GOARCH=amd64 go build -o udpclient .
	chmod 777 udpclient

imgbuild:
	docker build -t lnanjangud653/udpclient:"${IMG_VER}" .

imgpush:
	docker push lnanjangud653/udpclient:${IMG_VER}
