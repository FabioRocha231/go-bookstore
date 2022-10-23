## About the project
This project is a simple web server create to absorb the content and studies from golang programming language

### API docs

* `http://localhost:9010/book` get all books METHOD(GET).
* `http://localhost:9010/book` create a book `the object type to send to api it's the type Book` METHOD(POST).
* `http://localhost:9010/book/{id}` get a book by id METHOD(GET).
* `http://localhost:9010/book/{id}` delete a spcecific book METHOD(DELETE) 
* `http://localhost:9010/book/{id}` update a spcecific book METHOD(PUT).

### DATABASE configuration
you have to go to config folder and edit app.go file to update the
```gorm.Open("mysql", "root:123456789@Jj@/study?charset=utf8&parseTime=True&loc=Local")```
to you mySql Database configuration

### JSON OBJECT FORMAT

```
type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
```

### EXAMPLE
```json
{
	"Name": "do lixo ao luxo com golang",
	"Author": "Brazilian books",
	"Publication": "Go horse"
}
```


## Getting started

to run this project you need to clone this repository and use the `go run main.go` command in your CL shell
on just download the bin main and execute it. to execute the bin you have to create a a mysq database
with the root user and 123456789@Jj password with a study named table

### Layout

```tree
├── cmd
│   └── main
│   		└── main.go
├── README.md
├── build
│   ├── admin
│   │   └── Dockerfile
│   └── controller
│       └── Dockerfile
└─ pkg
  ├── config
  │   └── app.go
  ├── utils
  │   └── utils.go 
  └── controllers   
  │    └── book-controller.go
  └── models
  │    └── book.go   
  └── routes

```

A brief description of the layout:

* `main.go` the api entry point
