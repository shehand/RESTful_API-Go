# RESTful_API service - Go

## Installation

First of all make sure to install Go Programming Language
in your computer. Try [this link](https://golang.org/doc/install) to follow the installation.

And add the `GOPATH` path variable if it is not set. Use [this link](https://github.com/golang/go/wiki/SettingGOPATH) to follow a tutorial to set the `GOPATH` variable

Download the source code. You can either download it as a zip file and 
extract it or simply type the command in the terminal or bash or cmd,

`git clone 
https://github.com/shehand/RESTfu_API-Go.git`

change your directory into the project folder. And run the command,

`go run main.go`

Then it will start the server on http://localhost:8000. To test the api use postman or curl and make sure to follow the url like given below.

> NOTE
> 
> If you failed to start the server try to configure your database 
details with the project. To update database details find the file .env in the project directory or .env.example and change the database detils. If the error is about some missing dependencies try like below.
>
> if it says `github.com/gorilla/mux` is missin for a example try in the terminal
> 
>`go get github.com/{package-name}` here the package name is `gorilla/mux`
>

### Some Basic API urls

* http://localhost:8000/api/user/new -> Create User
* http://localhost:8000/api/user/login -> login
* http://localhost:8000/api/api/me/contacts -> get contacts

Feel free to open issues and PR's are welcome. Happy coding... ;)