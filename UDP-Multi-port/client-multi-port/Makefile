build:
	GOOS=linux GOARCH=amd64 go build -o multiportudpclient .
	chmod 777 multiportudpclient

imgbuild:
	docker build -t lnanjangud653/multiportudpclient:"${IMG_VER}" .

imgpush:
	docker push lnanjangud653/multiportudpclient:${IMG_VER}
