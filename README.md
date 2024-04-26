

# Weather API
This repository contains a simple Go project that provides information about weather conditions inside a given Market.

## Getting Started

These instructions will guide you through setting up and running this project on your local machine or docker.

## Installation
    

## Local
### Prerequisites

Before you begin, ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/dl/).

You should have tile38 server running on your system and define environment variable **TILE38_HOST** such as TILE38_HOST=localhost:9851. You can download it from [tile38.com](https://tile38.com/topics/installation)

You should have Redis server running on your system and define environment variable **REDIS_HOST** such as TILE38_HOST=localhost:6379. You can download it from [redis.io](https://redis.io/downloads/)

## Docker
### Prerequisites

Before you begin, ensure you have Docker installed on your system. You can download it from [docker.com](https://www.docker.com/get-started/).

### Running docker containers

```
docker-compose up --build
```




## Populating Database

Before you begin, ensure you have Python installed on your system. You can download it from [python.org](https://www.python.org/downloads/).

In order to initialize tile38 database by populating with example data, you can run the python script in path **scripts/import_data_tile38.py** while tile38 server is running on your environment. 

Before running the script make sure you have installed **Redis** for python.
```
pip install Redis
python import_data_tile38.py
```


## Usage

After running the program on your system, you can send request to server with required parameters.
If you have curl installed on your system ([Curl](https://curl.se/download.html)) :

```
curl -X GET -H "Content-Type: application/json" -d "{\"lat\": 41.13, \"long\": -74.18}" http://localhost:8080/weather

```

