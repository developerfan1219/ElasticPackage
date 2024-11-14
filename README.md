# ElasticPackage

Golang + Echo + Vue = Simple Order Dashboard.


## Requirements

- PostgreSQL 17
- Golang 1.23.0 windows/amd64
- Echo v4
- Vue 3 | Vuetify 3
- Node 22.8.0
- NPM 10.9.0


After completing the installation, your environment is ready for Installation.

## Install

### Initial DB Setup
- Edit db/config.json file
- In db directory, run ``` go mod tidy```
- copy & paste your db csv files to db/csv folder
- the file name has to be same with the table name. for example, ``` orders.csv order_items.csv ```
- migrate the dump. run the command in ``` db ``` folder ``` go run migrate.go ```

### Server
- move to ```server``` folder
- install modules ```go mod tidy```

### Vue
- move to ```vue-client``` folder
- install packages ``` npm install ```
- create ```env``` folder in the ```vue-client``` folder
- make ```.env``` file
- add the env variables
```
# .env
VITE_APP_API_BASE_URL=http://localhost:8080
VITE_APP_MAPBOX_ACCESS_TOKEN=
```
### Run 
- move to ```server``` folder and run ```go run server.go```
- move to ```vue-client``` folder and run ```npm run dev```
- visit ```http://localhost:3000```

### Pages
- Home Page - ```Part 2```
- Orders Page - ```Part 1```
