# USER REST API

REST API service responsible for managing user information across the entire application or ecosystem.

## Structure:
4 Domain layers:

- Models layer
- Repository layer
- UseCase layer
- Delivery layer

## API:

### POST /api/users

Creates new user

##### Example Input: 
```
{
	"first_name": "Jason",
	"last_name": "Ivanov",
	"email": "j.ivanov@mail.ru",
	"phone": "+334239999999",
	"login": "jason_ivanov_45",
	"password": "qwerty12345",
	"repeat_password": "qwerty12345"
} 
```

### GET /api/users

Returns all users

##### Example Response: 
```
{
	[
	  {
	    "uuid": "8a3a174e-5620-4374-a4ef-d56478249b15",
	    "first_name": "Jason",
	    "last_name": "Ivanov",
	    "email": "j.ivanov@mail.ru",
	    "phone": "+3752989824238",
	    "password": "qwerty12345",
	    "login": "jason_ivanov_45"
	  },
	  {
	    "uuid": "01be8933-da7a-41fa-9482-80529096fbd2",
	    "first_name": "Johny",
	    "last_name": "Petrov",
	    "email": "petrov.jo@mail.ru",
	    "phone": "+334239999988",
	    "password": "qwerty12345",
	    "login": "johny_petrov_99"
	  }
	]
} 
```

### GET /api/users/:uuid

Returns user with uuid

##### Example Response: 
```
{
	"uuid": "8a3a174e-5620-4374-a4ef-d56478249b15",
	"first_name": "Jason",
	"last_name": "Ivanov",
	"email": "j.ivanov@mail.ru",
	"phone": "+3752989824238",
	"password": "qwerty12345",
	"login": "jason_ivanov_45"
} 
```

### DELETE /api/users/:uuid

Deletes user with uuid

### PUT /api/users/:uuid

Updates user with uuid

##### Example Input: 
```
{
	"first_name": "Johny",
	"last_name": "Petrov",
	"email": "petrov.jo@mail.ru",
	"phone": "+334239999988",
	"login": "johny_petrov_99",
	"old_password": "qwerty12345",
	"new_password": "qwerty56789"
} 
```

## Requirements
- go 1.22.5
- docker & docker-compose
- julienschmidt/httprouter
- sirupsen/logrus
- ilyakaznacheev/cleanenv
- jackc/pgx

## Run Project

Create ```.env``` file in root directory and add following values:
```dotenv
DB_PASSWORD=your_database_password
```
Create ```logs/app.log``` file in root directory

Use ```make linux_run``` to build and run application for linux-based system

```or```

Use ```make win_run``` to build and run application for windows
