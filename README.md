# MiniAuth - API de Autenticación Modular en Go

**MiniAuth** es un servicio de autenticación robusto y modular escrito en Go. El proyecto está diseñado bajo principios de arquitectura limpia y seguridad, sirviendo como una base sólida para sistemas de gestión de usuarios en aplicaciones modernas.

## Objetivos del Proyecto

MiniAuth fue desarrollado para ofrecer un sistema de autenticación reutilizable, escalable y fácil de integrar. Se enfoca en la separación de responsabilidades y el uso de estándares de la industria para el manejo de credenciales.

## Características Implementadas

- **Gestión de Usuarios**: Registro y autenticación segura.
- **Seguridad de Contraseñas**: Uso intensivo de `bcrypt` para el hashing de credenciales.
- **Autenticación Basada en JWT**: Generación y validación de tokens JSON Web Tokens.
- **Arquitectura en Capas**: Estructura modular que separa Handlers, Servicios y Repositorios.
- **Migraciones Automáticas**: Control de versiones de base de datos con scripts SQL integrados.
- **Verificación de Email**: Soporte inicial para flujos de validación de identidad.

## Stack Tecnológico

- **Lenguaje**: Go 1.22+
- **Framework Web**: Gin Gonic.
- **Base de Datos**: PostgreSQL.
- **Seguridad**: JWT (JSON Web Token) y Bcrypt.

## Estructura del Proyecto

- `cmd/server/`: Punto de entrada de la aplicación.
- `internal/handlers/`: Controladores de rutas HTTP.
- `internal/services/`: Lógica de negocio de autenticación.
- `internal/repositories/`: Capa de persistencia (PostgreSQL).
- `internal/db/migrations/`: Scripts SQL para la estructura de la base de datos.
- `internal/security/`: Utilidades para JWT y hashing.

## Configuración y Despliegue

### Requisitos
- Go instalado.
- Servidor PostgreSQL activo.

### Instalación
1.  Clona el repositorio.
2.  Copia el archivo de ejemplo de variables de entorno:
    ```bash
    cp .env.example .env
    ```
3.  Instala las dependencias:
    ```bash
    go mod tidy
    ```
4.  Inicia el servidor:
    ```bash
    go run cmd/server/main.go
    ```

## Endpoints de la API

- `POST /auth/register`: Crea una nueva cuenta de usuario.
- `POST /auth/login`: Autentica al usuario y devuelve un token JWT.
- `GET /auth/verify`: Valida el token de sesión actual (Protegido).

## Roadmap

- [ ] Implementar rotación de Refresh Tokens.
- [ ] Integrar Rate Limiting para prevenir ataques de fuerza bruta.
- [ ] Soporte para roles y permisos (RBAC).

## Seguridad
MiniAuth prioriza la privacidad. Las contraseñas nunca se almacenan en texto plano y los errores de autenticación están diseñados para no revelar información sensible sobre la existencia de cuentas.

## Licencia
Distribuido bajo la licencia MIT.

