# Golang Authentication Template

This repository provides a simple example of how to implement authentication in a Golang application. It uses a PostgreSQL database to store user information, and JWT authentication to authenticate users.

## Getting started

To get started, clone this repository and install the dependencies:

git clone https://github.com/fahmisalman/Go-auth
cd go-auth
go mod init
go mod tidy
Once the dependencies are installed, you can start the application:

go run main.go
The application will listen on port 8080 by default. You can access the application in your web browser at http://localhost:8080/

## Authentication

To authenticate, you can send a POST request to the /auth/login endpoint with the following JSON body:

``` JSON
{
  "username": "user@example.com",
  "password": "password"
}
```
Use code with caution. Learn more
If the authentication is successful, the server will return a JSON Web Token (JWT) in the Authorization header of the response. You can then include this JWT in the Authorization header of subsequent requests to authenticate yourself to the server.

## Registering a new user

To register a new user, you can send a POST request to the /auth/register endpoint with the following JSON body:

``` JSON
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "username": "john.doe",
  "password": "password"
}
```
Use code with caution. Learn more
If the registration is successful, the server will return a JSON response with the following fields:

id: The ID of the new user
name: The name of the new user
email: The email address of the new user
Example usage

```
# Login
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "username": "user@example.com",
    "password": "password"
  }' \
  http://localhost:8080/auth/login

# Register a new user
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "username": "john.doe",
    "password": "password"
  }' \
  http://localhost:8080/auth/register
```

## Deployment

To deploy this application, you can use a variety of different tools and services. For example, you could deploy the application to a cloud platform such as Heroku or AWS Elastic Beanstalk. You could also deploy the application to a self-managed server using a tool such as Docker or Kubernetes.

## Conclusion

This repository provides a simple example of how to implement authentication in a Golang application. It is a good starting point for developers who are new to authentication, or who are looking to implement a specific authentication method in their application.
