# PCS-API

### API error code

Here are the different error codes that the API can return

| Code erreur | Description                                                                                                                                                    |
|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 1           | The password must be greater than 8 and less than 128 characters long, at least have an uppercase letter, a lowercase letter, a number and a special character |
| 2           | Invalid email                                                                                                                                                  |
| 3           | Bad typeUser                                                                                                                                                   |
| 4           | Missing content userDTO                                                                                                                                        |
| 5           | Email already exists                                                                                                                                           |
| 6           | Phone already exists                                                                                                                                           |
| 7           | Wrong login or password                                                                                                                                        |
| 8           | invalid token                                                                                                                                                  |
| 9           | Unauthorized chat access                                                                                                                                       |
| 10          | userID uuids in chatDTO are invalid                                                                                                                            |
| 11          | Invalid chat creation                                                                                                                                          |
| 12          | Invalid message in chatDTO                                                                                                                                     |
| 13          | Invalid message creation                                                                                                                                       |
| 14          | User is not the right type, invalid action                                                                                                                     |
| 15          | PropertyDTO invalid attribute                                                                                                                                  |
| 16          | Invalid number of files                                                                                                                                        |
| 17          | Property deletion not allowed                                                                                                                                  |
| 18          | Updated an unauthorized profile because the user is not admin or is not the profile owner                                                                      |
| 19          | Invalid service type                                                                                                                                           |
| 20          | No authorization to manipulate the service                                                                                                                     |
| 21          | Invalid property id                                                                                                                                            |
| 22          | The reservation date is invalid                                                                                                                                |
| 23          | Bill creation is invalid                                                                                                                                       |
| 24          | The user is not a traveler                                                                                                                                     |
| 25          | Invalid service                                                                                                                                                |
| 26          | Property payment creation error                                                                                                                                |
| 27          | Service payment creation error                                                                                                                                 |
| 28          | Wrong session stripe creation parameter                                                                                                                        |
| 29          | Invalid id                                                                                                                                                     |
| 30          | Reservation not found                                                                                                                                          |
| 31          | The durations are not the same                                                                                                                                 |


### Application error code

Here are the various error codes that may occur in the event of an error in the program

| Code erreur | Description                                  |
|-------------|----------------------------------------------|
| 1           | Error opening config.env file                |
| 2           | Error when trying to connect to the database |
| 3           | Error convert str to int (env)               |
| 5           | invalid token key                            |