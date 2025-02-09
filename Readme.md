# 2DCH7-20059: Pars Tasmim - Code Challenge
This repository contains the implementation of the Pars Tasmim code challenge.

## Project Structure
The project is organized as follows:

- contract/: Contains interface definitions and contracts used across the application.
- delivery/httpserver/: Houses the HTTP server implementation and related handlers.
- docs/swagger/: Includes Swagger documentation for the API endpoints.
- dto/: Data Transfer Objects used for communication between layers.
- entity/: Defines the core entities of the application.
- mocks/: Contains mock implementations for testing purposes.
- pkg/: Utility packages and helper functions.
- repository/: Data access layer managing interactions with the database.
- service/: Business logic and service implementations.
- validator/: Validates API inputs before passing them to next layers.

### Entities
#### User
Defines each user that want to use the reservation system.
#### Table
Identifies each table that exists in the restaurant.
#### Reservation
Used for storing reservations a user make.
#### Receipt
*NOT IMPLEMENTED*: After making a reservation, user needs to pay the price to make it definite.

### Services
#### User
- Register: Helps users to register themselves in the system
- Login: To login in the system.
- Logout: To logout from the system.
- IsTokenValid: Checks if the provided JWT token is valid.

#### Table
- Add: Adds a new table to restaurant. Each table can have seats between 4 and 10. Also, the number of the tables cannot exceed 10.
- Remove: Removes a table.
- Timetable: Helps the user to get the timetable of the restaurant at the specified date.

#### Reservation
- Book: Helps the user to book for a reservations.
- Cancel: Cancels a reservation. A reservation can only be canceled if there is more than three hours time to it.


## Getting Started
### Installation
Clone the repository:
```bash
git clone https://github.com/techerfan/2DCH7-20059.git
cd 2DCH7-20059
```
Set up environment variables:

Create a .env file in the root directory and configure the necessary environment variables.

### Start the services:

Use Docker Compose to build and start the services:

```bash
docker compose up --build
```
This will set up the application along with its dependencies, such as the PostgreSQL database, Redis, and Adminer panel.

### Running the Application
Once the services are up, the application will be accessible at `http://localhost:{PORT}`.

* Tables are not initialized. You can add a new table by the specified route in the documenation. Also, you are free to delete each table.
* You cannot add a table with the same number that is added before.

## API Documentation
Swagger documentation is available at `http://localhost:{PORT}/swagger/index.html`.

## TODO
- [ ] Add logger
- [ ] Write tests
- [ ] Implement locks for database
- [x] Add constraints