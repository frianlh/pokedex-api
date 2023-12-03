# Pokedex API
Pokedex API project in Go language

## Prerequisite
1. Go 1.21.3 or latest
2. Git
3. Docker
4. Fiber ([documentation](https://gofiber.io/))
5. GORM ([documentation](https://gorm.io/))
6. Go Migrate ([documentation](https://github.com/golang-migrate/migrate/tree/master/database/postgres))
7. Postgres ([documentation](https://www.postgresql.org/docs/))

## Project Setup
Run below commands to make project up and running
1. Clone Repository
   ```bash
   # Via SSH
   git clone git@github.com:frianlh/pokedex-api.git

   # Via HTTPS
   git https://github.com/frianlh/pokedex-api.git
   ```
2. Copy .env.example to .env
   ```bash
   # If using command prompt Windows
   copy .env.example .env
   
   # if using terminal, Ubuntu
   cp .env.example .env
   ```

3. Install Dependencies
   ```bash
   # Via Makefile
   make clean_module
   
   # Not via Makefile
   go mod tidy
   ```

   > Note: To **run project** and **run unit testing**, run [development mode](#run-project-development-mode) first.
4. Run Project
   ```bash
   # Via Makefile
   make run
   
   # Not via Makefile
   go run app/main.go
   ```

5. Run Unit Testing
   ```bash
   # Via Makefile
   make run
   
   # Not via Makefile
   go run app/main.go
   ```

## Run Project Development Mode
1. Run Development Mode
   ``` bash
   # Via Makefile
   compose_up
   
   # Not via Makefile
   docker-compose down
   doker network remove pokedex-api_pokedex_network; true
   doker-compose up --force-recreate
   ```

## Project Documentation
1. [API Documentation](https://www.postman.com/avionics-physicist-83460159/workspace/pokedex-api/collection/31514600-63602764-130e-4dcc-840f-2932906a3b22?action=share&creator=31514600)
2. [Database Schema](https://dbdiagram.io/d/Pokedex-656c270f56d8064ca045061a)

## Credential
1. Role Admin
   ```
   Email: dev.admin@gmail.com
   Pass: DevAdmin@123
   ```
2. Role User
   ```
   Email: dev.user@gmail.com
   Pass: DevAdmin@123
   ```

## Task
- [x] Documentation
  - [x] Database Schema
  - [x] API Documentation (Postman)
- [x] Pokedex API
  - [x] Authentication
    - [x] Login
  - [x] Monster Category
    - [x] List All Monster Category
  - [x] Monster Type
    - [x] List All Monster Type
  - [x] Monster
    - [x] Create Monster (only Admin can access)
    - [x] Get List Monster
    - [x] Get Detail Monster
    - [x] Update Monster (only Admin can access)
    - [x] Delete Monster (only Admin can access)
    - [x] Update Monster Catched Mark
- [x] Containerization
- [x] Manual Test and Unit Test