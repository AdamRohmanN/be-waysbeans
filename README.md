go mod init waysbeans <br>
go get -u github.com/gorilla/mux <br>
go get -u gorm.io/gorm <br>
go get -u gorm.io/driver/mysql <br>
go get github.com/go-playground/validator/v10 <br>
go get -u github.com/golang-jwt/jwt/v4 <br>
go get github.com/joho/godotenv <br>
<br>
- pkg/mysql/mysql.go before migration.go <br>
- don't forget database/migration to migrate the tables <br>
- and routes to initialize router <br>
- repositories before handlers <br>
- for many to many, needed gorm:"-" in the table models & table relations <br>
- "404 page not found" router pada routes.go belum di initialize <br>