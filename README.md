# Tienda

- [Tienda](#tienda)
  - [SPA: Panel de administracion](#spa-panel-de-administracion)
  - [API](#api)
    - [Api Usuarios = Auth](#api-usuarios--auth)
    - [Api Productos](#api-productos)
    - [API blogs](#api-blogs)
  - [SSR (Server side render)](#ssr-server-side-render)
    - [Products](#products)
    - [Blogs](#blogs)

## SPA: Panel de administracion

Peticion con el navegador

```r
#Use in the browser
http://localhost:8081/api/admin
```

## API

### Api Usuarios = Auth

1, get all users

```r
curl http://localhost:8081/api/users \
    --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyMTQ0ODMsImp0aSI6Ijc5In0.jfiR0soCJat9zn1d1grNERc3_BW39vtz41b9qDoxyq4" \
    | json_pp
```

2, Add user

```r
curl http://localhost:8081/api/users \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{ "email": "mario1@email.com", "password": "mario1" }'
```

3, Get user by id

```r
curl http://localhost:8081/api/users/80 \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMyMjQ0ODIsImp0aSI6Ijc5In0.1tLAnb-bsj7uJ0YREcNZoMf6MxXezvC5JGfggn9HxzI" \
  | json_pp
```

4, Update user

```r
curl http://localhost:8081/api/users \
    --include \
    --request "PUT" \
    --header "Content-Type: application/json" \
    --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMxNTk0MTUsImp0aSI6Ijc5In0.ro68MWf-Nki08rPhhIxAT6CRdhuXmA-pov4pvWkDApY" \
    --data '{ "id": 2, "email": "Train6" }'
```

5, Delete user

```r
curl http://localhost:8081/api/users/80 --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMyMjQ0ODIsImp0aSI6Ijc5In0.1tLAnb-bsj7uJ0YREcNZoMf6MxXezvC5JGfggn9HxzI"
```

6, Activate user

```r
curl http://localhost:8081/api/users/1/activate/XXX
```

7, Login

```r
curl http://localhost:8081/api/users/login \
  --include \
  --header "Content-Type: application/json" \
  --request "POST" \
  --data '{ "email": "mario2@email.com", "password": "mario2" }'
```

### Api Productos

1, get All products

```r
curl http://localhost:8081/api/products \
  --request "GET" | json_pp
```

2, add a new product

- id product will be ignored

```r
curl http://localhost:8081/api/products \
  --request "POST" \
  --include \
  --header "Token: ..." \
  --data '{ "name":"Producto 3", "price":1003, "image":"Some.3", "description":"Some product description 3" }'
```

2, update a product

- In the payload the id product is required

```r
curl http://localhost:8081/api/products \
  --request "PUT" \
  --include \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyNDM0NjEsImp0aSI6IjgyIn0.zTMVlrAwRMpaKtXqUu1-foFwqXaWdvYNlU8C05VLCHY" \
  --data '{ "id": 3, "name":"Producto 33", "price":200.33, "image":"Some.33", "description":"Some product description 3" }'
```

4, delete a product by id

```r
curl http://localhost:8081/api/products/3 \
  --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyNDM0NjEsImp0aSI6IjgyIn0.zTMVlrAwRMpaKtXqUu1-foFwqXaWdvYNlU8C05VLCHY" \
  --include
```

### API blogs

1, get all blogs

```r
curl http://localhost:8081/api/blogs --request "GET" | json_pp
```

2, find by id

```r
curl http://localhost:8081/api/blogs/2 --request "GET" | json_pp
```

3, show image

```r
#Use in the browser
http://localhost:8081/api/blogs/1/image
```

4, add blog

- id blog will be ignored

```r
curl http://localhost:8081/api/blogs --request "POST" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTExMTY4MTYsImp0aSI6Ijc5In0.7mai9dtJhPEpEMyYBDcEDf_IJ2w0PcPj-JPbEhEPdZs" \
  --data '{ "id": 0, "title": "Algo 2", "body": "<XX>...</XX><p>...", "abstract": "Some abstract ...", "thumbnail": "data:image/png;base64,xx...==", "author": null, "createdAt": null, "updateAt": null }' \
  | json_pp
```

5, update blog

```r
curl http://localhost:8081/api/blogs --request "PUT" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTExMTY4MTYsImp0aSI6Ijc5In0.7mai9dtJhPEpEMyYBDcEDf_IJ2w0PcPj-JPbEhEPdZs" \
  --data '{ "id": 1, "title": "Algo 11", "body": "<XX>11</XX><p>...", "abstract": "Some abstract 11", "thumbnail": "data:image/png;base64,xx...11", "author": null, "createdAt": null, "updateAt": null }' \
  | json_pp
```

4, delete a blog by id

```r
curl http://localhost:8081/api/blogs/3 \
  --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTExMTY4MTYsImp0aSI6Ijc5In0.7mai9dtJhPEpEMyYBDcEDf_IJ2w0PcPj-JPbEhEPdZs" \
  | json_pp
```

## SSR (Server side render)

### Products

```r
#Use in the browser
http://localhost:8081/products
```

### Blogs

```r
#Use in the browser
http://localhost:8081/blogs
```
