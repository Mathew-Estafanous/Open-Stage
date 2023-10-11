<p align="center">
  <img src="https://user-images.githubusercontent.com/56979977/116730566-a4dcae00-a9b6-11eb-95b7-39239fe35386.png" width="125">
  <h1 align="center">Open Stage</h1>

  <p align="center">
    A live Q&A platform that stimulates conversation, even from the quietest members.
    <br />
    <a href="https://open-stage-web.herokuapp.com/"><strong>Try It Out</strong></a>
    ·
    <a href="https://open-stage-api.herokuapp.com/docs"><strong>Explore API Docs</strong></a>
  </p>
</p>

# About The Project

Open Stage is a platform that bridges the communication gap by providing a platform for safe
and open dialogue. Simply create a room and allow others to voice their questions and upvote other's
questions as well. All done with the option of remaining anonymous to remain free from worries and judgement.

### Project Structure
The project is split into both the backend, written in Go, and the frontend written using React.js.

When structuring the backend, a key focus was to follow the **Domain Driven Design Principle.** This principle focuses on
having the domain/business logic at the heart of the app and all external implementation details are then hidden in outer layers.
When looking at the different packages of the backend, the both the [domain](https://github.com/Mathew-Estafanous/Open-Stage/tree/main/backend/domain)
and [service](https://github.com/Mathew-Estafanous/Open-Stage/tree/main/backend/service) packages contain the main business logic
of the application, while other packages implement logic that is dependant on external dependencies like a PostgreSQL database, etc.

Another important goal while building the backend of this app was to ensure wide test coverage throughout all layers of the app. This
ensures the validity of the code so that we can be confident that the app works when all the tests pass.

### Documentation
**API Docs: https://open-stage-api.herokuapp.com/docs**

At the heart of the API documentations is the swagger yaml files. Majority of the documentation was created using the
[go-swagger](https://github.com/go-swagger/go-swagger) library, in which the majority of this documentation can be found
in the [docs](https://github.com/Mathew-Estafanous/Open-Stage/tree/main/docs) package.

You can also find the application's database schema as both a [SQL file](https://github.com/Mathew-Estafanous/Open-Stage/tree/main/backend/docs/sql)
or take a look at the [schema diagram.](https://dbdiagram.io/d/606262f8ecb54e10c33dd900)

# Getting Started
Setting up the project in a local development environment should be simple and easy to do. The following are
steps of what you need to do, to get the project up and running.

### Prerequisites
The following must be installed on your machine before running or contributing.
* [Golang](https://golang.org/) - The API is built using Golang.
* [Docker](https://www.docker.com/) - Docker is used to easily spin up needed services, like a local database.
* [Node](https://nodejs.org/en/) - Node.js is important if you intend on locally running the frontend of the app.

### Running
``git clone https://github.com/Mathew-Estafanous/Open-Stage.git`` - Clone the repository.

``cd ./backend`` - Navigate to the backend directory of the project.

``docker-compose up`` - Startup all services, such as the REST API, database and redis server.

All the required services should be up and running, including the API on ``:8080``.

### Contributing
This project is still underdevelopment and continues to be improved upon on a regular basis. Contributions
are always welcome whether they are issue reports or pull requests to include a feature. Testing is very important,
which is why unit tests should be part of any new feature.

# Contact
**Mathew Estafanous -** mathewestafanous13@gmail.com

**Personal Website -** https://mathewestafanous.com/

**Project Link -** https://github.com/Mathew-Estafanous/Open-Stage
