# forum

## Authors: Alika03 and RIk0

## How to run 

Run the following commands:
1. For building an exec file: make build
2. For running the project: make run
3. For creating an image and running container: make docker


# Image
1. ![alt text](https://github.com/Alika03/forum/blob/222172b524fb82df778e167b9de6ee0e19ff0f7a/img/main.png?raw=true)

## Objectives

This project consists in creating a web forum that allows :

* communication between users.
* associating categories to posts.
* liking and disliking posts and comments.
* filtering posts.

### SQLite

In order to store the data in your forum (like users, posts, comments, etc.) we used the database library SQLite.

SQLite is a popular choice as embedded database software for local/client storage in application software such as web browsers. It enables you to create a database as well as controlling it by using queries.

To structure your database and to achieve better performance we highly advise you to take a look at the entity relationship diagram and build one based on your own database.

To know more about SQLite you can check the SQLite page.
### Authentication

In this segment the client is able to **register** as a new user on the forum, by inputting their credentials. We also created a **login session** to access the forum and be able to add posts and comments.

We used cookies to allow each user to have only one opened session. Each of this sessions contains an expiration date.

Instructions for user registration:

* It asks for email
    * When the email is already taken return an error response.
* It asks for username
* It asks for password
    * The password is encrypted when stored

The forum is able to check if the email provided is present in the database and if all credentials are correct. It will check if the password is the same with the one provided and, if the password is not the same, it will return an error response.
### Communication

In order for users to communicate between each other, they will have to be able to create posts and comments.

* Only registered users are able to create posts and comments.
* When registered users are creating a post they can associate one or more categories to it.
    * The implementation and choice of the categories is up to you.
* The posts and comments are visible to all users (registered or not).
* Non-registered users is only able to see posts and comments.

### Likes and Dislikes

Only registered users are able to like or dislike posts and comments.

The number of likes and dislikes are visible by all users (registered or not).
### Filter

The project has a filter mechanism, that will allow users to filter the displayed posts by :

* categories
* created posts
* liked posts

Client can look at filtering by categories as subforums. A subforum is a section of an online forum dedicated to a specific topic.

### Docker

For the forum project we used Docker.


This project contains:

* The basics of web :
    * HTML
    * HTTP
    * Sessions and cookies
* Using and setting up Docker
    * Containerizing an application
    * Compatibility/Dependency
    * Creating images
* SQL language
    * Manipulation of databases
* The basics of encryption


### Used packages

* All standard go packages are used.
* github.com/mattn/go-sqlite3
* golang.org/x/crypto/bcrypt
* github.com/satori/go.uuid
