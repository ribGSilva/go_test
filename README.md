This project is a study of how to create the basics of a microservice using gin with go

Uses mongo at default configs for testing

---------------------------------------
Post example:
http://localhost:8080/user
{
    "name": "josevaldo",
    "age": 54
}
Response:
{
    "id": "609e905afaf66e3399d37ecf"
}


---------------------------------------
Get example:
http://localhost:8080/user/{id}
Response:
{
    "id": "609e82fca83c2e9458773754",
    "name": "josevaldo",
    "age": 54
}
