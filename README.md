# Golang Crud App Demo
The following app is a golang demo of a simple CRUD app.
The solution can be run locally or as a dockerized solution.

# Requirements
The following tools will be needed
* docker
* docker-compose
* go (version 1.18.1 was development env)

Manual installation of the following packages
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/gorilla/mux
go get github.com/spf13/viper
```

# Building the Image
You can build the image using the following command from the root of the project
`docker build -t golang_crud -f Dockerfile . --no-cache`

To run the docker-compose, make sure that you have first the app image as described above. Next you can run the docker-compose as follows:
`docker-compose -f docker-compose.yml up`

# Running locally
If you want to run locally, you will need to edit the config json and change the following paramenters:
* host=localhost
* port=54322
This is assuming you run the database in the same way as configured in docker-compose.

# Why the wait script
Even if the `depends_on` parameter is specified in docker-compose,the service might be responding while the database is not open. This causes the app to fail. The wait script allows us to check if the database is open without crashing the app.

# Running through docker-compose
If you start the application through docker-compose, you will see some output as below.
Notice that the wait script will report if the db is up.

```
app_1          | --------------------------------------------------------
app_1          |  docker-compose-wait 2.7.2
app_1          | ---------------------------
app_1          | Starting with configuration:
app_1          |  - Hosts to be waiting for: [postgressdb:5432]
app_1          |  - Timeout before failure: 300 seconds 
app_1          |  - TCP connection timeout before retry: 30 seconds 
app_1          |  - Sleeping time before checking for hosts availability: 0 seconds
app_1          |  - Sleeping time once all hosts are available: 0 seconds
app_1          |  - Sleeping time between retries: 30 seconds
app_1          | --------------------------------------------------------
app_1          | Checking availability of postgressdb:5432
app_1          | Host postgressdb:5432 is now available!
app_1          | --------------------------------------------------------
app_1          | docker-compose-wait - Everything's fine, the application can now start!
app_1          | --------------------------------------------------------
app_1          | 2022/08/31 11:27:02 Loading Server Configurations...
app_1          | 2022/08/31 11:27:02 Connecting to database %s host=postgressdb user=postgres dbname=postgres port=5432
app_1          | 2022/08/31 11:27:02 Connected to Database...
pg_golang      | 
pg_golang      | PostgreSQL Database directory appears to contain a database; Skipping initialization
pg_golang      | 
pg_golang      | 2022-08-31 11:27:02.540 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
pg_golang      | 2022-08-31 11:27:02.540 UTC [1] LOG:  listening on IPv6 address "::", port 5432
pg_golang      | 2022-08-31 11:27:02.547 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
pg_golang      | 2022-08-31 11:27:02.886 UTC [1] LOG:  database system is ready to accept connections
pg_golang      | 2022-08-31 11:27:02.894 UTC [33] LOG:  incomplete startup packet
app_1          | 2022/08/31 11:27:02 Done auto migration
app_1          | 2022/08/31 11:27:02 Starting Server on port 8080
app_1          | 2022/08/31 11:28:00 Creating new database entry
app_1          | 2022/08/31 11:28:00 Inserting db record
```

Once the app is up you can use postman for example to create a record in the db.
You can do a post to the following URL `http://localhost:8080/api/products` using the body payload

```
{
"name": "test-product", "description": "random-description", "price": 100.00
}
```

This will result to a record being inserted as shown from the last 2 lines of logging in docker-compose