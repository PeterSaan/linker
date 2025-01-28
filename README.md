# Linker

## Starting project

**Make sure you have docker installed**

- Clone repo
- Run in project root `cp ./backend/.env.example ./backend/.env` 
- If you are running the project for the first time `docker-compose up --build -w` otherwise you can run `docker-compose watch`  
    - Note if there are any package changes etc, you will have to run `docker-compose up --build -w` again

## MVP

1. The user must create an account to use the app
2. The company must create an account to use the app
3. The user only sees companies that are hiring job titles picked by the user
4. The company only sees users that are looking for a job in the company's open roles
5. Users and companies can swipe right (interested) or left (not interested) on the opposite group
6. Users and companies get notified when they swiped right but the other swiped left
7. In the case of a sucessful match (both swiped right), a temporary chat room (based on activity) will be opened
