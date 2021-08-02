# Task Management Application

## Brief explanation

Create the **back-end** (REST API) part of the task management application (like [Trello](https://www.youtube.com/watch?v=noguPYxyv6g)) with no authentication and authorization required.

The main entity is a _Project_ (or _Board_) that always has its name and contains multiple _Columns_. _A Column_ always has a name, contains _Tasks_ and represents their status.

When a _Project_ is created at least one "default" _Column_ must be created in it as well.

A _Task_ can be created only inside the _Column_ and can be moved within the _Column_ (change priority) or across the _Columns_ (change status).

A _Task_ can have _Comments_ that could contain questions or _Task_ clarification information.

## Technical Requirements

The development could be divided into 3 phases: *Services*, *Handlers*, *DB*. Please, prefix each comment with the appropriate phase title like "Services: Add basic routes".

It is not required to develop the application phase by phase, but keep in mind that you have to add all these parts to the application.
Think about the architecture and build the basement for a future application extension.

### Phase 1 (Services)

Create the core functionality of the application. Define the service level of the architecture that takes care of business logic. Keep in mind that you&#39;ll need to add handlers and DB layer.

### Phase 2 (Handlers)

Extend your existing Core functionality with the HTTP server and handlers to be able to handle HTTP requests. Add the documentation for it.

### Phase 3 (DB)

Extend your application with the database layer: save and fetch the application data from the database.

## Technical Limitations

1. Any external dependency that is not specified in this document or is not compatible with the requirements must be discussed with technical mentors in the school chat. After approval, it will be added to this document.
2. Any dependency management tool is allowed (go modules are preferable).
3. Handlers must accept requests and send responses in JSON format.
4. Any external router compatible with standard with net/http server and handlers is allowed (e.g. gorilla/mux, go-chi/chi).
5. Documentation for the HTTP API must be present. It could be written with any tool (like Swagger) or in the Readme/Documentation file(s) in the repository.
6. All errors must be logged. Any other information could be logged (up to you).
7. Any external logger is allowed (we use uber/zap).
8. You must use _database/sql_ with any DB driver to work with the database (PostgreSQL is preferable, but is not required).
9. Tests must be present: unit tests and request tests (using net/http/httptest package).
10. Any dependencies for testing purposes are allowed.
11. Instruction on how to start-up the application must be present.
12. Pull Request on any platform (GitHub, GitLab, ...) must be opened.
13. The application must be deployed to any cloud provider: AWS, GCP, Azure, Heroku, etc. (Heroku is free and suits well for small projects). Please, attach the corresponding link in README file.

**Suggestions**

You can use go-playground/validator ([https://github.com/go-playground/validator](https://github.com/go-playground/validator)) to validate an HTTP request or anything else if needed.

**Allowed libraries**

- godotenv
- github.com/Yalantis/go-config
- github.com/spf13/viper
- github.com/golang-migrate/migrate
- github.com/pressly/goose
- github.com/rs/cors

## Business Rules

_User_ must be able to manage (create, read, update, delete) _Projects_:

- _Projects_ in a list are sorted by name.

_Project_ must contain at least one column:

- the first column created by default when a _Project_ created;
- the last column cannot be deleted.

_User_ must be able to manage (create, read, update, delete) _Columns_:

- _Columns_ in a list are sorted by their position specified by _User_;
- _Column_ name **must** be unique;
- when a _Column_ is deleted its tasks are moved to the _Column_ to the left of the current.

_User_ must be able to move a _Column_ left or right.

_User_ must be able to rename a _Column_

_User_ must be able to manage (create, read, update, delete) _Tasks_:

- _Task_ can be created only within a _Column_;
- User can view _Tasks_ in all _Columns_ of a _Project_;
- _User_ can update the name and the description of the _Task_;
- _User_ can delete the _Task_, with all _Comments_ related to this _Task_.

_User_ must be able to move a _Task_ across the _Columns_ (left or right) to change its status.

_User_ must be able to move a _Task_ within the _Column_ (up and down) to prioritize it.

_User_ must be able to manage (create, read, update, delete) _Comments_:

- _Comment_ can be created only within a _Task_;
- _Comments_ in a list are sorted by their creation date (from newest to oldest);
- _User_ can view _Comments_ related to a _Task_;
- _User_ can update the _Comment_ text;
- _User_ can delete the _Comment_.

## Fields and validation rules

**Project**

- Name (1-500 characters)
- Description (0-1000 characters)

**Column**

- Name/Status (1-255 characters, **unique** )

**Task**

- Name (1-500 characters)
- Description (0-5000 characters)

**Comment**

- Text (1-5000 characters)
