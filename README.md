# Users API

## Introduccion

1, Consultar todos los usuarios, get all

```r
curl http://localhost:8080/users \
    --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMyMjQ0ODIsImp0aSI6Ijc5In0.1tLAnb-bsj7uJ0YREcNZoMf6MxXezvC5JGfggn9HxzI" \
    | json_pp
```

2, Agregar un usuario

```r
curl http://localhost:8080/users \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{ "email": "mario1@email.com", "password": "mario1" }'
```

3, Obtener usuario por ID

```r
curl http://localhost:8080/users/80 \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMyMjQ0ODIsImp0aSI6Ijc5In0.1tLAnb-bsj7uJ0YREcNZoMf6MxXezvC5JGfggn9HxzI" \
  | json_pp
```

4, Modificar un usuarrio

```r
curl http://localhost:8080/users \
    --include \
    --request "PUT" \
    --header "Content-Type: application/json" \
    --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMxNTk0MTUsImp0aSI6Ijc5In0.ro68MWf-Nki08rPhhIxAT6CRdhuXmA-pov4pvWkDApY" \
    --data '{ "id": 2, "email": "Train6" }'
```

5, Eliminar un usuarrio

```r
curl http://localhost:8080/users/80 --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMyMjQ0ODIsImp0aSI6Ijc5In0.1tLAnb-bsj7uJ0YREcNZoMf6MxXezvC5JGfggn9HxzI"
```

6, Activar usuario

```r
curl http://localhost:8080/users/1/activate/XXX
```

7, Login

```r
curl http://localhost:8080/users/login \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{ "email": "mario2@email.com", "password": "mario2" }'
```
