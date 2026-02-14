package handlers

import (
    "encoding/json"
    "net/http"
    "taskflow/database"
    "taskflow/models"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query(`
        SELECT id, name, description, category, owner_id, created_at 
        FROM projects 
        ORDER BY created_at DESC
    `)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var projects []models.Project
    for rows.Next() {
        var p models.Project
        rows.Scan(&p.ID, &p.Name, &p.Description, &p.Category, &p.OwnerID, &p.CreatedAt)
        projects = append(projects, p)
    }

    json.NewEncoder(w).Encode(projects)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
    var project models.Project
    json.NewDecoder(r.Body).Decode(&project)

    var projectID int
    err := database.DB.QueryRow(
        "INSERT INTO projects (name, description, category, owner_id) VALUES ($1, $2, $3, $4) RETURNING id",
        project.Name, project.Description, project.Category, project.OwnerID,
    ).Scan(&projectID)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    project.ID = projectID
    json.NewEncoder(w).Encode(project)
}