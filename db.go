package main

import (
	"bufio"
	"database/sql"
	"errors"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Creates a local sqlite database file "database.sqlite" in the root directory
// if it does not already exist, and starts a database connection
func connectToDB() *sql.DB {
	// check if db file exists, and if necessary, create it
	if _, err := os.Stat("./database.sqlite"); err != nil {
		os.Create("./database.sqlite")
	}

	// open connection to db file
	db, err := sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		panic(err)
	}

	return db
}

// Reads the queries from migrate.sql line by line, and executes each
// query against the database connected to be connectToBB() to intialize
// any required tables
func Migrate() error {
	//open db connection
	db := connectToDB()

	//read queries from "migrate.sql" line by line
	file, err := os.Open("./migrate.sql")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// execute each sql query in file
		queryText := scanner.Text()

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer tx.Rollback()

		stmt, err := tx.Prepare(queryText)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec()
		if err != nil {
			log.Fatal(err)
			return err
		}

		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// use: call "waypointObject.create()" and it registers
// the waypoint based on its type values in the database
//
// the parameter (Waypoint) just before the function name is
// called a "receiver". it requires you call this method
// on a struct. it's like an object oriented design pattern
//
// notice that we're passing pointers (by reference) in functions that modify the struct,
// and passing by value in cases where we are not modifying the struct.
// consider this a "security feature"
func (wp *Waypoint) Create() error {
	// IMPORTANT NOTE: the ID is assigned by the database serialization:
	// the Waypoint passed as the receiver MUST HAVE an ID with some sentinel
	// value, let's say -1 (impossible ID), and then be assigned an ID
	// by the database upon being created in the database.
	// throw an error if an ID is passed that is not -1, since there
	// is an obvious issue. have a distinct type for this error
	// and describe it as something like "non-sentinel ID passed to Waypoint.create()"

	db := connectToDB()

	id := wp.ID
	name := wp.Name
	longitude := wp.Longitude
	latitude := wp.Latitude
	altitude := wp.Altitude

	if id != -1 {
		return errors.New(
			"non-sentinel ID passed to Waypoint.Create()" +
				"\n Expected ID: -1 but got ID: " + strconv.Itoa(id))
	}

	query := `INSERT INTO Waypoints (name, longitude, latitude, altitude) VALUES ($1, $2, $3, $4)`

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, longitude, latitude, altitude)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return err
	}

	wp.Get()

	return nil
}

// Update the database record for a given Waypoint based on its ID.
// The current values held in the fields of this Waypoint struct will
// be copied over to their corresponding columns in the database
func (wp Waypoint) Update() error {
	// using the receiver, update the waypoint's values in the database based on its ID

	db := connectToDB()

	id := wp.ID
	name := wp.Name
	longitude := wp.Longitude
	latitude := wp.Latitude
	altitude := wp.Altitude

	if id == -1 {
		return errors.New(
			"sentinel ID (-1) passed to Waypoint.Update(). " +
				"Did you remember to call Waypoint.Create() beforehand?")
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	query := `UPDATE Waypoints SET name = ($1), longitude = ($2), latitude = ($3), altitude = ($4) WHERE id = ($5)`

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, longitude, latitude, altitude, id)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Delete the database record for a given Waypoint based on its ID.
func (wp Waypoint) Delete() error {
	// using the receiver, delete the waypoint from the database based on its ID

	db := connectToDB()

	id := wp.ID

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer tx.Rollback()

	query := `DELETE FROM Waypoints WHERE id = ($1)`

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Writes the values stored in the database for this given Waypoint's ID
// to the Waypoint object.
// If this method is called on a Waypoint that has not yet been registed in the database,
// (i.e. Waypoint.Create() has not been called), then the method queries the database
// for records that match all four fields of (name, longitude, latitude, altitude).
// Otherwise, this method queries the database for records that match only the ID.
func (wp *Waypoint) Get() error {
	// the opposite of Waypoint.create()!
	// instead, populate all fields from the database
	// based on the ID provided

	db := connectToDB()

	id := wp.ID

	var row *sql.Row

	if id == -1 {
		//do based off fields
		query := `SELECT * FROM Waypoints WHERE 
			name = ($1) AND 
			longitude = ($2) AND
			latitude = ($3) AND
			altitude = ($4)`

		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer stmt.Close()

		row = stmt.QueryRow(wp.Name, wp.Longitude, wp.Latitude, wp.Altitude)
		if err != nil {
			log.Fatal(err)
			return err
		}

	} else {
		//do based off ID

		query := `SELECT * FROM Waypoints WHERE id = ($1)`

		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer stmt.Close()

		row = stmt.QueryRow(wp.ID)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	err := row.Scan(&wp.ID, &wp.Name, &wp.Longitude, &wp.Latitude, &wp.Altitude)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Definitions for AEACRoutes DB methods
// Follow the same general procedure as the Waypoints methods
func (AEACRoutes) Create() {

}

func (AEACRoutes) Update() {

}

func (AEACRoutes) Delete() {

}

func (*AEACRoutes) Get() {

}
