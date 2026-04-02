# Groupie Tracker

A **Go web application** that uses the Groupie Tracker API to display musician/band data, tour dates, and locations.

---

## Features

* Fetches artists, relations, locations, and dates from `https://groupietrackers.herokuapp.com/api/`.
* Merges data into a unified model: artists include `DatesLocations`, `Dates`, and `Locations`.
* Renders an HTML page with band profiles via templates.
- **Custom styling for improved visualization and readability of band data**.
* Serves static assets (media and CSS) for the web UI.
* Simple local server on `http://localhost:8080`.

---

## Usage

```bash
go run main.go
```

Open:

```
http://localhost:8080
```

---

## Project Structure

```
.
├── go.mod
├── helpers
│   ├── handlers.go
│   └── parsing.go
├── main.go
├── media
│   └── undo.png
├── README.md
└── templates
    ├── errors.html
    ├── index.html
    └── style.css
```

---