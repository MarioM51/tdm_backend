# Users API

- [Users API](#users-api)
  - [Introduccion](#introduccion)
    - [Api Usuarios = Auth](#api-usuarios--auth)
    - [Api Productos](#api-productos)

## Introduccion

### Api Usuarios = Auth

1, Consultar todos los usuarios, get all

```r
curl http://localhost:8081/users \
    --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyMTQ0ODMsImp0aSI6Ijc5In0.jfiR0soCJat9zn1d1grNERc3_BW39vtz41b9qDoxyq4" \
    | json_pp
```

2, Agregar un usuario

```r
curl http://localhost:8081/users \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{ "email": "mario1@email.com", "password": "mario1" }'
```

3, Obtener usuario por ID

```r
curl http://localhost:8081/users/80 \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMyMjQ0ODIsImp0aSI6Ijc5In0.1tLAnb-bsj7uJ0YREcNZoMf6MxXezvC5JGfggn9HxzI" \
  | json_pp
```

4, Modificar un usuarrio

```r
curl http://localhost:8081/users \
    --include \
    --request "PUT" \
    --header "Content-Type: application/json" \
    --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMxNTk0MTUsImp0aSI6Ijc5In0.ro68MWf-Nki08rPhhIxAT6CRdhuXmA-pov4pvWkDApY" \
    --data '{ "id": 2, "email": "Train6" }'
```

5, Eliminar un usuarrio

```r
curl http://localhost:8081/users/80 --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMyMjQ0ODIsImp0aSI6Ijc5In0.1tLAnb-bsj7uJ0YREcNZoMf6MxXezvC5JGfggn9HxzI"
```

6, Activar usuario

```r
curl http://localhost:8081/users/1/activate/XXX
```

7, Login

```r
curl http://localhost:8081/users/login \
  --include \
  --header "Content-Type: application/json" \
  --request "POST" \
  --data '{ "email": "mario2@email.com", "password": "mario2" }'
```

### Api Productos

1, get All products

```r
curl http://localhost:8081/products \
  --request "GET" | json_pp
```

2, add a new product

- id product will be ignored

```r
curl http://localhost:8081/products \
  --request "POST" \
  --include \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyNDM0NjEsImp0aSI6IjgyIn0.zTMVlrAwRMpaKtXqUu1-foFwqXaWdvYNlU8C05VLCHY" \
  --data '{ "name":"Producto 3", "price":1003, "image":"Some.3", "description":"Some product description 3" }'
```

2, update a product

- In the payload the id product is required

```r
curl http://localhost:8081/products \
  --request "PUT" \
  --include \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyNDM0NjEsImp0aSI6IjgyIn0.zTMVlrAwRMpaKtXqUu1-foFwqXaWdvYNlU8C05VLCHY" \
  --data '{ "id": 3, "name":"Producto 33", "price":200.33, "image":"Some.33", "description":"Some product description 3" }'
```

4, delete a product by id

```r
curl http://localhost:8081/products/3 \
  --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyNDM0NjEsImp0aSI6IjgyIn0.zTMVlrAwRMpaKtXqUu1-foFwqXaWdvYNlU8C05VLCHY" \
  --include
```
