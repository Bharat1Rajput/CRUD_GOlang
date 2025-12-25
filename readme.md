# Go CRUD API with Gorilla Mux

A simple REST API built with Go and Gorilla Mux for managing items.

## Features

- Create, Read, Update, Delete (CRUD) operations
- In-memory storage using slices
- JSON request/response handling
- RESTful API endpoints

## Prerequisites

- Go 1.21 or higher
- Gorilla Mux package

## Installation

1. Clone the repository
2. Install dependencies:
```bash
go get -u github.com/gorilla/mux
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

Server will start on `http://localhost:8080`

## API Endpoints

### 1. Get All Items
```bash
GET http://localhost:8080/items
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "Apple"
  }
]
```

### 2. Create Item
```bash
POST http://localhost:8080/items
Content-Type: application/json

{
  "name": "Apple"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "Apple"
}
```

### 3. Get Single Item
```bash
GET http://localhost:8080/items/1
```

**Response:**
```json
{
  "id": 1,
  "name": "Apple"
}
```

### 4. Update Item
```bash
PUT http://localhost:8080/items/1
Content-Type: application/json

{
  "name": "Orange"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "Orange"
}
```

### 5. Delete Item
```bash
DELETE http://localhost:8080/items/1
```

**Response:**
```json
{
  "message": "item deleted successfully"
}
```

## Testing with cURL

```bash
# Get all items
curl http://localhost:8080/items

# Create item
curl -X POST http://localhost:8080/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Apple"}'

# Get single item
curl http://localhost:8080/items/1

# Update item
curl -X PUT http://localhost:8080/items/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Orange"}'

# Delete item
curl -X DELETE http://localhost:8080/items/1
```

## Testing with Thunder Client / Postman

1. Set the correct HTTP method (GET, POST, PUT, DELETE)
2. Use URL: `http://localhost:8080/items` (or `/items/{id}`)
3. For POST/PUT: Add header `Content-Type: application/json`
4. For POST/PUT: Add JSON body in the Body tab

## Common Mistakes to Avoid

### ⚠️ Trailing Slash Issue
- ✅ **Correct:** `http://localhost:8080/items`
- ❌ **Wrong:** `http://localhost:8080/items/`

**The trailing slash causes a 404 error!** Gorilla Mux treats `/items` and `/items/` as different routes.

**Solution:** Either remove the trailing slash OR add `r.StrictSlash(true)` to your router:
```go
r := mux.NewRouter()
r.StrictSlash(true)  // Auto-handles trailing slashes
```

### ⚠️ Missing Content-Type Header
Always include `Content-Type: application/json` header for POST and PUT requests, otherwise the server may not parse the JSON body correctly.

### ⚠️ Using http:// not https://
Use `http://localhost:8080` not `https://localhost:8080` (unless you've configured SSL).


## Project Structure

```
.
├── main.go          # Main application file
├── go.mod           # Go module file
└── README.md        # This file
```

## Technologies Used

- **Go** - Programming language
- **Gorilla Mux** - HTTP router and URL matcher
- **JSON** - Data format for API communication

## License

MIT License