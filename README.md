# NetMon

This is a project that looks at viewing the network connections on a computer using netstat.

### How to start

- To run the project use...

```bash
go run main.go
```

- To exit use...

```bash
ctrl + c
```

## Considerations

#### Thursday 3rd October 2024

I already have a potential issue. Not all operating systems have `netstat` out the box so I may have to use another command to check what operating system their using and ensure that they
have `net-tools` installed into their linux distros.
