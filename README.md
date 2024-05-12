# Go Docker API Dashboard Application

This application provides an interface to interact with Docker containers, allowing users to list containers and serve a dashboard.

## Prerequisites

- Go > 1.20
- Docker API version > 1.45 

## Installation

Clone the repository to your local machine:

git clone https://github.com/surajupadhaya/containerDashboard.git

cd containerDashboard

## Dependencies
To install the necessary dependencies, run:

go mod init github.com/docker

go mod tidy

This will download the required packages, including the Docker client.

## Running the Application
Before running the application, ensure that Docker is running on your system.

To start the application, use:

go run main.go

The application will start an HTTP server on port 8080 and can be accessed at http://localhost:8080

## Endpoints
/ContainerHealth: Lists all Docker containers and their status in json 
![image](https://github.com/surajupadhaya/containerDashboard/assets/26745462/d08962c4-1188-46f5-adfd-8808f10b77e4)


/Dashboard: Serves the static dashboard HTML file.

![image](https://github.com/surajupadhaya/containerDashboard/assets/26745462/2b56ba04-bba4-43bc-902e-62139dbc3ddc)
