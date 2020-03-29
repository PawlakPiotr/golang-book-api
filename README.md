# Golang API for library web application

Web API interfeces written in Golang. It supports rest api requests using third party libraries (see libraries section). The project includes the ability to add, remove and edit books (basic CRUD) in the library. The project will be further developed. Also there will be a project with a graphic interface communicating with the api interfaces of this project.

## Technological stack 

### API 
API interfaces are written in Golang. Go can be downloaded from offical [site](https://golang.org/dl/). Go version 1.13 is used in this project.

`go1.13.1 darwin/amd64`


### Database
The database used in this project is MongoDB. You can download it from official [site](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/). The project needs an extra `mongo-diver` golang library. More information is in the next section (see third party libraries)

### Third party libraries

* MUX - `github.com/gorilla/mux`
* MongoDB - `go.mongodb.org/mongo-driver/mongo`

Each package can be downloaded by running the command `go get <package name>`. For example:
```
go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
```

Packages should be installed on `$GOROOT `or in the project `$GOPATH`. For example for this project `GOPATH` can be set to `/home/userName/golang-book-api/`.

## Configuration

Default api configuration is in the file `localhost.json` in the directory` src / gopang-book-api / config / `. It looks like this:

```
{
    "database": {
        "connection": "mongodb://localhost:27017",
        "name": "Library",
        "collections": ["books"]
    },
    "host": "localhost",
    "port": "8080"
}
```

## Building project

**Note:** To build project you need to have `GOPATH` set properly.

Go to `golang-book-api` directory and build project using golang cli.

```
cd src/golang-book-api
go build
```

After this you should see executable file called `golang-book-api` in `src/golang-book-api` directory.

## Running project

To start API simply run script:

`buildAndRun.sh`

If you want to run it manually go to `src/golang-book-api` directory and run file (you need to build it first):

*MacOS/linux*
```
./golang-book-api
```
*Windows*
```
golang-book-api.exe
```

After that you should see these logs of application in terminal: 
```
 2020-03-24 16:42:13 [DEBUG] Connceted do database Library
 2020-03-24 16:42:13 [INFO] [Starting API] Listening on PORT: 8080
```

### API Endpoints

If application starts successfully you can now acces API endpoints. Example of API endpoints:

* Returns all books stored in database
```
[GET] /api/v1/books

example:
http://localhost:8080/api/v1/books
```

* Returns book with given `ID` stored in database
```
[GET] /api/v1/books/{ID}

example:
http://localhost:8080/api/v1/books/4a18ad7a-56dd-4142-9b3d-299222cdca65
```


* Adds book with given parameters in `request.body`
```
[POST] /api/v1/books/add

example:
http://localhost:8080/api/v1/books/add

request body example:
{
    "title": "new title",
	"author" : {
		"firstname" : "FName",
		"lastname" : "LName"
	},
	"category" : "test CATEGORY",
	"cags" : ["new_tag1", "new_tag2"]
}
```

**Note** A swagger with all API endpoints will be prepared later.

## Aplication structure

```
├── src
│   └── golang-book-api
│       ├── config [Directory with API configuration files]
│       ├── controllers [Directory with all controllers implementing buisness logic]
│       ├── database [Directory with files connecting with MongoDB]
│       ├── main.go [Main application file]
│       ├── model [Directory with files which declares models of application]
│       ├── routes [Directory with defined routes of API endpoints]
│       └── utils [Directory with files implementing errors responses, logging engine etc.]
└── tests
    └── scripts [Bash scripts testing API endpints]
```
