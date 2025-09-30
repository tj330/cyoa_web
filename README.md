# cyoa_web

**cyoa_web** is a web-based application that brings "Choose Your Own Adventure" (CYOA) stories to life. Built in Go, this project features a modular, organized backend and a simple frontend so users can experience branching narratives in their browser.

## Features

- **Interactive Story Navigation:** Make choices to explore different story paths.
- **Extensible Story Format:** Add or modify stories easily (JSON).
- **RESTful API:** Story data can be served and manipulated via HTTP endpoints.
- **Separation of Concerns:** Clear structure for handlers, models, routes, utilities, and static assets.
- **Simple Frontend:** Minimal HTML/CSS for quick deployment and customization.

## Project Structure

```
cyoa_web/
├── cyoa                      # Executable binary
├── go.mod, go.sum            # Go dependencies
├── internal/
│   ├── api/
│   │   └── story_handler.go  # API handlers for story endpoints
│   ├── app/
│   │   └── app.go            # App startup and orchestration
│   ├── data/
│   │   └── story.json        # Example story data (JSON)
│   ├── models/
│   │   └── story.go          # Story model definitions
│   ├── routes/
│   │   └── routes.go         # Route definitions and setup
│   ├── static/
│   │   └── styles.css        # CSS styles
│   ├── templates/
│   │   └── index.html        # Main HTML template
│   └── utils/
│       └── utils.go          # Utility functions
└── main.go                   # Application entry point

```

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18+ recommended)

### Installation & Running

1. **Clone the repository:**
   ```sh
   git clone https://github.com/tj330/cyoa_web.git
   cd cyoa_web
   ```

2. **Run the app:**
   ```sh
   go run main.go
   ```

3. **Open your browser:**
   ```
   http://localhost:8080/
   ```

### Adding or Editing Stories

- Place your story files in `internal/data/`, following the `story.json` format.
- See `internal/models/story.go` for the expected structure and fields.

### API

- Story endpoints are managed in `internal/api/story_handler.go`.
- Routing is configured in `internal/routes/routes.go`.
- To extend the API, add new handlers or routes as needed.

## Development Notes

- **Frontend:** HTML/CSS lives in `internal/templates/` and `internal/static/`.
- **Backend:** All business logic and models are under `internal/`.
