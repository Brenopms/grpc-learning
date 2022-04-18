# GRPC Learning

Project for studying purposes. Learning about how to implement [gRPC](https://grpc.io/) in a microservices environment.

This repository contains four different small applications written in typescript and golang: An authentication service, an order service, a product service and an api gateway to handle http requests for each of the former ones.

This project was based on the two different articles: [NestJS: Microservices with gRPC, API Gateway, and Authentication — Part 1/2](https://levelup.gitconnected.com/nestjs-microservices-with-grpc-api-gateway-and-authentication-part-1-2-650009c03686) and [Microservices in Go with gRPC, API Gateway, and Authentication — Part 1/2](https://medium.com/gitconnected/microservices-with-go-grpc-api-gateway-and-authentication-part-1-2-393ad9fc9d30?source=user_profile---------12----------------------------).

![Application Diagram](https://miro.medium.com/max/1400/1*27q-rUMfeOFEi9KrFndjNg.png)

## Improvements

I have made a couple of changes/improvement regarding the original application:

- Replace [typeorm](https://typeorm.io/) for [prisma](prisma.io)
- Use native [scrypt](https://nodejs.org/api/crypto.html#cryptoscryptpassword-salt-keylen-options-callback) function instead of using [bcrypt](https://www.npmjs.com/package/bcrypt) module
- Add username to User entity
- Change auth Entity to User
- Add validations in product service using go validator
- Add validations in order service using go validator
- Improve error treatment for failing product service client
- Accept order quantity as a parameter
- Calculate order price according

## To-do

- Figure it out how to remove protobuf file duplication for golang
- Containerize each service
- Add kubernetes configuration for running locally and deploy
- Add unit tests
- Add CI/CD configuration using github actions
