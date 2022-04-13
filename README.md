
# Inprinte Backoffice-backend 🧙‍♂️

This is the backend part of the backoffice part



## Installation

To run the project, you need golang 1.17, so you can install it using 
[this link](https://khongwooilee.medium.com/how-to-update-the-go-version-6065f5c8c3ec)
    
## Run Locally

Clone the project

```bash
https://github.com/inprinte/backoffice-backend.git
```

Go to the project directory

```bash
cd backoffice-backend
```

Start the server

```bash
go run main.go 
```


## Deployment

To deploy this project

Clone the project 

```bash
https://github.com/inprinte/backoffice-backend.git
```

Go to the project directory

```bash
cd backoffice-backend
```

Build the docker image

```bash
docker build --tag inprinte-backoffice-backend .
```

Run the docker image specifying the port 

```bash
sudo docker run -d -p 8081:8080 inprinte-backoffice-backend
```


## Authors

- [@romainsilvy](https://github.com/romainsilvy)
- [@Ayatooo](https://github.com/Ayatooo)
- [@matteodinville](https://github.com/matteodinville)



