# Inprinte backend

This is the backend part of the project Inprinte



## Installation

To run the project, you need golang 1.17, so you can install it using 
[this link](https://khongwooilee.medium.com/how-to-update-the-go-version-6065f5c8c3ec)
    
## Run Locally

Clone the project

```bash
https://github.com/inprinte/backend
```

Go to the project directory

```bash
cd backend
```

Start the server

```bash
go run main.go 
```


## Deployment

To deploy this project

Clone the project 

```bash
https://github.com/inprinte/backend
```

Go to the project directory

```bash
cd backend
```

Build the docker image

```bash
docker build --tag inprinte-backend .
```

Run the docker image specifying the port 

```bash
sudo docker run -d -p 8080:8080 inprinte-backend
```


## API Reference

#### Get home informations

```http
GET /
```

#### Get shop informations

```http
GET /boutique
```

#### Get product informations

```http
GET /produit/{id_product}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id_product` | `string` | **Required**. Id of product to fetch |

#### Get user informations

```http
GET /user/{id_user}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id_user` | `string` | **Required**. Id of user to fetch |


#### Login

```http
POST /connexion
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email` | `json` | **Required**. email of the user |
| `password` | `json` | **Required**. password of the user |


#### Sign up 

```http
POST /inscription
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `firstname` | `json` | **Required**. firstname of the user |
| `lastname` | `json` | **Required**. lastname of the user |
| `email` | `json` | **Required**. email of the user |
| `password` | `json` | **Required**. password of the user |
| `phone` | `json` | **Required**. phone of the user |


## Authors

- [@romainsilvy](https://github.com/romainsilvy)
- [@Ayatooo](https://github.com/Ayatooo)
- [@matteodinville](https://github.com/matteodinville)

