# Go Fiber Web Application - User Registration and Login

This is a simple web application built using the **Go Fiber** framework. The application allows users to register with their name, email, and password, and then log in using their email and password. The application hashes passwords for security before storing them in a MySQL database.

## Features

- **User Registration**: Allows users to register with their name, email, and password.
- **User Login**: Allows users to log in using their email and password.
- **Password Hashing**: Passwords are securely hashed before being stored in the database.
- **Error Handling**: The application handles errors like incorrect credentials and missing fields.
- **Environment Variables**: Sensitive data such as the database connection string is stored in environment variables.

## Technologies Used

- **Go Fiber**: Web framework for building the API.
- **MySQL**: Database for storing user data.
- **bcrypt**: Used to hash passwords.
- **godotenv**: Loads environment variables from a `.env` file.

## Setup Instructions

### Prerequisites

Before running the application, make sure you have the following installed:

- [Go](https://golang.org/doc/install)
- [MySQL](https://dev.mysql.com/downloads/)
- [Postman](https://www.postman.com/downloads/) (for testing the API)

### Environment Setup

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/your-username/go-fiber-registration-login.git
   cd go-fiber-registration-login
