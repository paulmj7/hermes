# Hermes

Hermes is a an api to relay the contents of a volume in real time, similar to Apache's live file browsing. When paired with Homebase [COMING SOON], this combination provides a distributed, multivolume storage server.

From [Wikipedia](https://en.wikipedia.org/wiki/Hermes), "Hermes is considered the herald of the gods ... Hermes functioned as the emissary and messenger."

## Installation

First, customize the config.json file in the src folder to serve your root volumes, as well as change the port if necessary.

To run Hermes

```bash
cd src
go run .
```

## Usage

- POST /api
  - 200 sends the root volumes being served as JSON
- POST /api/move
  - 200 sends the contents of the current directory as JSON
- POST /api/retrieve
  - 200 sends the url route to download the specified file
  - ex. returns x91js -> localhost:5000/api/send?key=x91js
- GET /api/send?key=
  - 200 streams the file for download

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
