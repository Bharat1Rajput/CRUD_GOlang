# Go CRUD API with MongoDB and Gorilla Mux

A REST API built with Go, MongoDB, and Gorilla Mux for managing users with persistent database storage.

## Features

- Create, Read, Update, Delete (CRUD) operations
- MongoDB integration for persistent storage
- JSON request/response handling
- RESTful API endpoints
- Proper error handling and context management

## Prerequisites

- Go 1.21 or higher
- MongoDB (running on localhost:27017)
- Required Go packages

## Installation

1. Clone the repository
2. Install dependencies:
```bash
go get -u github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson
go mod tidy
```

3. Make sure MongoDB is running:
```bash
# Start MongoDB
mongod
```

4. Run the server:
```bash
go run server/main.go
```

Server will start on `http://localhost:8080`

## API Endpoints

### 1. Get All Users
```bash
GET http://localhost:8080/users
```

**Response:**
```json
[
  {
    "id": "507f1f77bcf86cd799439011",
    "name": "Bharat Singh",
    "email": "Bharat@123.com",
    
  }
]
```

### 2. Create User
```bash
POST http://localhost:8080/user
Content-Type: application/json

{
  "name": "Bharat Singh",
  "email": "Bharat@123.com"
}
```

**Response:**
```json
{
  "id": "507f1f77bcf86cd799439011",
  "name": "Bharat Singh",
  "email": "Bharat@123.com"
}
```

### 3. Get Single User
```bash
GET http://localhost:8080/user/507f1f77bcf86cd799439011
```

**Response:**
```json
{
  "id": "507f1f77bcf86cd799439011",
  "name": "Bharat Singh",
  "email": "Bharat@123.com"
}
```

### 4. Update User
```bash
PUT http://localhost:8080/user/507f1f77bcf86cd799439011
Content-Type: application/json

{
  "name": "Bharat Singh",
  "email": "Bharat@123.com"
}
```

**Response:**
```json
{
  "id": "507f1f77bcf86cd799439011",
  "name": "Bharat Singh",
  "email": "Bharat@123.com"
}
```

### 5. Delete User
```bash
DELETE http://localhost:8080/users/507f1f77bcf86cd799439011
```

**Response:**
```json
{
  "message": "User deleted successfully"
}
```

## Testing with cURL

```bash
# Get all users
curl http://localhost:8080/users

# Create user
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name": "Bharat Singh","email": "Bharat@123.com"}'

# Get single user (replace ID with actual MongoDB ObjectID)
curl http://localhost:8080/user/507f1f77bcf86cd799439011

# Update user
curl -X PUT http://localhost:8080/user/507f1f77bcf86cd799439011 \
  -H "Content-Type: application/json" \
  -d '{"name": "Bharat Singh","email": "Bharat@123.com"}'

# Delete user
curl -X DELETE http://localhost:8080/user/507f1f77bcf86cd799439011
```

## Testing with Thunder Client / Postman

1. Set the correct HTTP method (GET, POST, PUT, DELETE)
2. Use URL: `http://localhost:8080/users` (or `/user/{id}`)
3. For POST/PUT: Add header `Content-Type: application/json`
4. For POST/PUT: Add JSON body in the Body tab
5. Use actual MongoDB ObjectID for {id} parameter (24-character hex string)

## Common Mistakes to Avoid

### ⚠️ Trailing Slash Issue
- ✅ **Correct:** `http://localhost:8080/users`
- ❌ **Wrong:** `http://localhost:8080/users/`

**The trailing slash causes a 404 error!** Gorilla Mux treats `/users` and `/users/` as different routes.

**Solution:** Either remove the trailing slash OR add `r.StrictSlash(true)` to your router:
```go
r := mux.NewRouter()
r.StrictSlash(true)  // Auto-handles trailing slashes
```

### ⚠️ Missing Content-Type Header
Always include `Content-Type: application/json` header for POST and PUT requests, otherwise the server may not parse the JSON body correctly.

