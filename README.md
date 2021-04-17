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
[![schema](https://user-images.githubusercontent.com/56979977/114913465-37922000-9def-11eb-9093-250cce0c6e86.png)](https://dbdiagram.io/d/606262f8ecb54e10c33dd900)

## TODO
The project is still under development and is continually changing so here are a few things underworks
- [x] Moving from MySQL to PostgreSQL.
- [ ] Refactor the database schema to a better format.
- [ ] Create REST endpoint documentation for the API
- [ ] Host the API on a cloud platform (AWS or Heroku)

## Contact
**Mathew Estafanous -** mathewestafanous13@gmail.com

**Personal Website -** https://mathewestafanous.com/

**Project Link -** https://github.com/Mathew-Estafanous/Ur-Codebin
