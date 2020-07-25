# Hermes

[![Build Status](https://travis-ci.com/paulmj7/hermes.svg?branch=master)](https://travis-ci.com/paulmj7/hermes)

Hermes is an filesystem management api to relay the contents of a volume in real time and provide system controls over HTTP. When paired with [Homebase](https://github.com/paulmj7/homebase), they provide a distributed, multivolume storage server.

From [Wikipedia](https://en.wikipedia.org/wiki/Hermes), "Hermes is considered the herald of the gods ... Hermes functioned as the emissary and messenger."

## Installation

First, customize the config.json file in the src folder to serve your root volumes, as well as change the port if necessary.

To run Hermes, either build from the Docker file or =>
```bash
cd src
go run .
```

## Usage

- POST /api
  - 200 sends the root volumes being served as JSON
- POST /api/change_dir
  - 200 sends the contents of the current directory as JSON
- POST /api/retrieve
  - 200 sends the url route to download the specified file
  - ex. returns x91js -> localhost:5000/api/send?key=x91js
- GET /api/send?key=
  - 200 streams the file for download
- POST /api/upload
  - 201 sends file through multipart form to the server at the specified path
- POST /api/create
  - 201 creates a named folder at the specified path
- PUT /api/move
  - 204 moves folder or file to the specified path
- DELETE /api/delete
  - 204 deletes file or folder at the specified path

Payload Scheme
```
{
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://github.com/paulmj7/hermes/blob/master/LICENSE)
