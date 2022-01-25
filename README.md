# Users API

## Introduccion

1, Consultar todos los usuarios

```r
curl http://localhost:8080/users | json_pp
```

2, Agregar un usuario

```r
curl http://localhost:8080/users \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{ "email": "mario1@mail.com", "password": "mario1" }'
```

3, Obtener usuario por ID

```r
curl http://localhost:8080/users/1
```

4, Modificar un usuarrio

```r
curl http://localhost:8080/users \
    --include \
    --request "PUT" \
    --header "Content-Type: application/json" \
    --header "token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMwNzU5NzQsImlzcyI6IjQifQ.KZQ71teq3y5mm8bb7hPBZJxmZX99o7bp2EsL4pAkE5Q" \
    --data '{ "id": 1, "email": "Train2" }'
```

5, Eliminar un usuarrio

```r
curl http://localhost:8080/users/2 --request "DELETE"
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
    --data '{ "email": "mario3@mail.com", "password": "mario3" }'
```
