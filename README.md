# Inprinte backoffice frontend

This is the frontend part of the backoffice of the project Inprinte



## Installation

To run the project, you need node js 17, you can install this using 
[this link](https://phoenixnap.com/kb/update-node-js-version)
    
## Run Locally

Clone the project

```bash
https://github.com/inprinte/backoffice-frontend
```

Go to the project directory

```bash
cd backoffice-frontend/inprinte-app
```

Install dependencies

```bash
npm install
```

Start the server

```bash
npm start
```


## Deployment

To deploy this project

Clone the project 

```bash
https://github.com/inprinte/backoffice-frontend
```

Go to the project directory

```bash
cd backoffice-frontend/inprinte-app
```

Build the docker image

```bash
docker build --tag inprinte-backoffice-frontend .
```

Run the docker image specifying the port 

```bash
sudo docker run -d -p 3001:3000 inprinte-backoffice-frontend
```


## Authors

- [@romainsilvy](https://github.com/romainsilvy)
- [@Ayatooo](https://github.com/Ayatooo)
- [@matteodinville](https://github.com/matteodinville)
