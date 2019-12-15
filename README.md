# reader
 sample REST API how to make skeleton code on golang. for easy learn

# depdency
depedency for installing this skeleton code.

```
go get github.com/badoux/checkmail
go get github.com/jinzhu/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm/dialects/mysql" //If using mysql 
go get github.com/jinzhu/gorm/dialects/postgres //If using postgres
go get github.com/joho/godotenv
go get gopkg.in/go-playground/assert.v1
```

# How to use
we have structure like this,and you can create your own project like this

```
config        //for configuration about create middleware,create route version 1.0 
controller   //make controller you can make own process using models query etc.
core         //core of enginne this skeleton it's about processing it can be use more than one project so it's useable function
database     //you can make operation about migration,and seeder on this folder
models       //make models data im using ORM, and schema for migration data.
main.go     //main process from config,and core first running apps
.env        //environment you can make some easy config from this file
```
# contribution

you can request merge if you have some good idea.im open how to make this can be reusable
clean and easy to learn for newbie from golang
