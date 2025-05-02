# GoLang RESTFul API Prototype
A prototype of RestAPI with GoLang with basic CRUD functions using Gorm and PostGreSQL for database functions.
More features may be added in the future.

# Running the API Locally
1. Open CMD, change the directory to the file of the RESTFul API (Cd C:/users/usename/RestfulAPI, for e.g.).
2. Assuming you have Go already install on your computer, type `Go Run .` or `Go Run main.go` to start the API.

# Testing CRUD function endpoint via Postman
With the Postman program downloaded and opened, and with the API connected to the proper database, to test the CRUD operations:
1. **GET**: ensure the setting is changed to **GET** and then type ``http://localhost:8080/users/:id`` with ``:id`` denoting the entry within the database (e.g. ``http://localhost:8080/users/:1``). If the result shows a JSON data other than error messages, then the GET function is successful.
2. **POST**: ensure the setting is changed to **POST** and then type ``http://localhost:8080/users/`` before going to "Body" tab underneath the address bar. Select **Raw** and type a query sample like the format below:
   ```
    {
    "username":"testdata",
    "password":"testpassword",
    "email":"test@testmail.com"
    }
   ```
   Once the data is entered and sent, successful entry will show the complete data structure of the data you have sent to the database, which will look like:

   ```
    {
      "data": {
        "id": 0,
        "username": "testdata",
        "password": "testpassword",
        "email": "test@testmail.com",
        "created_at": "2025-05-01T17:31:18.3458927+07:00",
        "updated_at": "2025-05-01T17:31:18.3458927+07:00"
      }
    }
    ```
3. **PATCH**: ensure the setting is changed to **PATCH** and then type ``http://localhost:8080/users/:id`` with ``:id`` denoting which entry do you wish to update within the database (e.g. ``http://localhost:8080/users/:6``). To update the database, use the code format similar to the POST method, with different input data as shown below:
    ```
    {
        "username":"newdata1",
        "password":"newpassword",
        "email":"test@testmail.com"
    }
    ```
     Once the data is entered and sent, successful entry will show the complete data structure of the data you wished to update, which will look like:
    ```
      {
        "data": {
            "id": 6,
            "username": "newdata1",
            "password": "newpassword",
            "email": "test@testmail.com",
            "created_at": "2025-05-01T17:31:18.345892Z",
            "updated_at": "2025-05-01T17:48:28.0824109+07:00"
        }
      }
    ```
  4. **DELETE**: ensure the setting is changed to **DELETE** and then type ``http://localhost:8080/users/:id`` with ``:id`` denoting which entry do you wish to update within the database (e.g. ``http://localhost:8080/users/:6``). If the execution is successful, you will get this message:
    ```
    {
        "message": "An employee has been deleted"
    }
    ```
