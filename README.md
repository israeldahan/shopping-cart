## Shopping Cart Project

### Description
This project is a shopping cart application that allows users to add items to their cart and get Bill.

### Usage
Run the program: `go: go mod download && go mod verify && go run .` // default port is 2345
Run the Program with Docker: `docker-compose up -d` // default port is 3000

when you run the program, you will see the following request:
```
POST http://localhost:<port>/bill 
    Body: {
        "cart": [
            { "name": "milk" },
            { "name": "milk" },
            { "name": "bread"},
            { "name": "bread"},
            { "name": "milk" },
            { "name": "bread"},
            { "name": "bread"},
            { "name": "bread"},
            { "name": "milk" },
            { "name": "bread"}
        ]
    }
```