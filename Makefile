run:
	go run src/*.go

install:
	go get -v github.com/gin-contrib/cors
	go get -v github.com/gin-gonic/gin
	go get -v github.com/jinzhu/gorm
	go get -v github.com/jinzhu/gorm/dialects/sqlite
	go get -v github.com/youpy/go-wav
	go get -v github.com/gin-contrib/static