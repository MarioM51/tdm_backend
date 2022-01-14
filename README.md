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
    --data '{ "id": 1, "email": "mario1@mail.com", "password": "mario1" }' | json_pp
```

3, Obtener usuario por ID

```r
curl http://localhost:8080/users/1
```

4, Modificar un usuarrio

```r
curl http://localhost:8080/users \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{ "id": 1, "email": "Train2", "password": "Coltrane2" }'
```

5, Eliminar un usuarrio

```r
curl http://localhost:8080/users/2 --request "DELETE"
```
