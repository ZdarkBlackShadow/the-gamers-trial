# 🎮 The Gamers Trial

Welcome to The Gamers Trial! This project is an image-based quiz web application, designed to test players' knowledge of the video game universe. Users answer questions based on visual elements (such as mini-maps or screenshots) to earn points.

## 🛠️ Prerequisites

To run and develop the application, you will need the following tools installed on your machine:

- Git : For cloning the repository.

- Go : Version 1.20 or higher.

- MySQL : Required for the GORM database.

## 🚀 Configuration and Installation

### 1. Project Cloning

```bash
git clone https://github.com/ZdarkBlackShadow/the-gamers-trial.git
cd the-gamers-trial
```

### 2. Environment Configuration

Rename the example environment file and update the database connection information.

```bash
cp .env.example .env
```
`Modify the .env file with your database information (DB_HOST, DB_USER, etc.)`


### 3. Database Initialization and Migration

This command will create the database structure (tables) and, if you have added the seed() function, insert the basic data (users, default questions, etc.).

```go
go run migration/database.go reset-db
```


### 4. Starting the Application

Run the Go server. The application should be accessible at http://localhost:8080 (unless configured differently in your code or .env file).

```go
go run main.go
```


## 📁 Project Structure

- `cmd/` : Functions for executing one-off operations (e.g., migrations, seeding).

- `config/` : Application configuration (database, logs).

- `controller/` : Handles HTTP requests (uses Fiber).

- `service/` : Contains the business logic.

- `repository/` : Database access layer (uses GORM).

- `model/` : Data structures (GORM entities, requests, views).

- `routes/` : Definition of API routes.

- `views/` : HTML template files.

- `doc/` : documantation file