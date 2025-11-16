# Routes

This file describe the routes of the app

- `/home`
    - [ ] GET
- `/authentification`
    - [ ] GET
    - `/register`
        - [ ] POST
    - `/login`
        - [ ] POST
    - `/logout`
        - [ ] POST
- `/question`
    - [ ] GET
    - [ ] POST
- `/profile`
    - [ ] GET
    - [ ] PUT
    - [ ] DELETE
- `/image/:path_of_image`
    - [ ] GET
- `/score`
    - [ ] GET

## Details

### `/home`

- #### GET

Golang Structure

```go
type Home struct {
    HasError    bool
    ErrorDetail struct {
        ErrorType string
        ErrorText string
    }
    IsLogin  bool
    UserInfo struct {
        Name  string
        Score int
    }
}
```

### `authentification`

- #### GET

Golang Structure

```go
type Authentification struct {
    HasError       bool
    ErrorDetail struct {
        ErrorType string
        ErrorText string
    }
}
```

### `/authentification/register`

- #### POST

expected JSON from the front

```json
{
    email: string,
    pseudo: string,
    password: string,
    password_confirmation: string
}
```

- data **valid**: rediret to `/home`
- data **invalid** or error in the server: rediret to `/auth/register` with error in the structure

### `authentification/login`
entification
- #### POST

expected JSON from the front

```json
{
    email_or_pseudo: string,
    password: string
}
```

- data **valid**: rediret to `/home`
- data **invalid** or error in the server: rediret to `/auth/register` with error in the structure

### `authentification/logout`

- #### POST

just a request **POST** for logout the connected user

### `/question`

- #### GET

Golang Structure

```go
type Question struct {
    HasError       bool
    ErrorDetail    struct {
        ErrorType string
        ErrorText string
    }
    State          bool
    ID             int
    Content        string
    HasImage       bool
    ImagePath      string
    Option []struct {
        ID   int
        Text string
    }
    UserHasCorrect bool
    OldUserScore   int
    NewUserScore  int
}
```

If `State` is set at **true**, the user must answer **else** the user has already answer to the question and discover the truth. After 30sec, the user is redirect to a new question.

- #### POST

expected JSON from the front

```json
{
    question_id: int,
    user_answer_id: int
}
```

### `/profile`

- #### GET

Golang Structure

```go
type Profile struct {
    HasError    bool
    ErrorDetail struct {
        ErrorType string
        ErrorText string
    }
    Pseudo string
    Email  string
    Score  string
}
```

- #### PUT

expected JSON from the front

```json
{
    old_pseudo: string
    new_pseudo: string
}
```

- ### DELETE

just a **DELETE** request for delete the user who are connected.

### `/image/:path_of_image`

- ### GET

Send the image of the path.

### `/score`

- ### GET

Golang Structure

```go
type Profile struct {
    HasError    bool
    ErrorDetail struct {
        ErrorType string
        ErrorText string
    }
    List struct {
        Pseudo string
	    Score int
    }
}
```