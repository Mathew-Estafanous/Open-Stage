# Open-Stage
A platform that bridges the communication gap by providing a platform for safe
and open dialogue that is free from judgement. Simply create a room
and allow others to voice their questions and upvote other's questions as well. 
All done with the option of remaining anonymous to remain free from worries of possibly
being judged.

## Project Structure
The project is structered using **Domain Driven Design principles** with the goal of focusing on
the actual business logic of the application while keeping the specific technology details decoupled.
The business details can be found in the [domain](https://github.com/Mathew-Estafanous/Open-Stage/tree/main/domain) 
and [service](https://github.com/Mathew-Estafanous/Open-Stage/tree/main/service) packages while having external
dependencies like the infrastructure and handlers outside of the specific business model. This structure allows 
for easy unit testing since each layer is independant of the other.

### Database Schema
The following is an image of the database schema that is used for the web app.
[![schema](https://user-images.githubusercontent.com/56979977/115421799-494a3d80-a1ca-11eb-9a6b-975733597be2.png)](https://dbdiagram.io/d/606262f8ecb54e10c33dd900)

## TODO
The project is still under development and is continually changing so here are a few things underworks
- [X] Moving from MySQL to PostgreSQL.
- [X] Refactor the database schema to a better format.
- [X] Host the API on Heroku.
- [ ] Create REST endpoint documentation for the API

## Contact
**Mathew Estafanous -** mathewestafanous13@gmail.com

**Personal Website -** https://mathewestafanous.com/

**Project Link -** https://github.com/Mathew-Estafanous/Ur-Codebin
