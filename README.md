# go-api
Simple Go API for User Management. Launches on http://localhost:3000. To launch, run command `go run main.go` from project root. Requires PostgreSQL server running. Check example env for settings. 
## Routes
Simple routes
    - View All Users - http://localhost:3000/users</a> (GET)
    - Create New User - http://localhost:3000/users (POST)
    - View Single User - http://localhost:3000/users/{id}</a> (GET)
    - View Random User - http://localhost:3000/random-user</a> (GET)
    - Introduce User - http://localhost:3000/users/{id}/hello</a> (GET)
    - Edit Single User - http://localhost:3000/users/{id} (PUT)
    - Delete Single User - http://localhost:3000/users/{id} (DELETE)
    - Delete All Users - http://localhost:3000/users/ (DELETE)

### User Data Example
{
    "fname": "John",
    "lname": "Smith",
    "pname": "JJ",
    "email": "john.smith@example.com",
    "skills": [
        "PHP",
        "GoLang",
        "JS"
    ],
    "experience": 15,
    "since": "2001-02-22T00:00:00Z"
}