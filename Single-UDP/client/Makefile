build:
	GOOS=linux GOARCH=amd64 go build -o singleudpclient .
	chmod 777 singleudpclient

imgbuild:
	docker build -t lnanjangud653/singleudpclient:"${IMG_VER}" .

imgpush:
	docker push lnanjangud653/singleudpclient:${IMG_VER}
