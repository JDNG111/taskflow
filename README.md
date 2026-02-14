<div align="center">

# 🚀 SkillFlow

### *Gestor Inteligente de Aprendizaje y Proyectos*

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![React Version](https://img.shields.io/badge/React-18.2-61DAFB?style=for-the-badge&logo=react&logoColor=white)](https://react.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-336791?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)

[Características](#-características-principales) •
[Demo](#-demo) •
[Instalación](#-instalación-rápida) •
[Tecnologías](#-stack-tecnológico) •
[Roadmap](#-roadmap) •
[Contribuir](#-cómo-contribuir)

<img src="https://via.placeholder.com/800x400/667eea/ffffff?text=SkillFlow+Dashboard" alt="SkillFlow Dashboard" width="100%"/>

</div>

---

## 📖 Sobre el Proyecto

**SkillFlow** es una aplicación full-stack moderna que transforma la manera en que organizas tu aprendizaje técnico. Diseñada especialmente para desarrolladores y estudiantes autodidactas, combina la metodología Kanban con un sistema inteligente de seguimiento de progreso.

### 💡 ¿Por qué SkillFlow?

- 🎯 **Enfoque Visual**: Visualiza tu progreso de aprendizaje en un tablero Kanban intuitivo
- 📊 **Datos Accionables**: Estadísticas en tiempo real sobre tu productividad
- 🔥 **Stack Moderno**: Go + React + PostgreSQL para máximo rendimiento
- 🚀 **Escalable**: Arquitectura preparada para miles de usuarios concurrentes
- 💻 **100% Open Source**: Código abierto y gratuito para siempre

---

## ✨ Características Principales

<table>
<tr>
<td width="50%">

### 🗂️ Gestión de Proyectos
- Crea proyectos de aprendizaje ilimitados
- Organiza por categorías personalizadas
- Descripción detallada con Markdown

### 📊 Tablero Kanban
- Vista de 3 columnas: Pendiente → En Progreso → Completada
- Drag & drop para mover tareas (próximamente)
- Colores por prioridad

</td>
<td width="50%">

### 🏷️ Sistema de Etiquetas
- **Categorías**: Aprender, Practicar, Proyecto, Documentar
- **Prioridades**: Alta 🔴, Media 🟡, Baja 🟢
- Filtros avanzados

### 📈 Analytics
- Gráficos de progreso por proyecto
- Distribución de tareas por categoría
- Historial de completados

</td>
</tr>
</table>

### 🔐 Seguridad y Autenticación

- Sistema de registro y login
- Contraseñas hasheadas con bcrypt
- Sesiones seguras
- Datos privados por usuario

---

## 🎬 Demo

### 📸 Capturas de Pantalla

Creación de Actividades
<img width="1889" height="910" alt="image" src="https://github.com/user-attachments/assets/61f484e8-8cea-44d2-821f-2868f31de79b" />

Tablero
<img width="1879" height="895" alt="image" src="https://github.com/user-attachments/assets/0237a9d3-290b-4c2a-8f60-5e54e238ac11" />

Estadisticas
<img width="1901" height="895" alt="image" src="https://github.com/user-attachments/assets/e77c841c-7725-4fe4-95bf-e1a3a594bb9a" />

---

## 🛠️ Stack Tecnológico

<table>
<tr>
<td align="center" width="33%">

### Backend
![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white)

- **Lenguaje**: Go 1.21+
- **Router**: Gorilla Mux
- **WebSockets**: Gorilla WebSocket
- **Auth**: bcrypt
- **CORS**: rs/cors

</td>
<td align="center" width="33%">

### Frontend
![React](https://img.shields.io/badge/React-61DAFB?style=flat-square&logo=react&logoColor=black)

- **Framework**: React 18.2
- **UI**: Bootstrap 5
- **Gráficos**: Chart.js
- **HTTP**: Axios
- **Routing**: React Router v6

</td>
<td align="center" width="33%">

### Base de Datos
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=flat-square&logo=postgresql&logoColor=white)

- **DBMS**: PostgreSQL 15
- **Driver**: lib/pq
- **Índices**: Optimizados
- **Relaciones**: Foreign Keys

</td>
</tr>
</table>

### 🏗️ Arquitectura
```
skillflow/
├── backend/
│   ├── main.go                 # Punto de entrada
│   ├── handlers/               # Controladores HTTP
│   │   ├── auth.go
│   │   ├── projects.go
│   │   └── tasks.go
│   ├── models/                 # Estructuras de datos
│   ├── database/               # Conexión DB
│   └── websocket/              # Manejo WebSockets
│
├── frontend/
│   ├── src/
│   │   ├── components/         # Componentes React
│   │   │   ├── Dashboard.jsx
│   │   │   ├── KanbanBoard.jsx
│   │   │   └── TaskCard.jsx
│   │   ├── context/            # Context API
│   │   ├── services/           # API calls
│   │   └── App.js
│   └── public/
│
└── db/
    └── schema.sql              # Esquema de base de datos
```

---

## 🚀 Instalación Rápida

### Prerrequisitos
```bash
# Verifica versiones
go version      # >= 1.21
node --version  # >= 18.0
psql --version  # >= 15.0
```

### 📦 Paso 1: Clonar Repositorio
```bash
git clone https://github.com/JDNG111/taskflow.git
cd skillflow
```

### 🗄️ Paso 2: Configurar Base de Datos
```bash
# Crear base de datos
psql -U postgres
CREATE DATABASE skillflow;
\q

# Ejecutar migraciones
psql -U postgres -d skillflow -f db/schema.sql
```

### ⚙️ Paso 3: Backend
```bash
cd backend

# Descargar dependencias
go mod download

# Configurar credenciales
# Editar database/connection.go con tus datos:
# - Host: localhost
# - User: tu_usuario
# - Password: tu_contraseña
# - Database: skillflow

# Iniciar servidor (puerto 8080)
go run main.go
```

### 🎨 Paso 4: Frontend
```bash
# En otra terminal
cd frontend

# Instalar dependencias
npm install

# Iniciar aplicación (puerto 3000)
npm start
```

### ✅ Paso 5: ¡Listo!

Abre tu navegador en `http://localhost:3000` 🎉

---

## 📚 Uso

### Crear un Proyecto
```javascript
// POST /api/projects
{
  "name": "Aprender Go",
  "description": "Dominar Goroutines y Channels",
  "created_by": 1
}
```

### Agregar Tarea
```javascript
// POST /api/tasks
{
  "title": "Leer documentación de Concurrency",
  "description": "Capítulos 1-3",
  "priority": "alta",
  "category": "aprender",
  "status": "pendiente",
  "project_id": 1
}
```

### Actualizar Estado
```javascript
// PUT /api/tasks/5
{
  "status": "completada"
}
```

---

## 🎯 Roadmap

### En Progreso 🔨
- [ ] Drag & Drop en tablero Kanban
- [ ] Filtros avanzados de tareas
- [ ] Modo oscuro 🌙

### Próximas Features 🔜
- [ ] Compartir proyectos con otros usuarios
- [ ] Notificaciones por email
- [ ] Exportar datos a PDF/CSV
- [ ] API REST documentada (Swagger)
- [ ] Aplicación móvil nativa

### Largo Plazo 🚀
- [ ] Dockerización completa
- [ ] Tests automatizados (80%+ coverage)
- [ ] CI/CD con GitHub Actions
- [ ] Integración con GitHub para importar repos
- [ ] IA para sugerencias de aprendizaje

---

## 🤝 Cómo Contribuir

¡Las contribuciones son lo que hace que la comunidad open source sea increíble! Cualquier contribución es **muy apreciada**.

### Proceso

1. **Fork** el proyecto
2. Crea tu **Feature Branch** (`git checkout -b feature/NuevaCaracteristica`)
3. **Commit** tus cambios (`git commit -m 'Add: Nueva característica increíble'`)
4. **Push** al Branch (`git push origin feature/NuevaCaracteristica`)
5. Abre un **Pull Request**

### Guía de Estilo

- Go: Sigue [Effective Go](https://golang.org/doc/effective_go)
- React: Usa Hooks, evita clases
- Commits: Formato [Conventional Commits](https://www.conventionalcommits.org/)

---

## 🐛 Reportar Bugs

¿Encontraste un bug? [Abre un issue aquí](https://github.com/tuusuario/skillflow/issues/new?template=bug_report.md)

---

## 💬 Comunidad y Soporte

- 📧 Email: navarroestudiante1010@gmail.com

---

## 📄 Licencia

Distribuido bajo la Licencia MIT. Ver [`LICENSE`](LICENSE) para más información.
```
MIT License

Copyright (c) 2024 Tu Nombre

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software...
```

---

## 🙏 Agradecimientos

- [Gorilla Toolkit](https://www.gorillatoolkit.org/) - Excelente toolkit para Go
- [React Bootstrap](https://react-bootstrap.github.io/) - Componentes UI
- [Chart.js](https://www.chartjs.org/) - Gráficos interactivos
- [PostgreSQL](https://www.postgresql.org/) - Base de datos confiable
- Comunidad de desarrolladores que inspiran este proyecto

---

## 📊 Estadísticas del Proyecto

![GitHub stars](https://img.shields.io/github/stars/tuusuario/skillflow?style=social)
![GitHub forks](https://img.shields.io/github/forks/tuusuario/skillflow?style=social)
![GitHub watchers](https://img.shields.io/github/watchers/tuusuario/skillflow?style=social)

![GitHub issues](https://img.shields.io/github/issues/tuusuario/skillflow)
![GitHub pull requests](https://img.shields.io/github/issues-pr/tuusuario/skillflow)
![GitHub last commit](https://img.shields.io/github/last-commit/tuusuario/skillflow)

---

<div align="center">

### ⭐ Si te gusta este proyecto, dale una estrella en GitHub ⭐

Hecho con ❤️ por [Tu Nombre](https://github.com/JDNG111)

[⬆ Volver arriba](#-skillflow)

</div>
