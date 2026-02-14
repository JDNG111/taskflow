package handlers

import (
    "encoding/json"
    "net/http"
    "taskflow/database"
    "taskflow/models"
    
    "github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    projectID := vars["id"]

    rows, err := database.DB.Query(`
        SELECT id, title, description, status, priority, category, 
               COALESCE(assigned_to, 0) as assigned_to, 
               project_id, created_by, created_at 
        FROM tasks 
        WHERE project_id = $1
        ORDER BY 
            CASE priority 
                WHEN 'alta' THEN 1 
                WHEN 'media' THEN 2 
                WHEN 'baja' THEN 3 
            END, created_at DESC
    `, projectID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var t models.Task
        var assignedTo int
        rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Priority, &t.Category, 
                  &assignedTo, &t.ProjectID, &t.CreatedBy, &t.CreatedAt)
        t.AssignedTo = assignedTo
        tasks = append(tasks, t)
    }

    json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)

    var taskID int
    err := database.DB.QueryRow(
        `INSERT INTO tasks (title, description, status, priority, category, project_id, created_by) 
         VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
        task.Title, task.Description, task.Status, task.Priority, task.Category, 
        task.ProjectID, task.CreatedBy,
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

    // Si assigned_to es 0, poner NULL en la base de datos
    var assignedTo interface{}
    if task.AssignedTo == 0 {
        assignedTo = nil
    } else {
        assignedTo = task.AssignedTo
    }

    _, err := database.DB.Exec(
        `UPDATE tasks SET 
            title=$1, description=$2, status=$3, priority=$4, 
            category=$5, assigned_to=$6 
         WHERE id=$7`,
        task.Title, task.Description, task.Status, task.Priority, 
        task.Category, assignedTo, taskID,
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

func GetStats(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("user_id")
    
    var totalTasks, completedTasks, pendingTasks, inProgressTasks int
    
    // Total de tareas
    database.DB.QueryRow("SELECT COUNT(*) FROM tasks WHERE created_by = $1", userID).Scan(&totalTasks)
    
    // Completadas
    database.DB.QueryRow("SELECT COUNT(*) FROM tasks WHERE created_by = $1 AND status = 'completed'", userID).Scan(&completedTasks)
    
    // Pendientes
    database.DB.QueryRow("SELECT COUNT(*) FROM tasks WHERE created_by = $1 AND status = 'pending'", userID).Scan(&pendingTasks)
    
    // En progreso
    database.DB.QueryRow("SELECT COUNT(*) FROM tasks WHERE created_by = $1 AND status = 'in_progress'", userID).Scan(&inProgressTasks)
    
    // Por categoría
    var learnCount, practiceCount, projectCount, docCount int
    database.DB.QueryRow("SELECT COUNT(*) FROM tasks WHERE created_by = $1 AND category = 'aprender'", userID).Scan(&learnCount)
    database.DB.QueryRow("SELECT COUNT(*) FROM tasks WHERE created_by = $1 AND category = 'practicar'", userID).Scan(&practiceCount)
    database.DB.QueryRow("SELECT COUNT(*) FROM tasks WHERE created_by = $1 AND category = 'proyecto'", userID).Scan(&projectCount)
    database.DB.QueryRow("SELECT COUNT(*) FROM tasks WHERE created_by = $1 AND category = 'documentar'", userID).Scan(&docCount)
    
    stats := map[string]interface{}{
        "total": totalTasks,
        "completed": completedTasks,
        "pending": pendingTasks,
        "in_progress": inProgressTasks,
        "by_category": map[string]int{
            "aprender": learnCount,
            "practicar": practiceCount,
            "proyecto": projectCount,
            "documentar": docCount,
        },
        "progress_percentage": float64(completedTasks) / float64(totalTasks) * 100,
    }
    
    json.NewEncoder(w).Encode(stats)
}	