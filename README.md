# lolhichat

## About
LolHiChat is a simple web forum where users can post whatever they feel like posting!
Unfortunately this has not been completed within the CVWO deadlines :sad:

## Features
There are 4 roles that users of LolHiChat can interact with a discussion thread:
- [x] They can be a casual user who can view the discussion threads that they search 
- [x] They can be a logged-in user, who can interact with the discussion threads on the site
- [ ] They can be a subscriber to a thread, who will have quick access to the thread in their feed
- [ ] They can be a host to a thread, which grants them admin powers to moderate the discussion in the thread

As a casual user, I can
- [x] Search for threads that I am interested in, so that I can see what others have to say about topics of my interest.
- [x] Sign up for an account and log in, so that I can interact more meaningfully with the threads on the site

As a logged-in user, I can
- [x] Comment under the discussion threads, so that I can share my views to the other users of the platform
- [x] Reply to the comments of others, so that I can participate in the discussion with them
- [x] Like the comments of others, so that I can show support for opinions that I align with

As a subscriber, I can
- [ ] Access the threads that I have been subscribed to on my homepage feed, so that I have easier access to the topics that I am interested in

As a host, I can
- [ ] Create the discussion thread, so that I can encourage like-minded others to talk about a topic I am interested in
- [ ] Pin comments in the discussion thread, so that I can highlight comments that I wish to be shown to others when they first enter the thread
- [ ] Delete comments in the discussion thread, so that I can keep the conversation in the thread civil 

## Implementation
This project is split up into 2 parts:
1. Frontend, written in React
2. Backend, written in Go

### Running the Frontend
1.  After navigating to the base directory of lolhichat, navigate to the frontend directory
    ```bash
    $ cd frontend
    ```
2.  Install the dependencies as defined in `package.json`
    ```bash
    $ npm install
    ```
3.  Run the react app
    ```bash
    $ npm start
    ```

### Running the Backend
1.  After navigating to the base directory of lolhichat, navigate to the backend directory
    ```bash
    $ cd backend
    ```
2.  Create a postgresql database with the below commands, and create the required tables based on `database/db.sql`
    ```
    $ psql -U <your-username>
    postgres=# CREATE DATABASE "lolhichat" WITH OWNER <your-username>
    postgres=# GRANT ALL PRIVILEGES ON DATABASE "lolhichat" TO <your-username>
    postgres=# \i database/db.sql
    ```
3.  Create a `.env` file to set the required environment variables in the project
    ```bash
    DB_HOST="localhost"
    DB_USER=<your-username>
    DB_PASSWORD=<your-password>
    DB_NAME="lolhichat"
    DB_PORT="5432"

    JWT_SECRET_KEY="rv9unfvedwnvund^%$#"
    ```
4.  Install the dependencies and run the go app with
    ```bash
    $ go run cmd/server/main.go
    ```

## Project Structure
### Frontend 
The organisation of the frontend is as such, with the react root found in `src/index.tsx` 
and the main react app found in `src/App.tsx`
```md
frontend/
├── public/
├── src/
│   ├── api/
│   ├── components/
│   ├── helpers/
│   ├── hooks/
│   ├── pages/
│   ├── App.css
│   ├── App.test.tsx
│   ├── App.tsx
│   ├── index.css
│   ├── index.tsx
│   ├── logo.svg
│   ├── react-app-env.d.ts
│   ├── reportWebVitals.ts
│   └── setupTests.ts
├── package-lock.json
├── package.json
└── tsconfig.json
```

The other subfolders under `src/` are listed below
* `api/` contains the setup code for the axios api client, and the interfaces to receive api responses from the backend
    ```md
    frontend/src/api
    ├── apiClient.tsx
    ├── apiResponse.tsx
    └── axios.d.ts
    ```
* `components/` contains React components from the pages that have been abstracted out
    ```md
    frontend/src/components
    ├── CommentCard.tsx
    ├── DirectToLogin.tsx
    ├── Logout.tsx
    ├── NavBar.tsx
    ├── NewComment.tsx
    ├── QuickSearchBar.tsx
    └── ThreadCard.tsx
    ```
* `helpers/` contains general helper functions for the app
    ```md
    frontend/src/helpers
    ├── authChecks.tsx
    └── errorMessage.tsx
    ```
* `hooks/` contains custom react hooks that are used for the app
    ```md
    frontend/src/hooks
    ├── useApiClient.tsx
    └── useAuth.tsx
    ```
* `pages/` contains the main webpages that users will mainly be routed to
    ```md
    frontend/src/pages
    ├── Home.tsx
    ├── Login.tsx
    ├── NewThread.tsx
    ├── SearchResults.tsx
    ├── SignUp.tsx
    ├── ViewComment.tsx
    └── ViewThread.tsx
    ```

### Backend
The organisation of the backend is as such, with the main go app found in `cmd/server/main.go`
```md
backend/
├── cmd/
│   └── server/
|       └── main.go 
├── internal/
│   ├── controllers/
│   ├── dataaccess/
│   ├── database/
│   ├── helpers/
│   ├── middleware/
│   ├── models/
│   └── routes/
├── go.mod
└── go.sum
```

The other subfolders under `internal/` are listed below:
* `controllers/` contains the handlers for the various http routes
    ```md
    backend/internal/controllers
    ├── home.go
    ├── likes.go
    ├── login.go
    ├── newComment.go
    ├── newThread.go
    ├── renewToken.go
    ├── search.go
    ├── signUp.go
    └── viewThread.go
    ```
* `dataaccess/` contains functions that interact with the database
    ```md
    backend/internal/dataaccess
    ├── comments.go
    ├── likes.go
    ├── threads.go
    └── users.go
    ```
* `database/` contains the database schema, as well as the setup code to link the backend app to the postgresql database
    ```md
    backend/internal/database
    ├── db.sql
    └── setup.go
    ```
* `helpers/` contains general helper functions for the app
    ```md
    backend/internal/helpers
    ├── hashing.go
    ├── tokens.go
    ├── url.go
    └── username.go
    ```
* `middleware/` contains code for the middleware to process the http headers
    ```md
    backend/internal/middleware
    ├── auth.go
    └── headers.go
    ```
* `models/` contains the type structures that model the database entries
    ```md
    backend/internal/models
    ├── comment.go
    ├── like.go
    ├── thread.go
    └── user.go
    ```
* `routes/` contains the code for assiging the http routes to the various handler and middleware functions
    ```md
    backend/internal/routes
    └── routes.go
    ```

## Tech Stack
### Frontend
The frontend is written in [TypeScript](https://www.typescriptlang.org/) with [React](https://reactjs.org/), and makes use of the following tools:
- [Axios](https://www.npmjs.com/package/axios)
- [Material UI](https://mui.com/) 
- [React Router Dom](https://www.npmjs.com/package/react-router-dom)
- [React Time Ago](https://www.npmjs.com/package/react-time-ago)

More details can be found under `frontend/package.json`.

### Backend
The backend is written in [Go](https://go.dev/) with [Gin Web Framework](https://gin-gonic.com/en/), making use of the following tools:
- [CORS middleware for Gin](https://github.com/gin-contrib/cors)
- [Gorm](https://gorm.io/) along with its [postgresql driver](https://github.com/go-gorm/postgres)
- [GoDotEnv](https://pkg.go.dev/github.com/lpernett/godotenv)
- [Golang JWT](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)
- [Golang crypto](https://pkg.go.dev/crypto)
- [Golang sqids](https://sqids.org/go)

More details can be found under `backend/go.mod`.