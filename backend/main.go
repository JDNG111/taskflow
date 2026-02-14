package main

import (
    "log"
    "net/http"
    "taskflow/database"
    "taskflow/handlers"
    "taskflow/websocket"
    
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    // Conectar a la base de datos
    database.Connect()
    defer database.Close()

    // Inicializar WebSocket Hub
    hub := websocket.NewHub()
    go hub.Run()

    // Configurar rutas
    router := mux.NewRouter()
    
    // API Routes
    api := router.PathPrefix("/api").Subrouter()
    api.HandleFunc("/register", handlers.Register).Methods("POST")
    api.HandleFunc("/login", handlers.Login).Methods("POST")
    api.HandleFunc("/projects", handlers.GetProjects).Methods("GET")
    api.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
    api.HandleFunc("/projects/{id}/tasks", handlers.GetTasks).Methods("GET")
    api.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
    api.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
    api.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
api.HandleFunc("/stats", handlers.GetStats).Methods("GET")
    
    // WebSocket
    router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        websocket.ServeWs(hub, w, r)
    })

    // Configurar CORS - PERMITIR SOLO EL PUERTO DEL FRONTEND (3000)
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"}, // ⚠️ SOLO 3000, no 3001
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Content-Type", "Authorization"},
    })

    handler := c.Handler(router)

    log.Println("🚀 Servidor backend corriendo en http://localhost:8081") // ⚠️ PUERTO 8081
    log.Fatal(http.ListenAndServe(":8081", handler)) // ⚠️ PUERTO 8081
}