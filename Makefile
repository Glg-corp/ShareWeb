run:
	go run src/*.go

build:
	go build -o ShareWeb -i src/*.go

install:
	go get -v github.com/gin-contrib/cors
	go get -v github.com/gin-gonic/gin
	go get -v github.com/jinzhu/gorm
	go get -v github.com/jinzhu/gorm/dialects/sqlite