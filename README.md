# GO STORE AND WEBSOCKETS

This project works as a backend for a store, it has products, that has relations with brands and categories,
also allows to create users with priviledges.

For security JWT  protocol is used.

New functions have been included, users registered as *postMaker* can now create new posts and everyone can see the posts inside the backend of the store.

## Documentation
The documentation is stored in the following addres.

[Swagger documentation](http://jose827corrza.github.io/go-store-websockets/ "Swagger documentation")

## Deploy URL

[Deploy](https://go-store-websockets-production.up.railway.app)



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


To run the server locally, run the following command

> go run main.go

**CHECK THE REQUIRED VALUES IN THE .env FILE, GUIDE USING THE .env.example**

## Docker for Local

To build the docker image, move to database folder and run the following commands

> docker build . -t go-estore-websockets

Next to it, to start the container, 

> docker run -p 54321:5432 go-estore-websockets