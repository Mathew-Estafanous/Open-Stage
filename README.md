<p align="center">
  <img src="https://user-images.githubusercontent.com/56979977/116730566-a4dcae00-a9b6-11eb-95b7-39239fe35386.png" width="125">
  <h1 align="center">Open Stage</h1>

  <p align="center">
    A live Q&A platform that stimulates conversation, even from the quetest members.
    <br />
    <a href="https://open-stage-platform.herokuapp.com/docs"><strong>Explore API Docs</strong></a>
    ·
    <a href="https://github.com/Mathew-Estafanous/Open-Stage/issues"><strong>Report An Issue</strong></a>
  </p>
</p>

# About The Project

Open Stage is a platform that bridges the communication gap by providing a platform for safe
and open dialogue. Simply create a room and allow others to voice their questions and upvote other's
questions as well. All done with the option of remaining anonymous to remain free from worries and judgement.

### Project Structure
The project is structured using **Domain Driven Design principles** with the goal of focusing on
the actual business logic of the application while keeping the specific technology details decoupled.
Business details can be found in the [domain](https://github.com/Mathew-Estafanous/Open-Stage/tree/main/domain)
and [service](https://github.com/Mathew-Estafanous/Open-Stage/tree/main/service) packages while having external
dependencies like the infrastructure and handlers outside the specific business model.

### Documentation
**API Docs: https://open-stage-platform.herokuapp.com/docs**

At the heart of the API documentations is the swagger yaml files. Majority of the documentation was created using the
[go-swagger](https://github.com/go-swagger/go-swagger) library, in which the majority of this documentation can be found
in the [docs](https://github.com/Mathew-Estafanous/Open-Stage/tree/main/docs) package.

You can also find the application's database schema as both a [SQL file](https://github.com/Mathew-Estafanous/Open-Stage/blob/main/_sql/schema.sql)
or take a look at the [schema diagram](https://dbdiagram.io/d/606262f8ecb54e10c33dd900)

# Getting Started
Setting up the project in a local development environment should be simple and easy to do. The following are
steps of what you need to do, to get the project up and running.

### Prerequisites
The following must be installed on your machine before running the project.
* [Golang](https://golang.org/) - The API is built and run using the golang language.
* [Docker](https://www.docker.com/) - Docker is used to easily spin up needed services, like a local database.

### Running
``git clone https://github.com/Mathew-Estafanous/Open-Stage.git`` - Clone the repository.

``docker-compose up`` - Startup all services, including API and database.

All the required services should be up and running, including the API on ``:8080``.

### Contributing
This project is still underdevelopment and continues to be improved upon on a regular basis. Contributions
are always welcome whether they are issue reports or pull requests to include a feature. Testing is very important,
which is why unit tests should be part of any new feature contribution.

# Contact
**Mathew Estafanous -** mathewestafanous13@gmail.com

**Personal Website -** https://mathewestafanous.com/

**Project Link -** https://github.com/Mathew-Estafanous/Ur-Codebin
