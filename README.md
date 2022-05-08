<h1>GoMasters REST API</h1>

This REST API was created as a study project.</br>

Used libs:
* router: go-chi/chi
* validation: go-playground/validator;
* uuid: google/uuid;
* postgres driver: jackc/pgx;
* read envs: kelseyhightower/envconfig;
* logger: go.uber.org/zap;
* read yaml: gopkg.in/yaml.

DB: PostgreSQL.</br>

Requests:
<pre>
GET / - get index
GET /users - get all users
POST /users - create user
GET /users/{id} - get user
PUT /users/{id} - edit user
DELETE /users/{id} - delete user
</pre>

Main entity:
<pre>
type User struct {
    ID uuid
    Firstname string
    Lastname string
    Email string
    Age int
    Created time.Time
}
</pre>

INDEX</br>
![Index](https://user-images.githubusercontent.com/21006294/167303132-684c359b-3021-4c88-bb18-9ad9540f54e5.png)

GET ALL RECORDS</br>
![DB ALL](https://user-images.githubusercontent.com/21006294/167303126-14bbfaca-4cdd-4095-8057-663a7caa01ae.png)
![Postman ALL](https://user-images.githubusercontent.com/21006294/167303133-a0bbdd95-644b-4b5b-94d9-e1184b532ec5.png)

POST RECORD</br>
![Postman POST](https://user-images.githubusercontent.com/21006294/167303135-526bf30c-b2c6-4656-add5-c1d081b9726b.png)
![DB POST](https://user-images.githubusercontent.com/21006294/167303138-905efca6-91f4-4764-99f9-a68c2a124a48.png)

GET RECORD BY ID</br>
![Postman GET BY ID](https://user-images.githubusercontent.com/21006294/167303134-3b1e8879-0f56-4d7e-b3b7-acdb3e8f57ab.png)

PUT (UPDATE) RECORD BY ID</br>
![Postman PUT](https://user-images.githubusercontent.com/21006294/167303141-021b8629-90a3-4ef1-8fa2-18291d79668e.png)
![DB PUT](https://user-images.githubusercontent.com/21006294/167303142-ceb85462-81fe-46ff-9812-effafdd0933a.png)

DELETE RECORD BY ID</br>
![Postman DELETE](https://user-images.githubusercontent.com/21006294/167303130-933deb23-ff10-48d9-926c-77144b509f06.png)
![DB DELETE](https://user-images.githubusercontent.com/21006294/167303131-e0afc1b6-7eaa-425d-bfa3-1de2143b6327.png)