package handlers

import (
    "encoding/json"
    "net/http"
    "taskflow/database"
    "taskflow/models"
    
    "golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var req models.RegisterRequest
    json.NewDecoder(r.Body).Decode(&req)

    // Hash password
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

    // Insertar usuario
    var userID int
    err := database.DB.QueryRow(
        "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id",
        req.Username, req.Email, string(hashedPassword),
    ).Scan(&userID)

    if err != nil {
        http.Error(w, "Error al crear usuario", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Usuario creado exitosamente",
        "user_id": userID,
    })
}

func Login(w http.ResponseWriter, r *http.Request) {
    var req models.LoginRequest
    json.NewDecoder(r.Body).Decode(&req)

    var user models.User
    var passwordHash string
    err := database.DB.QueryRow(
        "SELECT id, username, email, password_hash FROM users WHERE email = $1",
        req.Email,
    ).Scan(&user.ID, &user.Username, &user.Email, &passwordHash)

    if err != nil {
        http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
        return
    }

    // Verificar contraseña
    err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
    if err != nil {
        http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Login exitoso",
        "user":    user,
    })
}