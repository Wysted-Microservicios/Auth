# Template Go-SQL Events

## Instalación

### Docker

Tiene dos Dockerfile, uno para producción y otro para entorno de desarrollo.

`Dockerfile.dev`
`Dockerfile.prod`

Ambos exponen el puerto: `8080`

### Docker compose

Correr:

```bash
docker compose up app
```

Con esto se levanta el servicio con base de datos y NATS que es el message broker que se usa para la comunicación entre micro-servicios.

> [!WARNING]
> Se va a crear todos los STREAMS de NATS en desarrollo, pero en producción debe ejecutar un script o similar que cree los STREAMS manualmente.

## API Reference (Swagger)

Docs: https://github.com/swaggo/swag

Para la documentación, primero instalar:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

#### Index

```http
  GET /swagger/index.html
```

#### docs.json

```http
  GET /swagger/docs.json
```

Comandos para actualizar documentación luego de haber hecho los comentarios:

```bash
swag fmt
swag init -o src/cmd/http/docs
```

## Base de datos

Las modificaciones a la base de datos se colocan en la carpeta `db/prisma/schema.prisma`, para producción se tiene que ocupar otro método para realizar las migraciones de la base de datos.

Luego, para tener los modelos en código Go, instale.

```bash
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
go install github.com/glerchundi/sqlboiler-crdb/v4@latest
```

Para hacer migraciones y tener los modelos en codigo Go.

```bash
./scripts/updateDb.bash
```

Esto creará en `src/package/db/models` todos los modelos de la base de datos en local en código Go compilado con todos los métodos y tests hechos.

## Funcionalidades

-   📝 Registrarse: Permite a los nuevos usuarios crear una cuenta.
-   🔐 Iniciar sesión: Acceso al sistema para usuarios registrados.
-   🔄 Cambiar contraseña: Opción para actualizar la contraseña actual del usuario.
-   🚀 Compatibilidad con NATS: Integración para comunicación eficiente mediante el sistema de mensajería NATS.
-   🧪 Testing a todos los servicios y sus funciones
