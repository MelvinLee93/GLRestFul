# GoLang RESTFul API Prototype
A prototype of RestAPI with GoLang with basic CRUD functions using Gorm and PostGreSQL for database functions.
More features may be added in the future.

## Setting up and Running the Project
1. Download the API and place it on a folder.
2. Install Go from the provided link [here](https://go.dev/dl/).
3. Open CMD, change the directory to the file where you have placed the RESTFul API prototype (for e.g. Cd C:/users/usename/RestfulAPI).
4. Once you are in the directory, type `Go get` to download all Golang plugins and modules needed for the project.
5. To run the API project, type `Go Run .` or `Go Run main.go` to start the API.

## Setting up the Database
1. Download and install PostGreSQL from the provided link [here](https://www.postgresql.org/download/).
> [!NOTE]
> During the installation, when asked for the username and password, keep the username as "postgres" while keep the password as "S8rxc34Vi". These variables can be changed on the config.go file when needed.
3. Download PGAdmin4 from the provided link [here](https://www.pgadmin.org/download/pgadmin-4-windows/).
4. Once installed, open PGAdmin4 and login to the database with the password you have created for the server during the PostGreSQL installation.
5. Assuming you have not created the database, right click on the "Database" category just under "PostgreSQL 17" and name it "db_employee".
6. Once the database is created, right click on db_employee, and click "Query Tool".
7. In the Query tab, copy the query below, paste it on the tab, and then click the play button or press F5 to execute the script:
   ```
   create table users ( 
     primary key(id),
	 id serial not null, 
     username VARCHAR(100) UNIQUE not null, 
     password VARCHAR(100) UNIQUE not null, 
     email VARCHAR(100) UNIQUE,
	 created_at TIMESTAMP,
	 updated_at TIMESTAMP
	 );
   ```

## Testing Via Postman
Firstly, download and install Postman from the provided link [here](https://www.postman.com/downloads/).
### Testing the JWT Authentication
#### Login Test
1. On the address bar provided in Postman, to the left of it, switch the setting from **GET** to **POST**.
2. On the "Headers" tab, add in the key as "Authorization" with the value "Bearer penacony".
3. On the "Authorization" tab, change the Auth Type into "JWT Bearer".
4. Ensure the Algorithm is HS256 and the Secret set to {{Bearer}}.
5. On the "Body" tab, choose "x-www-form-urlencoded" and type in these values:
   ![image](https://github.com/user-attachments/assets/694090e9-1134-4181-9488-85cd00d464f2)
6. Once done, click the "send" button and assuming it is successful, it will show this on the console at the bottom:
   ![image](https://github.com/user-attachments/assets/9aa474d2-8a16-4dea-aa24-79e5a78d254e)
### Testing Access To Restricted Route
1. On the address bar provided, type ``http://localhost:8080/restricted``.
2. On the "Authorization" tab, switch the Auth Type to "Bearer Token". Leave the Token to have {{Bearer}} value.
3. On the headers, remove "penacony" and add it with the token you have attained from the previous step, so it will look like this:
   ![image](https://github.com/user-attachments/assets/942cab09-d7c8-4ba5-be8a-a20476b4b2f4)
4. Click the "send" button and upon success, it should show the message "Welcome Galactic Baseballer!" on the console below.

## Testing CRUD function endpoints

> [!NOTE]
> Guest CRUD function is also allowed, but only the **GET** function is available. Putting ``http://localhost:8080/`` will yield the message "Welcome, Guest!" on the console below. Getting the data will be similar as the **GET** on the restricted route, except it will be ``http://localhost:8080/:id`` with ``:id`` denoting the entry within the database (e.g. ``http://localhost:8080/1``).

Most of the CRUD operations are within the restricted route:
1. **GET**: ensure the setting is changed to **GET** and then type ``http://localhost:8080/restricted/:id`` with ``:id`` denoting the entry within the database (e.g. ``http://localhost:8080/restricted/1``). If the result shows a JSON data other than error messages, then the GET function is successful.
2. **POST**: ensure the setting is changed to **POST** and then type ``http://localhost:8080/restricted/`` before going to "Body" tab underneath the address bar. Select **Raw** and type a query sample like the format below:
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
3. **PATCH**: ensure the setting is changed to **PATCH** and then type ``http://localhost:8080/users/:id`` with ``:id`` denoting which entry do you wish to update within the database (e.g. ``http://localhost:8080/restricted/1``). To update the database, use the code format similar to the POST method, with different input data as shown below:
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
  4. **DELETE**: ensure the setting is changed to **DELETE** and then type ``http://localhost:8080/restricted/:id`` with ``:id`` denoting which entry do you wish to update within the database (e.g. ``http://localhost:8080/restricted/:6``). If the execution is successful, you will get this message:
    ```
    {
        "message": "An employee has been deleted"
    }
    ```
