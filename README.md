# Tienda

- [Tienda](#tienda)
  - [TODO](#todo)
  - [SPA: Panel de administracion](#spa-panel-de-administracion)
  - [API](#api)
    - [Api Usuarios = Auth](#api-usuarios--auth)
    - [Api Productos](#api-productos)
      - [Product likes](#product-likes)
      - [Product image](#product-image)
      - [Product comments](#product-comments)
    - [API blogs](#api-blogs)
      - [Blog Images](#blog-images)
      - [Blog Likes](#blog-likes)
      - [Blog Comments](#blog-comments)
    - [API Orders](#api-orders)
  - [SSR (Server side render)](#ssr-server-side-render)
    - [Products](#products)
    - [Blogs](#blogs)

## TODO

1, Refactorizar para uso de commons mas adecuado
2, Refactorizar para uso de constantes cumunes

## SPA: Panel de administracion

Peticion con el navegador

```r
#Use in the browser
http://192.168.1.81:80/api/admin
```

## API

### Api Usuarios = Auth

1, get all users

```r
curl http://192.168.1.81:80/api/users \
    --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyMTQ0ODMsImp0aSI6Ijc5In0.jfiR0soCJat9zn1d1grNERc3_BW39vtz41b9qDoxyq4" \
    | json_pp
```

2, Add user

```r
curl http://192.168.1.81:80/api/users \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{ "email": "mario1@email.com", "password": "mario1" }'
```

3, Get user by id

```r
curl http://192.168.1.81:80/api/users/80 \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMyMjQ0ODIsImp0aSI6Ijc5In0.1tLAnb-bsj7uJ0YREcNZoMf6MxXezvC5JGfggn9HxzI" \
  | json_pp
```

4, Update user

```r
curl http://192.168.1.81:80/api/users \
    --include \
    --request "PUT" \
    --header "Content-Type: application/json" \
    --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMxNTk0MTUsImp0aSI6Ijc5In0.ro68MWf-Nki08rPhhIxAT6CRdhuXmA-pov4pvWkDApY" \
    --data '{ "id": 2, "email": "Train6" }'
```

5, Delete user

```r
curl http://192.168.1.81:80/api/users/80 --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMyMjQ0ODIsImp0aSI6Ijc5In0.1tLAnb-bsj7uJ0YREcNZoMf6MxXezvC5JGfggn9HxzI"
```

6, Activate user

```r
curl http://192.168.1.81:80/api/users/1/activate/XXX
```

7, Login

```r
curl http://192.168.1.81:80/api/users/login \
  --include \
  --header "Content-Type: application/json" \
  --request "POST" \
  --data '{ "email": "mario2@email.com", "password": "mario2" }'
```

### Api Productos

1, get All products

```r
curl http://192.168.1.81:80/api/products \
  --request "GET" | json_pp
```

2, add a new product

- id product will be ignored

```r
curl http://192.168.1.81:80/api/products \
  --request "POST" \
  --include \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTcxNzU4NjAsImp0aSI6Ijc5In0.F-PyYWdcBkbY-xIprWwFpH57tBjl2xRTuSgmI-F4S8s" \
  --data '{ "name":"Producto 1", "price":1001, "description":"Some product description 1" }'
```

3, update a product

- In the payload the id product is required

```r
curl http://192.168.1.81:80/api/products \
  --request "PUT" \
  --include \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyNDM0NjEsImp0aSI6IjgyIn0.zTMVlrAwRMpaKtXqUu1-foFwqXaWdvYNlU8C05VLCHY" \
  --data '{ "id": 3, "name":"Producto 33", "price":200.33, "image":"Some.33", "description":"Some product description 3" }'
```

4, delete a product by id

```r
curl http://192.168.1.81:80/api/products/3 \
  --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyNDM0NjEsImp0aSI6IjgyIn0.zTMVlrAwRMpaKtXqUu1-foFwqXaWdvYNlU8C05VLCHY" \
  --include
```

5, Add product to homescreen, or delete product from homescreen

- The request to add or remove the product from homescreen it's the same, the backend detect
  automatically if it has to romove or add.

```r
curl http://192.168.1.81:80/api/products/1/onHomeScreen --request "PUT" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE2NjE5ODUsImp0aSI6Ijc5In0.-g8Dk7ynvbBPugoZSRZ1hww4VuxIf5EFxni6Jph1mq8" 

```

#### Product likes

5, add like to product

- The tokes is optional, where if this is emply the anonymous user will be used

```r
curl http://192.168.1.81:80/api/products/4/like \
  --request "POST" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyNDM0NjEsImp0aSI6IjgyIn0.zTMVlrAwRMpaKtXqUu1-foFwqXaWdvYNlU8C05VLCHY"
```

6, remove like to product

- The tokes is optional, where if this is emply the anonymous user will be used

```r
curl http://192.168.1.81:80/api/products/1/like \
  --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyNDM0NjEsImp0aSI6IjgyIn0.zTMVlrAwRMpaKtXqUu1-foFwqXaWdvYNlU8C05VLCHY"
```

#### Product image

7, See image product by Id Image

- Use the browser and this will show the image

```r
http://192.168.1.81/api/products/image/1

```

8, add image to product

```r
curl 'http://192.168.1.81/api/products/2/images' \
  -H 'Content-Type: multipart/form-data; boundary=----WebKitFormBoundaryjNS3Ls6gNAOJjc2g' \
  --data-raw $'------WebKitFormBoundaryjNS3Ls6gNAOJjc2g\r\nContent-Disposition: form-data; name="file"; filename="small_green.png"\r\nContent-Type: image/png\r\n\r\n\r\n------WebKitFormBoundaryjNS3Ls6gNAOJjc2g--\r\n'
```

8, add image to product

```r
curl 'http://192.168.1.81/api/products/image/3' \
  --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTczNDc5MjcsImp0aSI6Ijc5In0.hJDKQw1JTP4XD_vRix4u5hbM89lTbgSpSdP_1FqlhU8"
```

#### Product comments

- The id user is get from the token

1, Add comment to product

```r
curl http://192.168.1.81:80/api/products/1/comment \
  --request "POST" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg5ODYwNTcsImp0aSI6Ijc5In0.QK0ltvGbTp5V3PoZj-gUiFjHGOPiGYu8n1YU83J1fNQ" \
  --data '{ "IdProduct": 1, "content": "First comment", "rating": 3 }'
```

1,1 Add response to a product

- required response_to

```r
curl http://192.168.1.81:80/api/products/13/comment/14 \
  --request "POST" \
  --data '{ "content": "First response to product 13 to comment 14" }' \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQwNjA2NjUsImp0aSI6Ijc5In0.QdLuY-6vWwMYUeeHYkNiFRrHS9RxkDKyiW1i79aorf8"
```

2, Delete comment

```r
curl http://192.168.1.81:80/api/products/1/comment/9 \
  --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg5ODYwNTcsImp0aSI6Ijc5In0.QK0ltvGbTp5V3PoZj-gUiFjHGOPiGYu8n1YU83J1fNQ"
```

3, Get all comments

```r
curl http://192.168.1.81:80/api/products/comments \
  --request "GET" \
  --header "Token:  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQwODUzMjQsImp0aSI6Ijc5In0.zLsCmrWSWYxOUFVY1Z5k4ssY27h17jXNMWl4pltlW2Y"
  
```

### API blogs

1, get all blogs

```r
curl http://192.168.1.81:80/api/blogs --request "GET" | json_pp
```

2, find by id

```r
curl http://192.168.1.81:80/api/blogs/2 --request "GET" | json_pp
```

4, add blog

- id blog will be ignored

```r
curl http://192.168.1.81:80/api/blogs --request "POST" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE4Mjk3MjEsImp0aSI6Ijc5In0.hKTme8DOE65PAoGYKQvb1qLUboNhQaAj4rMkLQPLLyk" \
  --data '{ "title": "Algo 2", "body": "<XX>...</XX><p>...", "abstract": "Some abstract ...", "thumbnail": "data:image/png;base64,xx...==", "author": null, "createdAt": null, "updateAt": null }' \
  | json_pp
```

- Create a blog with onHomeScreen

```r
curl http://192.168.1.81:80/api/blogs --request "POST" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE4Mjk3MjEsImp0aSI6Ijc5In0.hKTme8DOE65PAoGYKQvb1qLUboNhQaAj4rMkLQPLLyk" \
  --data '{ "onHomeScreen": "2022-01-01T01:11:01Z", "title": "Algo 3", "body": "<XX>...</XX><p>...", "abstract": "Some abstract ...", "thumbnail": "data:image/png;base64,xx...==", "author": null, "createdAt": null, "updateAt": null }' \
  | json_pp
```

5, update blog

- The id is required in the json body

```r
curl http://192.168.1.81:80/api/blogs --request "PUT" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE4Mjk3MjEsImp0aSI6Ijc5In0.hKTme8DOE65PAoGYKQvb1qLUboNhQaAj4rMkLQPLLyk" \
  --data '{ "id": 5, "title": "Algo 333", "body": "<XX>333</XX><p>...", "abstract": "Some abstract 333", "thumbnail": "data:image/png;base64,xx...333" }' \
  | json_pp
```

- Update just a field, it's required the ID and the field to update

```r
curl http://192.168.1.81:80/api/blogs --request "PUT" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE4Mjk3MjEsImp0aSI6Ijc5In0.hKTme8DOE65PAoGYKQvb1qLUboNhQaAj4rMkLQPLLyk" \
  --data '{ "id": 5, "onHomeScreen": "2022-01-01T01:13:33Z" }' \
  | json_pp
```

4, delete a blog by id

```r
curl http://192.168.1.81:80/api/blogs/3 \
  --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTExMTY4MTYsImp0aSI6Ijc5In0.7mai9dtJhPEpEMyYBDcEDf_IJ2w0PcPj-JPbEhEPdZs" \
  | json_pp
```

#### Blog Images

Note: The image is addded and updated via base64 in the correspondenting requests for that

3, show image

```r
#Use in the browser
http://192.168.1.81:80/api/blogs/1/image
```

#### Blog Likes

5, add like to blog

- The tokes is optional, where if this is emply the anonymous user will be used

```r
curl http://192.168.1.81:80/api/blogs/2/like \
  --request "POST" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyNDM0NjEsImp0aSI6IjgyIn0.zTMVlrAwRMpaKtXqUu1-foFwqXaWdvYNlU8C05VLCHY"
```

6, remove like to blog

- The tokes is optional, where if this is emply the anonymous user will be used

```r
curl http://192.168.1.81:80/api/blogs/2/like \
  --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDUyNDM0NjEsImp0aSI6IjgyIn0.zTMVlrAwRMpaKtXqUu1-foFwqXaWdvYNlU8C05VLCHY"
```

#### Blog Comments

Only logged user can comment.

1, Add comment

- The id user is get from the token

```r
curl http://192.168.1.81:80/api/blogs/1/comment \
  --request "POST" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQ2NjM0OTgsImp0aSI6Ijc5In0.HP_e5OEEkWB4qJMxi2GUqboyoi912JXS_kdDzvYQwzo" \
  --data '{ "IdBlog": 1, "text": "First comment", "rating": 3 }'
```

2, Delete comment

```r
curl http://192.168.1.81:80/api/blogs/1/comment/9 \
  --request "DELETE" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg5ODYwNTcsImp0aSI6Ijc5In0.QK0ltvGbTp5V3PoZj-gUiFjHGOPiGYu8n1YU83J1fNQ"
```

3, Add response to comment of blog

- required response_to

```r
curl http://192.168.1.81:80/api/blogs/13/comment/14 \
  --request "POST" \
  --data '{ "content": "First response to product 13 to comment 14" }' \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQwNjA2NjUsImp0aSI6Ijc5In0.QdLuY-6vWwMYUeeHYkNiFRrHS9RxkDKyiW1i79aorf8"
```

### API Orders

1, Add Order

```r
curl http://192.168.1.81:80/api/orders \
  --include \
  --request "POST" \
  --data '{ "id_user": 1, "products": [ { "id_product": 1, "amount": 3 }, { "id_product": 91, "amount": 1 } ] }'
```

2, Find by ids

```r
curl http://192.168.1.81:80/api/orders/find \
  --include \
  --request "POST" \
  --data '[1,2,3]'
```

3, Delete by id

```r
curl http://192.168.1.81:80/api/orders/1 --request "DELETE" | json_pp
```

4, Confirm order

```r
curl http://192.168.1.81:80/api/orders/1/confirm \
  --request "PUT" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTUxOTA3MzQsImp0aSI6IjIifQ.tinqy68DlRk2IMpOTwdsgUVOos6-FFL3gQYX3oKg2AE"
  | json_pp
```

5, Find All

- admin user required

```r
curl http://192.168.1.81:80/api/orders \
  --request "GET" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTUzNTkyNzksImp0aSI6Ijc5In0.LjaV5IU3ZcVETZkRssuEpz3gL-FQTUE4GqboJkVvBEM" \
  | json_pp
```

6, Confirm order

- admin user required

```r
curl http://192.168.1.81:80/api/orders/1/accept \
  --request "PUT" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTUxOTA3MzQsImp0aSI6IjIifQ.tinqy68DlRk2IMpOTwdsgUVOos6-FFL3gQYX3oKg2AE" \
  | json_pp
```

6, Get orders of user logged

- token required

```r
curl http://192.168.1.81:80/api/orders/findByUserLogged \
  --request "GET" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTYwMTAyMzQsImp0aSI6IjIifQ.wk5HVH6fW4H5Q6xqWz32ACDSq3KGVBKGntYWzRiUSVc" \
  | json_pp
```

7, get info payment

```r
curl http://192.168.1.81:80/api/orders/paymentInfo \
  --request "GET" \
  --header "Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjUwMDQ1OTgsImp0aSI6Ijc5In0.V9xqyLx1tuDkmmVsNuRMj8FfwcVcTmJnKGX9PePMe_4" \
  | json_pp
```

## SSR (Server side render)

### Products

```r
#Use in the browser
http://192.168.1.81:80/products
```

```r
#Use in the browser
#El id del producto debe estar al final y le debe de prececeder un guion medio (-)
http://192.168.1.81/products/nombre-producto-11
```

### Blogs

```r
#Use in the browser
http://192.168.1.81:80/blogs
```

```r
#Use in the browser
#The Blog's ID must be at the end of the url, and must be precedided by a middle dash (-)
http://192.168.1.81/blogs/el-titulo-del-blog-1
```
