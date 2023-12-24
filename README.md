# MiniRoute
<p align="center">
  <img src="https://github.com/GLVSKiriti/MiniRoute/assets/116095646/937d590e-b7f5-4b16-b2db-336d2dd0dbee" alt="MiniRoute Logo" width="600">
</p>

An URL shortner web application built with React, TypeScript, GO lang, PostgreSQL. With this application users can shorten the long url and can share those shortened urls with others. Shortened urls are redirected to original long urls on accessing

## ✨ Features
- **Create, Read, Update, and Delete (CRUD) URLS**: Easily manage your URLS with simple and user friendly UI.
- **Containerized the app**: Both backend and frontend are containerized using Dockerfile and docker-compose
- **Reliable Backend**: Ensured reliable backend by comprehensive testing in GO.

## 💡 Technologies Used

<a href="https://www.docker.com/" target="_blank"> <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/react/react-original.svg" width="70px" height="70px"/> </a>&nbsp;
<a href="https://www.docker.com/" target="_blank"> <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/sass/sass-original.svg" width="70px" height="70px"/> </a>&nbsp;
<a href="https://go.dev/" target="_blank"> <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" width="60px" height="60px"/> </a>&nbsp;
<a href="https://www.docker.com/" target="_blank"> <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" width="70px" height="70px"/> </a>&nbsp;

## ⚙ Project Setup 
- **Clone the repository**: 
   ```bash
   git clone https://github.com/GLVSKiriti/MiniRoute.git
   ```
- **Setup .env file**:
   ```bash
   cd ./Backend/ && touch .env
   ```
   Now paste this in .env file by replacing with your database details
  ```bash
  host=<your_postgresql_databse_host>
  port=5432
  user=<your_postgresql_databse_user>
  password=<your_postgresql_databse_password>
  dbname=<your_postgresql_databse_dbname>
  SECRETKEY=<Any_Secret_Key_Of_your_Choice>
   ```
  Make sure that you have created these 2 tables in your database
  ```bash
   CREATE TABLE USERS (uid SERIAL PRIMARY KEY,email varchar(36) NOT NULL,password varchar(100) NOT NULL);

   CREATE TABLE URLMAPPINGS (uid int,id int,longurl text NOT NULL,shorturl varchar(100) NOT NULL,PRIMARY KEY(uid,id),FOREIGN KEY (uid) REFERENCES users(uid));
   ```

- **Just run the below commands and access frontend at `http://localhost:5173/`**
  
- **With docker**: (Very Simple)
  ```bash
    docker compose up
   ```
- **Without docker**
  
  - **Install dependencies**:
   ```bash
   make install
   ```
  - **Start frontend**
   ```bash
   make frontend
   ```
  - **Start backend**
   ```bash
   make backend
   ```
   
## 🤝 Contributing
- You are welcome to come up with new features or resolving issues
- To contribute:
  - Fork this repo and clone the forked repo
  - Dont push your changes directly to main branch
  - Please make sure you raise PR from a new branch not from main

## **Feel Free To Reach Me If You Have Any Doubts In Contributing Or Project Setup** 


