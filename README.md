# Getting Started
This project relies on the [Echo](https://echo.labstack.com/guide/) framework. Read a bit more about it if you aren't familiar.

This project uses go 1.18. Details about the packages it relies on are in the `go.mod` file.

To run the project for development, simply run:
```
go run main.go controllers.go odlc.go pathfinding.go
```

To build a binary for the project, run:
```
go build
```


When running, the root of the project is available at `localhost:1323`.

For testing endpoints for development, it's highly recommended that you use an API client. [Insomnia](https://insomnia.rest/) is a great choice.

# File Structure
## main.go
Controllers are paired routes and HTTP request methods. Middleware, logging, and everything else should be configured here.

## controllers.go
Controllers are declared here.

## odlc.go
Any functions related to interfacing with ODLC are declared here.

## pathfinding.go
Any functions related to interfacing with pathfinding are declared here.

## rcomms.go
Any functions related to interfacing with the RCOMMS system are declared here. Most likely, this include instantiation of socket connections via a middleware to allow for endpoints in the form of controllers to be reached. 

# Tests
Tests are ran on a per file basis. For example, tests for `odlc.go` will be in `odlc_test.go`.

Tests can be ran with `go test <test files to run> <files being tested>`.
Alternatively, `go test` will automatically test all files in the project.

Be sure to write a test for every function you make, barring some exceptions like `main`.

There are options for coverage and verbosity as part of `go test`.


# Database
The database is a SQLite based database generated by shell scripts.

## Tables
*format: Tablename (Corresponding Go Struct)*
### waypoints (Waypoint)
- id: serial, primary key, internal
- name: text, identifier, arbitrary and provided by the competition
- longitude (degrees)
- latitude (degrees)
- altitude (meters above sea level, nullable)

### aeac_routes (AEACRoute)
- id: serial, primary key, internal
- number: integer, provided by the competition, used to identify routes to the judges
- start_waypoint_name: name of the starting waypoint
- end_waypoint_name:   name of the ending waypoint
- passengers: number of passengers being transported
- max_vehicle_weight: maximum weight of the vehicle allowed for the route (kg or lbs? unsure yet. just a number for now)
- value: $ value of route, floating point value
- remarks: text, indicates any additional details about the parameters of the route
- order: cardinal number indicating the nth order we are taking this route in. nullable. if null, we are not taking this route.
