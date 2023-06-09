# Golang JWT MySQL

## API Reference

#### Login

```http
  POST /login
```

| Parameter  | Type     | Description   |
| :--------- | :------- | :------------ |
| `username` | `string` | **Required**. |
| `password` | `string` | **Required**. |

username `user` and `admin` password `secret`

#### Register

```http
  POST /register
```

| Parameter  | Type     | Description   |
| :--------- | :------- | :------------ |
| `username` | `string` | **Required**. |
| `password` | `string` | **Required**. |

#### Logout

```http
  GET /logout
```

#### Get All Produk

```http
  GET /produk
```

#### Add Produk

```http
  POST /produk
```

| Parameter     | Type     | Description   |
| :------------ | :------- | :------------ |
| `nama_produk` | `string` | **Required**. |
| `stok`        | `string` | **Required**. |

#### Edit Produk

```http
  PUT /produk
```

| Parameter     | Type     | Description   |
| :------------ | :------- | :------------ |
| `id`          | `string` | **Required**. |
| `nama_produk` | `string` | **Optional**. |
| `stok`        | `string` | **Optional**. |

#### Delete Produk

```http
  DELETE /produk
```

| Parameter | Type     | Description   |
| :-------- | :------- | :------------ |
| `id`      | `string` | **Required**. |

#### Get Cart

```http
  GET /cart
```

#### Add to Cart

```http
  POST /cart
```

| Parameter   | Type     | Description   |
| :---------- | :------- | :------------ |
| `id_produk` | `string` | **Required**. |
| `jml`       | `string` | **Required**. |

#### Edit Cart

```http
  PUT /cart
```

| Parameter   | Type     | Description   |
| :---------- | :------- | :------------ |
| `id`        | `string` | **Required**. |
| `id_produk` | `string` | **Required**. |
| `jml`       | `string` | **Required**. |

#### Checkout

```http
  POST /checkout
```
