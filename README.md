
# AWS Cache Project
This project is an AWS cache that allows you to pull data from the AWS API and store it in a MongoDB database. It was developed while working at Xero.


## Features
Fetches data from the AWS API.
Stores the fetched data in a MongoDB database.
Provides a RESTful API to interact with the cached data.

### Prerequisites
Docker
Docker Compose
Go (Golang)
AWS CLI

## Setup
1. Clone the Repository
git clone https://github.com/rifaavalon/go-cache-aws.git

`cd go-cache-aws`

3. Set Up AWS Credentials
Ensure your AWS credentials are configured locally. You can set them up using the AWS CLI:
`aws configure`

4. Build and Run the Docker Containers
Use Docker Compose to build and run the Docker containers for your Go application and MongoDB:

`docker-compose up --build`


This will start your Go application on port 8080 and MongoDB on port 27017.

## Usage
Running the Application
To run the application locally without Docker, ensure you have Go installed and run:

`go run main.go`


## API Endpoints
`GET /instances`: Retrieve all instances.

`GET /instances/:id`: Retrieve a specific instance by ID.

`POST /instances`: Create a new instance.


## Example Requests
Get All Instances
`curl http://localhost:8080/instances`
Get a Specific Instance
`curl http://localhost:8080/instances/1`
Create a New Instance
`curl -X POST http://localhost:8080/instances -d '{"field1":"value1", "field2":"value2"}' -H "Content-Type: application/json"`

## Project Structure

```
.
├── Dockerfile
├── docker-compose.yml
├── main.go
├── aws.go
├── db.go
├── handlers.go
├── ec2.go
├── go.mod
├── go.sum
└── README.md
```
## Configuration


### Environment Variables

MONGO_URI: The URI for connecting to MongoDB. Set this in your Docker Compose file or as an environment variable.


### Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.


### License
This project is licensed under the MIT License. See the LICENSE file for details.
