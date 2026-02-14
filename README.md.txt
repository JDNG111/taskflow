# TaskFlow - Gestión de Proyectos Colaborativa 🚀

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![React](https://img.shields.io/badge/React-18-61DAFB?style=for-the-badge&logo=react)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-316192?style=for-the-badge&logo=postgresql)
![WebSocket](https://img.shields.io/badge/WebSocket-Real%20Time-010101?style=for-the-badge)

## 📋 Descripción

**TaskFlow** es una aplicación full-stack de gestión de proyectos y tareas, construida con **Go** en el backend y **React** en el frontend. Permite a los usuarios crear proyectos, asignar tareas y colaborar en tiempo real con un tablero estilo Kanban.

## ✨ Características

- ✅ **Autenticación de usuarios** - Registro y login seguro
- ✅ **Gestión de proyectos** - Crear y visualizar proyectos
- ✅ **Tablero Kanban** - Organiza tareas en columnas (Pendientes, En Progreso, Completadas)
- ✅ **Tiempo real** - Actualizaciones en vivo con WebSockets
- ✅ **API RESTful** - Backend robusto en Go
- ✅ **Base de datos relacional** - PostgreSQL con relaciones complejas

## 🛠️ Tecnologías Utilizadas

### Backend
- **Go** - Lenguaje principal
- **Gorilla Mux** - Router HTTP
- **Gorilla WebSocket** - Comunicación en tiempo real
- **PostgreSQL** - Base de datos
- **bcrypt** - Encriptación de contraseñas
- **CORS** - Middleware para seguridad

### Frontend
- **React 18** - Librería UI
- **React Router DOM** - Navegación
- **Axios** - Cliente HTTP
- **Hooks** - useState, useEffect
- **Context API** - Manejo de estado

### Base de Datos
- **PostgreSQL 15+**
- **pgAdmin 4** - Administración visual

## 📁 Estructura del Proyecto
taskflow/
├── backend/ # API en Go
│ ├── database/ # Conexión a PostgreSQL
│ ├── handlers/ # Manejadores HTTP
│ ├── models/ # Estructuras de datos
│ ├── websocket/ # Comunicación en tiempo real
│ └── main.go # Punto de entrada
├── frontend/ # Aplicación React
│ ├── public/ # Archivos estáticos
│ └── src/ # Código fuente
│ ├── components/# Componentes React
│ └── services/ # Servicios API
└── README.md



## 🚀 Instalación y Ejecución

### Prerrequisitos
- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Git

### 1. Clonar el repositorio
```bash
git clone https://github.com/JDNG111/taskflow.git
cd taskflow

### 2. Configurar Base de Datos
CREATE DATABASE taskflow;
Ejecuta el script SQL en database.sql (incluido en el repositorio)

### 3. Backend
cd backend
go mod download
go run main.go
# Servidor en http://localhost:8081

### 4. Frontend
cd frontend
npm install
npm start
# Aplicación en http://localhost:3000

🔮 Próximas Mejoras
Autenticación JWT

Asignación de tareas a usuarios

Comentarios en tareas

Notificaciones en tiempo real

Modo oscuro

Pruebas unitarias

👨‍💻 Autor
Julian Navarro - @JDNG111

📄 Licencia
Este proyecto está bajo la Licencia MIT - ver el archivo LICENSE para más detalles.

🙏 Agradecimientos
A la comunidad de Go y React

A todos los que prueben la aplicación

⭐️ ¡Si te gusta este proyecto, no olvides darle una estrella en GitHub! ⭐️
