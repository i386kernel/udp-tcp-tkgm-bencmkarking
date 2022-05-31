build:
	GOOS=linux GOARCH=amd64 go build -o multiportudpserver .
	chmod 777 multiportudpserver

imgbuild:
	docker build -t lnanjangud653/multiportudpserver:"${IMG_VER}" .

imgpush:
	docker push lnanjangud653/multiportudpserver:${IMG_VER}