### ⚠️ Using http:// not https://
Use `http://localhost:8080` not `https://localhost:8080` (unless you've configured SSL).

### ⚠️ Exposing Server Errors to Client
**Bad Practice:**
```go
if err != nil {
    json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})  // ❌ Exposes internal errors
}
```

**Good Practice:**
```go
if err != nil {
    log.Println("Database error:", err)  // Log internally
    json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch user"})  // ✅ Generic message to client
}
```

**Why?** Sending `err.Error()` to clients can expose:
- Database connection strings
- Internal file paths
- Security vulnerabilities
- Sensitive implementation details

### ⚠️ Not Returning User ID After Update
**Bad Practice:**
```go
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    // ... update logic ...
    json.NewEncoder(w).Encode(user)  // ❌ user.ID is empty!
}
```

**Good Practice:**
```go
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    // ... update logic ...
    user.ID = id  // ✅ Set the ID from URL params
    json.NewEncoder(w).Encode(user)
}
```

### ⚠️ Using log Inside Error Handler
**Bad Practice:**
```go
if err != nil {
    log.Fatal(err)  // ❌ Terminates entire server!
}
```

**Good Practice:**
```go
if err != nil {
    log.Println("Error:", err)  // ✅ Just logs, doesn't crash server
    // Handle error gracefully
}
```

**Why?** `log.Fatal()` calls `os.Exit(1)` which shuts down your entire server. Use it only in initialization code (like `main()`), not in request handlers.

### ⚠️ Not Using DisconnectMongoDB()
**Important:** Go requires explicit resource management. Unlike Node.js which automatically cleans up connections, Go needs you to explicitly close database connections.

**Bad Practice:**
```go
func main() {
    database.ConnectMongoDB()
    // ❌ Connection never closed - resource leak!
    http.ListenAndServe(":8080", router)
}
```

**Good Practice:**
```go
func main() {
    database.ConnectMongoDB()
    defer database.DisconnectMongoDB()  // ✅ Ensures cleanup even if panic occurs
    http.ListenAndServe(":8080", router)
}
```

**Why?** 
- Prevents connection leaks
- Ensures graceful shutdown
- Releases database resources properly
- `defer` guarantees execution even if panic occurs

### ⚠️ Invalid MongoDB ObjectID
MongoDB uses 24-character hexadecimal ObjectIDs. Using invalid IDs will cause errors:
- ✅ **Valid:** `507f1f77bcf86cd799439011`
- ❌ **Invalid:** `1`, `abc`, `12345`

## Project Structure

```
.
├── server/
│   └── main.go           # Main application entry point
├── models/
│   └── user.go          # User model with bson tags
├── handlers/
│   └── user.go          # User CRUD handlers
├── config/
│   └── db.go            # MongoDB connection setup
├── go.mod               # Go module file
├── go.sum               # Dependency checksums
└── README.md            # This file
```

## Model Structure

The User model uses MongoDB BSON tags and primitive.ObjectID:

```go
type User struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name      string             `json:"name" bson:"name"`
    Email     string             `json:"email" bson:"email"`
}
```

**Key Points:**
- `primitive.ObjectID` is MongoDB's native ID type (not int or string)
- `bson` tags map Go fields to MongoDB document fields
- `json` tags control JSON serialization
- `omitempty` allows MongoDB to auto-generate IDs

## Technologies Used

- **Go** - Programming language
- **Gorilla Mux** - HTTP router and URL matcher
- **MongoDB** - NoSQL database for data persistence
- **MongoDB Go Driver** - Official MongoDB driver for Go
- **BSON** - Binary JSON for MongoDB documents
- **JSON** - Data format for API communication

## Environment Configuration

For production, use environment variables for MongoDB connection:

```go
// config/db.go
mongoURI := os.Getenv("MONGODB_URI")
if mongoURI == "" {
    mongoURI = "mongodb://localhost:27017" // fallback
}
```

## 👨‍💻 Author

**Bharat Rajput**
- GitHub: [@Bharat1Rajput](https://github.com/Bharat1Rajput)
- LinkedIn: [Bharat Singh](https://www.linkedin.com/in/bharat-singh-1288a4254)
- Email: bharattsingh33@gmail.com



## License

MIT License
