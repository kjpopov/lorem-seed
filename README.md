# Lorem Seed

Lorem Seed is a command-line tool written in Go for simulating stress/load against APIs and testing database clusters under high load. It can be used to seed data or perform load/stress tests by hitting a specified API endpoint with customizable payloads (currently limited to title and description).

## Installation

To install Lorem Seed, you will need to have Go installed on your system. Then, you can use the `go get` command to install the package:
```
go get github.com/kjpopov/lorem-seed
```

## Usage

To use Lorem Seed, you can invoke the `lorem-seed` command with the following flags:

- `-u` or `--url`: the URL of the API endpoint to hit (defaults to `http://localhost:1323/api/todo`)
- `-i` or `--interval`: the interval between requests, in seconds (defaults to `10`)
- `-p` or `--print-response`: a flag to indicate whether the response from the server should be printed

For example, to perform a load/stress test with a request interval of 5 seconds and printing the server responses:

```
lorem-seed -u http://example.com/api/todo -i 5 -p
```

## Contributions

Contributions are welcome! If you have ideas for improvements or want to report a bug, please open an issue or submit a pull request.

## Limitations
Note that Lorem Seed is designed to generate placeholder text that is similar to the original "lorem ipsum" text, but it is not a perfect replica. The generated text may differ slightly in length and word choice.
Currently, the tool only allows for submission of the title and description fields as part of the payload. To use this tool, you will need an API endpoint that can accept these fields. In the future, the ability to customize the payload will be added.