# Hermes

[![Build Status](https://travis-ci.com/paulmj7/hermes.svg?branch=master)](https://travis-ci.com/paulmj7/hermes)

Hermes is an filesystem management api to relay the contents of a volume in real time and provide system controls over HTTP. When paired with [Homebase](https://github.com/paulmj7/homebase), they provide a distributed, multivolume storage server.

From [Wikipedia](https://en.wikipedia.org/wiki/Hermes), "Hermes is considered the herald of the gods ... Hermes functioned as the emissary and messenger."

## Installation

Running with Docker =>
```bash
docker build -t hermes .
docker run -p <PORT>:3000 hermes
```

Using the source code =>

```bash
go get github.com/paulmj7/hermes/hermes
```

```go
import "github.com/paulmj7/hermes/hermes"

func main() {
  port := ":3000"
	roots := []string{"desired/path/to/folder"}
	hiddenFiles := make(map[string]bool)
	hiddenFiles["block/this/path"] = true
	corsEnabled := false
	w := hermes.Worker{port, roots, hiddenFiles, corsEnabled}
	w.Serve()
}
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
