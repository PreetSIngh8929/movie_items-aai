# movie_items-aai
items api for movie app

Api to create ,search ,get movies and book and cancel tickets.Gorilla mux is used as a router in this api.Elastic Search is used to store data.Uber-go/zap package is used as a logger.
# ## Major Technologies
- [ ] Go
- [ ] gorilla mux
- [ ] Elastic Search
- [ ] Uber-go/Zap

### Install & Setup

To setup and install this sample Leaderboard project, follow the below steps:
- Clone this project by the command: 

```
$ git clone https://github.com/PreetSIngh8929/movie_items-aai.git
```
- Then install dependencies using go get 

- Finally, run the below command to start the project.

```
go run main.go
```
# Endpoints

- POST ```/movies``` create a movie
- GET ```"/movies/{id}"``` get movies
- PUT ```"/movies/{id}"``` book ticket
- POST```"/movies/search"``` search a movie
