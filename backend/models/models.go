package models

import "time"

type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"password,omitempty"`
    CreatedAt time.Time `json:"created_at"`
}

type Project struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Category    string    `json:"category"`
    OwnerID     int       `json:"owner_id"`
    CreatedAt   time.Time `json:"created_at"`
}

type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    Priority    string    `json:"priority"`
    Category    string    `json:"category"`
    AssignedTo  int       `json:"assigned_to"`
    ProjectID   int       `json:"project_id"`
    CreatedBy   int       `json:"created_by"`
    CreatedAt   time.Time `json:"created_at"`
    DueDate     time.Time `json:"due_date"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type RegisterRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}