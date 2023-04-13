# GO STORE AND WEBSOCKETS

#dependencies used and its reason

> go get github.com/gorilla/mux

For routing.

> go get github.com/gorilla/websocket

For the websocket functionality

> go get github.com/joho/godotenv

For environment variables, very important.

> go get github.com/golang-jwt/jwt/v4

This will ease the set of JWT protocol

> go get github.com/lib/pq

To enable the use of Postgres

## Development

Make sure you have the .env file in your root folder, containing the requested variables avoiding the typos
you can guide with the .env.example file.

**VERY IMPORTANT TO NOTE THAT THE PORT CURRENTLY IS RECEIVING A STRING, WHICH ALSO NEEDS BESIDES THE PORT NUMBER, THE ":". THIS CAN BE SOLVED/IMPROVED IN A FUTURE**

To run the server locally, run the following command

> go run main.go

## Docker for Local

To build the docker image, move to database folder and run the following commands

> docker build . -t go-estore-websockets

Next to it, to start the container, 

> docker run -p 54321:5432