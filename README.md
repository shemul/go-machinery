# Go Machinery example

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)

## About <a name = "about"></a>

A very basic example of task queuing in Machinery and Golang. In the example a golang web server takes input as task and worker consume that task.

## Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them.

```
go environment
```

### Usage <a name = "usage"></a>

A step by step series of examples that tell you how to get a development env running.

start with dependencies download

```
go mod download
```

Spawn a redis server for `Broker` and `ResultBackend`

```
docker run -d -p 6379:6379 redis
```

run the webserver server on port 5000

```
go run main.go server
```

finally, run the worker 

```
go run main.go worker
```


a little demo with curl.
```
curl --request POST 'localhost:5000/send_task' --header 'Content-Type: application/json' --data-raw '{"email": "anyone@gmail.com","subject": "Hi","body":
```

in worker 

![image](https://cdn-images-1.medium.com/max/1600/1*aBmyIx1P7jUIfO3SNGKWqg.png)
