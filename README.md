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