package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// func TestHello(t *testing.T) {
// 	e := echo.New()

// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, Hello(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, "Hello, World!", rec.Body.String())
// 	}
// }

func TestPostWaypoints(t *testing.T) {
	cleanDB()

	e := echo.New()

	wp1 := Waypoint{
		ID:        -1,
		Name:      "gcomWP1",
		Longitude: -6.2,
		Latitude:  149.1651830,
		Altitude:  20.0,
	}

	wp2 := Waypoint{
		ID:        -1,
		Name:      "gcomWP2",
		Longitude: -36.3637798,
		Latitude:  147.1651830,
		Altitude:  20.0,
	}

	wp3 := Waypoint{
		ID:        -1,
		Name:      "gcomWP3",
		Longitude: -37.3637798,
		Latitude:  146.1641830,
		Altitude:  10.0,
	}

	wpslice := []Waypoint{wp1, wp2, wp3}

	jsonBody, _ := json.Marshal(wpslice)

	req := httptest.NewRequest(http.MethodPost, "/waypoints", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, PostWaypoints(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		successMsg := JSONResponse{
			Type: "Message",
			Message: "Waypoints successfully registered!",
		}
		var returnedMsg JSONResponse
		json.Unmarshal([]byte(rec.Body.Bytes()), &returnedMsg)
		assert.Equal(t, successMsg, returnedMsg)
	}

	cleanDB()
}

func TestGetWaypoints(t *testing.T) {
	cleanDB()

	e := echo.New()

	wp1 := Waypoint{
		ID:        -1,
		Name:      "gcomWP1",
		Longitude: -6.2,
		Latitude:  149.1651830,
		Altitude:  20.0,
	}

	wp2 := Waypoint{
		ID:        -1,
		Name:      "gcomWP2",
		Longitude: -36.3637798,
		Latitude:  147.1651830,
		Altitude:  20.0,
	}

	wpslice := []Waypoint{wp1, wp2}

	jsonBody, _ := json.Marshal(wpslice)

	req := httptest.NewRequest(http.MethodPost, "/waypoints", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, PostWaypoints(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		successMsg := JSONResponse{
			Type: "Message",
			Message: "Waypoints successfully registered!",
		}
		var returnedMsg JSONResponse
		json.Unmarshal([]byte(rec.Body.Bytes()), &returnedMsg)
		assert.Equal(t, successMsg, returnedMsg)

		//check if get returns all the waypoints

		req = httptest.NewRequest(http.MethodGet, "/waypoints", nil)
		rec := httptest.NewRecorder()

		c = e.NewContext(req, rec)

		if assert.NoError(t, GetWaypoints(c)) {
			wp1.ID = 1
			wp2.ID = 2
			wpslice = []Waypoint{wp1, wp2}

			jsonBody, _ = json.Marshal(wpslice)

			// Info.Println(string(jsonBody))

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, string(jsonBody)+"\n", rec.Body.String())
		}
	}

	cleanDB()
}

func TestPostRoutes(t *testing.T) {
	cleanDB()

	e := echo.New()

	r1 := AEACRoutes{
		ID:               -1,
		Number:           1,
		StartWaypoint:    "alpha",
		EndWaypoint:      "delta",
		Passengers:       5,
		MaxVehicleWeight: 30.5,
		Value:            13.2,
		Remarks:          "",
		Order:            0,
	}

	r2 := AEACRoutes{
		ID:               -1,
		Number:           2,
		StartWaypoint:    "delta",
		EndWaypoint:      "zeta",
		Passengers:       10,
		MaxVehicleWeight: 20.05,
		Value:            17.3,
		Remarks:          "remark",
		Order:            1,
	}

	rSlice := []AEACRoutes{r1, r2}

	jsonBody, _ := json.Marshal(rSlice)

	req := httptest.NewRequest(http.MethodPost, "/routes", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, PostRoutes(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		successMsg := JSONResponse{
			Type: "Message",
			Message: "AEACRoutes registered!",
		}
		var returnedMsg JSONResponse
		json.Unmarshal([]byte(rec.Body.Bytes()), &returnedMsg)
		assert.Equal(t, successMsg, returnedMsg)
	}

	cleanDB()
}

func TestGetRoutes(t *testing.T) {
	cleanDB()

	e := echo.New()

	r1 := AEACRoutes{
		ID:               -1,
		Number:           1,
		StartWaypoint:    "alpha",
		EndWaypoint:      "delta",
		Passengers:       5,
		MaxVehicleWeight: 30.5,
		Value:            13.2,
		Remarks:          "",
		Order:            0,
	}

	r2 := AEACRoutes{
		ID:               -1,
		Number:           2,
		StartWaypoint:    "delta",
		EndWaypoint:      "zeta",
		Passengers:       10,
		MaxVehicleWeight: 20.05,
		Value:            17.3,
		Remarks:          "remark",
		Order:            1,
	}

	rSlice := []AEACRoutes{r1, r2}

	jsonBody, _ := json.Marshal(rSlice)

	req := httptest.NewRequest(http.MethodPost, "/routes", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, PostRoutes(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		successMsg := JSONResponse{
			Type: "Message",
			Message: "AEACRoutes registered!",
		}
		var returnedMsg JSONResponse
		json.Unmarshal([]byte(rec.Body.Bytes()), &returnedMsg)
		assert.Equal(t, successMsg, returnedMsg)

		//check if get returns all the routes

		req = httptest.NewRequest(http.MethodGet, "/routes", nil)
		rec := httptest.NewRecorder()

		c = e.NewContext(req, rec)

		if assert.NoError(t, GetRoutes(c)) {
			r1.ID = 1
			r2.ID = 2
			routeSlice := []AEACRoutes{r1, r2}

			jsonBody, _ = json.Marshal(routeSlice)

			// Info.Println(string(jsonBody))

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, string(jsonBody)+"\n", rec.Body.String())
		}
	}

	cleanDB()
}

func TestGetNextRoute(t *testing.T) {
	cleanDB()

	e := echo.New()

	r1 := AEACRoutes{
		ID:               -1,
		Number:           1,
		StartWaypoint:    "alpha",
		EndWaypoint:      "delta",
		Passengers:       5,
		MaxVehicleWeight: 30.5,
		Value:            13.2,
		Remarks:          "",
		Order:            5,
	}

	r2 := AEACRoutes{
		ID:               -1,
		Number:           2,
		StartWaypoint:    "delta",
		EndWaypoint:      "zeta",
		Passengers:       10,
		MaxVehicleWeight: 20.05,
		Value:            17.3,
		Remarks:          "remark",
		Order:            3,
	}

	r3 := AEACRoutes{
		ID:               -1,
		Number:           3,
		StartWaypoint:    "delta",
		EndWaypoint:      "alpha",
		Passengers:       10,
		MaxVehicleWeight: 20.05,
		Value:            17.3,
		Remarks:          "next waypoint",
		Order:            0,
	}

	rSlice := []AEACRoutes{r1, r2, r3}

	jsonBody, _ := json.Marshal(rSlice)

	req := httptest.NewRequest(http.MethodPost, "/routes", strings.NewReader(string(jsonBody)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, PostRoutes(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		successMsg := JSONResponse{
			Type: "Message",
			Message: "AEACRoutes registered!",
		}
		var returnedMsg JSONResponse
		json.Unmarshal([]byte(rec.Body.Bytes()), &returnedMsg)
		assert.Equal(t, successMsg, returnedMsg)

		//check that getNextRoute returns the route with the lowest order

		req = httptest.NewRequest(http.MethodGet, "/nextroute", nil)
		rec := httptest.NewRecorder()

		c = e.NewContext(req, rec)

		if assert.NoError(t, GetNextRoute(c)) {
			r3.ID = 3
			// routeSlice := []AEACRoutes{r1, r2}

			jsonBody, _ = json.Marshal(r3)

			// Info.Println(string(jsonBody))

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, string(jsonBody)+"\n", rec.Body.String())
		}
	}

	cleanDB()
}
