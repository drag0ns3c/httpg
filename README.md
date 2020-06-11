# HTTPG

HTTPG is a command line utility to expose a directory over http. This can be useful when a target machine doesn't have any tools installed natively.

It is distributed as a standalone binary with no dependencies and can be run on all platforms.

As a bonus, it will render a status page that will output details of the machine.

## Usage

The following command will launch the file server listening on `0.0.0.0:9000` and will serve the directory the command is run from.
```bash
./httpg -port 9000 -host 0.0.0.0 -dirPath .
```

You can also find info on the host machine by visiting `/__info__`.

## Options

| Option | Default | Description |
| ------ | ------- | ----------- |
| `dirPath` | `.` | The path to serve files from |
| `host` | `0.0.0.0` | The host address to listen on |
| `port` | `8080` | The port number to listen on |
| `help` | | Display the help text