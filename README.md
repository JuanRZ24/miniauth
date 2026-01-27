# MiniAuth

MiniAuth es una API de autenticación modular escrita en Go, enfocada en buenas prácticas de seguridad, claridad arquitectónica y reutilización.

## Objetivos

MiniAuth fue creado como un ejercicio práctico para reforzar fundamentos de seguridad en backend y como base de autenticación reutilizable para sistemas futuros.


## Features
- Registros de usuariosaa
- Login basico con email y contraseña
- Hash de contraseñas (bcrypt)
- Validaciones basicas
- Arquitectura por capas

## Roadmap

### Autenticación
- [ ] JWT con rotación
- [ ] Refresh tokens
- [ ] Verificación de email
- [ ] Recuperación de contraseña
- [ ] Roles y permisos

### Seguridad
- [ ] Rate limiting en login
- [ ] Bloqueo por intentos fallidos


## Arquitectura

MiniAuth sigue una arquitectura en capas:

- Handler: manejo de HTTP y validación de requests
- Service: lógica de negocio
- Repository: acceso a datos
- Models: entidades del dominio

## Stack

- Go
- Gin
- PostgreSQL
- bcrypt

## Requisitos

- Go 1.22+
- PostgreSQL

## Endpoints

POST /auth/register  
POST /auth/login  
****

## Seguridad

- Las contraseñas nunca se almacenan en texto plano
- Se usa bcrypt para hashing
- Los errores de login no revelan si el email existe

