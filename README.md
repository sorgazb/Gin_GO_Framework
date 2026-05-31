# Gin_GO_Framework

![Go](https://img.shields.io/badge/Go-1.24+-00add8?style=for-the-badge&logo=go&logoColor=white)&nbsp;![Gin](https://img.shields.io/badge/Gin-v1.10.0-00add8?style=for-the-badge&logo=go&logoColor=white)&nbsp;![REST API](https://img.shields.io/badge/REST-API%20JSON-6366f1?style=for-the-badge)&nbsp;![Estado](https://img.shields.io/badge/Estado-En%20progreso-22c55e?style=for-the-badge)

> **Gin_GO_Framework** es una aplicación web en **Go** construida con el framework **Gin v1.10**, que implementa enrutamiento con parámetros, respuestas en texto plano y una mini **API REST JSON** para la gestión de usuarios con validación de campos.

---

## 📋 Descripción

El proyecto sirve como base de aprendizaje para construir APIs y servidores web con Go y Gin. Incluye:

- **Enrutamiento dinámico** con parámetros de ruta (`/:nombre`).
- **API REST JSON** con endpoint `POST /usuarios` y binding automático de JSON.
- **Validación de campos** requeridos (`nombre` y `email`) con respuestas de error descriptivas.
- **Almacenamiento en memoria** de usuarios registrados durante la sesión.
- **Plantillas HTML** y **archivos estáticos** listos para ampliar.
- **Documentación interna** en `docs/` con notas sobre el servidor web básico en Go.

---

## 🛣️ Endpoints disponibles

| Método | Ruta | Descripción |
|--------|------|-------------|
| `GET` | `/` | Respuesta de texto: `Hola mundo` |
| `GET` | `/saludo/:nombre` | Saludo personalizado con el nombre |
| `POST` | `/usuarios` | Registra un usuario nuevo (JSON: `nombre`, `email`) |

**Ejemplo de petición POST:**
```json
{
  "nombre": "Sergio",
  "email": "sergio@example.com"
}
```

**Respuesta:**
```json
{
  "message": "Usuario Registrado",
  "datos": [
    { "nombre": "Sergio", "email": "sergio@example.com" }
  ]
}
```

---

## 🏗️ Estructura del Proyecto

```txt
Gin_GO_Framework/
├── routes/
│   └── routes.go          # Definición de rutas y lógica de handlers
├── static/                # Archivos estáticos (CSS, JS, imágenes)
├── templates/             # Plantillas HTML
├── docs/
│   └── Servidor Web Basico en Go.md   # Notas y documentación interna
├── main.go                # Punto de entrada: instancia Gin y arranca en :8080
├── go.mod                 # Módulo Go y dependencias
├── go.sum                 # Checksums de dependencias
└── README.md              # Documentación del proyecto
```

---

## ⚙️ Instalación y Ejecución

Clonar el repositorio:
```bash
git clone https://github.com/sorgazb/Gin_GO_Framework.git
cd Gin_GO_Framework
```

Descargar dependencias:
```bash
go mod tidy
```

Levantar el servidor:
```bash
go run main.go
```

El servidor estará disponible en:
```txt
http://localhost:8080
```

Probar el endpoint POST con curl:
```bash
curl -X POST http://localhost:8080/usuarios \
  -H "Content-Type: application/json" \
  -d '{"nombre": "Sergio", "email": "sergio@example.com"}'
```

---

## 🤝 Contribución

Haz fork del repositorio.

Crea una rama de trabajo:
```bash
git checkout -b feature/mi-nueva-funcionalidad
```

Realiza tus cambios y haz commit.

Abre un Pull Request describiendo tus mejoras.

---

<p align="center">Aprendizaje de Gin Framework &nbsp;·&nbsp; Sergio Orgaz Bravo</p>
