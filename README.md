go mod init waysbeans <br>
go get -u github.com/gorilla/mux <br>
go get -u gorm.io/gorm <br>
go get -u gorm.io/driver/mysql <br>
go get github.com/go-playground/validator/v10 <br>
<br>
- pkg/mysql/mysql.go before migration.go <br>
- don't forget database/migration to migrate the tables <br>
- and routes to initialize router <br>
- repositories before handlers <br>
- for many to many, needed gorm:"-" in the table models & table relations <br>