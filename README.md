# admin.claps.dev

## Project setup
```
yarn install
```

### Compiles and hot-reloads for development
```
yarn serve
```

### Compiles and minifies for production
```
yarn build
```

##  Deploy to the server side

### Execute the command to generate the dist package
```
yarn build
```
the dist package is ready to be deployed to the server

## Port and other configuration settings

### For development
Edit the .env.development file
```
# port for development
VUE_APP_PORT=8001

# base api url for development
VUE_APP_API_URL='http://localhost:7003/'
```

### For server
Edit the .env file
```
# port
VUE_APP_PORT=8001

# base api url
VUE_APP_API_URL='https://admin-api.claps.dev/'
```
