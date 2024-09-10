# AWS Serverless API in Golang

This repository contains an example of building a Serverless API using AWS Lambda, API Gateway, and the Serverless Framework. The API is written in Golang. 

This is the example code from DEV101 "Build a Serverless API in Go" presented at AWS Summit Toronoto on September 11th. 2024.

## Prerequisites

* AWS account
* Go programming language installed and configured
* AWS CLI installed and configured with your AWS credentials
* Serverless Framework installed globally `npm install -g serverless`

## Setup

1. Clone this repository:
```
git clone https://github.com/your-repo/aws-serverless-api-golang.git
cd aws-serverless-api-golang
```

2. Install the project dependencies: `go get ./...`

## Development
To run the API locally, use the following command:

```
serverless offline
```

This will start the API Gateway locally and watch for changes in your Go code.

## Deployment

To deploy the API to AWS, use the following command:

```
serverless deploy
```

This will package your Go code, create an AWS Lambda function, and configure the API Gateway.

## Usage

After deployment, you can access the API using the provided API Gateway endpoint. The endpoint URL will be displayed in the deployment output.