# POE-debugServer
It's a data Server for Developing a data stats server which service for Path of exile.

It's made for simulating the way how Path of exile offical server gives data.
For debugging our applications or websites, either trade tool.
Only set with internally default character data or trade data, without update data from Offical server.

(Which mean it's a localhost server used for confirming your data struct is right or not, also helping you to debug your application to make sure your handle is right if the server gave the unexpect struct or shut down for maintenance.
Also have a board to contral the debug server to different status and ensure your application errors handle is fine.)



**Expect feature**:
- maintenance control                            (to test when the server shut down accidentally)
- maintenance control with timer
- API simulate with contralable results          (including trade and character data.)
- Gives illegal data struct randomly.            (including gives depercated data, maybe?)
- **.... etc**

___
# Install
git clone --depth 10 https://github.com/kwangsing3/POE-debugServer.git
# Requirement
go.mod
```
module poe-debugserver

go 1.13

require (
	github.com/codegangsta/negroni v1.0.0
	github.com/gorilla/context v1.1.1
	github.com/gorilla/mux v1.8.0
	github.com/phyber/negroni-gzip v1.0.0
	github.com/rs/cors v1.7.0
	github.com/unrolled/render v1.1.1
	gopkg.in/tylerb/graceful.v1 v1.2.15
)

```
___

## Default domin:
http://localhost:9999
## API path:
API path follow as the offcial server, so it could very easily switch between production and develop while developing.


### API document.
```
/
```
 (unfinished)
### Setting Server status by query.
```
/status
```

| KEY        | VALUE  |
| ------------- | -----:|
| Maintenance         | boolen |

### Get account name by character name 
```
/character-window/get-account-name-by-character
```
( default result: "defaultAccount")

| KEY        | VALUE  |
| ------------- | -----:|
| character         | string |

### Setting Server status by query.
```
/character-window/get-characters
```

| KEY        | VALUE  |
| ------------- | -----:|
| accountName         | string |

### Get Current leagues.
```
/api/leagues
```

### Setting Server status by query.
```
/character-window/get-characters
```

| KEY        | VALUE  |
| ------------- | -----:|
| Maintenance         | boolen |

### Get Ladders
```
/api/ladders
```