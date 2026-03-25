# Groupie Tracker

A **Go web application** that uses the Groupie Tracker API to display musician/band data, tour dates, and locations.

---

## Features

* Fetches artists, relations, locations, and dates from `https://groupietrackers.herokuapp.com/api/`.
* Merges data into a unified model: artists include `DatesLocations`, `Dates`, and `Locations`.
* Renders an HTML page with band profiles via templates.
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
groupie-tracker/
├─ go.mod
├─ main.go
├─ media/               # Static media assets
├─ template/
│  ├─ index.html
│  └─ style.css
└─ utils/
   ├─ parse.go          # API parsing and data-shaping logic
   └─ utils.go          # template rendering helper
```

---