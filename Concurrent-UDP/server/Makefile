build:
	GOOS=linux GOARCH=amd64 go build -o udpserver .
	chmod 777 udpserver

imgbuild:
	docker build -t lnanjangud653/udpserver:"${IMG_VER}" .

imgpush:
	docker push lnanjangud653/udpserver:${IMG_VER}

