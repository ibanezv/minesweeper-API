# API-Minesweeper Endpoints

## Endpoints Games
- Game Creation
    - Method: POST
    - URI:'/games'
    - Body Request:
        ```
            {
                "id":0,
                "user_id":12,   
                "count_rows":10,
                "count_cols":10,
                "count_mines":10,
                "state":""
            }
        ```
    - Description: This endpoint generate a new game and distributions grid to.

- Game Get
    - Method: GET
    - URI:'/games/{ID}'
    - Body Response:
        ```
            {
                "id":5,
                "user_id":12,   
                "count_rows":10,
                "count_cols":10,
                "count_mines":10,
                "state":""
            }
        ```
    - Description: Get a game

- Game Distribution Get
    - Method: GET
    - URI:'/games/{ID}/distributions'
    - Body Response:
        ```
            [
                {
                    "game_id":5,
                    "row_number":0,   
                    "col_number":0,
                    "value":"",
                    "state":"hidden",
                },
                {
                    "game_id":5,
                    "row_number":0,   
                    "col_number":1,
                    "value":"selected",
                    "state":"showed",
                }
            ]
        ```
    - Description: Get a specific distribution.
- Set a cell as selected
    - Method: PATCH
    - URI:'/games/{id}/distributions'
    - Body Response:
        ```
            {
                "game_id":5,
                "row_number":0,   
                "col_number":1,
                "value":"selected",
                "state":"",
            }
        ```
    - Description: Set a new value to a cell in distribution grid, value could be 'selected'|'flag'|'question'

## Endpoints Users

- User Creation
    - Method: POST
    - URI:'/users'
    - Body Request:
        ```
            {
                "id":0,
                "nick_name":"test",
                "account_id":5,
            }
       ```
    - Description: This endpoint generate a new user.

- User Get
    - Method: GET
    - URI:'/users/{ID}'
    - Body Response:
        ```
            {
                "id":1,
                "nick_name":"test",
                "account_id":5,
            }
        ```
    - Description: Get a user.

- Get user games
    - Method: GET
    - URI:'/users/{ID}/games'
    - Body Response:
        ```
        [
            {
                "id":5,
                "user_id":12,   
                "count_rows":10,
                "count_cols":10,
                "count_mines":10,
                "state":"in_process"
            },
            {
                "id":6,
                "user_id":12,   
                "count_rows":10,
                "count_cols":10,
                "count_mines":10,
                "state":"finished"
            }
        ]
        ```
    - Description: Get games of user.

## Endpoints Accounts
- Account Creation
    - Method: POST
    - URI:'/acounts'
    - Body Request:
        ```
            {
                "id":0,
                "email":"test@test.com",
            }
       ```
    - Description: This endpoint generate a new account.

- Account Get
    - Method: GET
    - URI:'/accounts/{ID}'
    - Body Response:
        ```
            {
                "id":1,
                "email":"test@test.com",
            }
        ```
    - Description: Get a account.

- Get acount users
    - Method: GET
    - URI:'/accounts/{ID}/users'
    - Body Response:
        ```
        [
            {
                "id":1,
                "nick_name":"test",
                "account_id":5,
            },
            {
                "id":2,
                "nick_name":"test2",
                "account_id":5,
            }
        ]
        ```
    - Description: Get user of account.

<br>

[Back home](/README.md)
