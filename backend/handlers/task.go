package handlers

import (
    "encoding/json"
    "net/http"
    // "strconv"      // Comenta si no se usa
    "taskflow/database"
    "taskflow/models"
    
    "github.com/gorilla/mux"  // Este SÍ se usa en GetTasks
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    projectID := vars["id"]

    rows, err := database.DB.Query(`
        SELECT id, title, description, status, assigned_to, project_id, created_by, created_at 
        FROM tasks 
        WHERE project_id = $1
        ORDER BY created_at DESC
    `, projectID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var t models.Task
        rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.AssignedTo, &t.ProjectID, &t.CreatedBy, &t.CreatedAt)
        tasks = append(tasks, t)
    }

    json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)

    var taskID int
    err := database.DB.QueryRow(
        "INSERT INTO tasks (title, description, project_id, created_by) VALUES ($1, $2, $3, $4) RETURNING id",
        task.Title, task.Description, task.ProjectID, task.CreatedBy,
    ).Scan(&taskID)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    task.ID = taskID
    json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    taskID := vars["id"]

    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)

    _, err := database.DB.Exec(
        "UPDATE tasks SET title=$1, description=$2, status=$3, assigned_to=$4 WHERE id=$5",
        task.Title, task.Description, task.Status, task.AssignedTo, taskID,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Tarea actualizada"})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    taskID := vars["id"]

    _, err := database.DB.Exec("DELETE FROM tasks WHERE id = $1", taskID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Tarea eliminada"})
}