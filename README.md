# CRUD-API-USING-GO

- Implementation of Basic Crud API using Clean Architecture in Golang
- Here, we can create a user where user does CRUD operation on a book<br><br>

- How?
    - After cloning the repo you need to follow the few steps:
        - Run `make setup` command in the root directory using terminal
        - The above command will create and run a docker container that is built with `mysql` image
        - After that, `make run` command, this will do the DB schema migration and the server starts listening at `8080` port.
        - Postman Collection for the endpoints have been added as well.<br><br><br>

- To check the Database:
    - Run `docker exec -it mysql-crud-api mysql -p`
    - Enter the password `password`
    - Run `use crud_api_db;`
    
- Future Enhancements:
    - Make email_id as unique field and apply race condition.
    - Add Authentication using middlewares.
    - Try to implement more complex business logics for this application.
    - Containerise the whole app.
    - Add Separate commands for migration up and down using cobra cli and viper.
