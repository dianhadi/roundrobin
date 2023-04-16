# Round Robin

This is an example project of round robin API that will receive HTTP POSTS and send them to an instance of application API. The round robin API will receive the response from the application API and send it back to the client. 

The archithecture: 
- 1 Round Robin API instance
- Multiple Application API instances

## How to Use

### Prerequisites
This repo is developed with **Docker**, so you need to install docker first before continue. 

These are dependencies that will be run in docker:
1. Golang for programming language. *You may need to install Go if you want to make changes of the code.*

### Getting Started

1. Clone this repository.
2. Open a terminal and navigate to the project root directory.
3. Run the following command to start the project:

```
docker-compose up
```

*You also may use `docker-compose build` and `docker-compose down` for development purpose*

URL will be accessable via http://localhost

## How to Test

### Simple Test
To perform simple testing, you can use the curl command to send a POST request to the Application API endpoint with the following payload:

```
curl -X POST http://localhost:8080/ -H 'Content-Type: application/json' \
-d '{"game":"Mobile Legends", "gamerID":"GYUTDTE", "points":20}'
```

### Down or slow Application Instances

To simulate a scenario where one of the Application API instances goes down or responds slowly, you can use the docker pause command to pause the container for that instance.

```
docker ps
docker pause <CONTAINER ID>
```

*You can also use Docker Desktop UI to pause and unpause docker container*

### Concurrent Access to Round Robin Instance

To simulate a scenario where there are multiple concurrent requests to the Round Robin API, you can use the Postman Runner feature. First, create a new request in Postman to send a POST request to the Round Robin API with the same payload as before. Then, create a collection with this request and add it to a new runner.

In the runner, set the iteration count to the number of concurrent requests you want to send, and set the delay between requests to 0. This will send multiple requests to the Round Robin API at the same time, simulating concurrent access.

You can view the logs for the Round Robin API container to see how the requests are being routed to the different Application API instances.









### Code Documentation 
Code documentation can be found [here](./DOCUMENTATION.md) 