# go-api
Simple Go API for User Management. Launches on http://localhost:10000. To launch, run command `go run main.go`.
## Routes
Simple routes
- View All Users - http://localhost:10000/users (GET)
- Create New User - http://localhost:10000/users (POST)
- View Single User - http://localhost:10000/users/{id} (GET)
- View Random User - http://localhost:10000/random-user (GET)
- Introduce User - http://localhost:10000/users/{id}/hello (GET)
- Edit Single User - http://localhost:10000/users/{id} (PUT)
- Delete Single User - http://localhost:10000/users/{id} (DELETE)
- Delete All Users - http://localhost:10000/users/ (DELETE)

### User Data Example
{
    "id": 1,
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