# Lets talk about Logging

Principles

- dont log in libs, return errors or handle them not log them in libs
- But I have a hight level lib and I must log. Then use interfaces!
- Use a save method in main to exit and respect the defer calls (dont use log.Panic/log.Fatal)
- dont have dependencies at compile time in your core libs for logging (interface to the rescue)
- One place to set logger settings
- Use structured logs to query them
- No warning log level: Nobody reads warnings, because by definition nothing went wrong.

## Task description

[Go] Improve logging principle
We believe that improve the logging mechanism decouple dependency to a special logging lib.
it will reduce dependencies in our core code and increase the maintainability and will improve our understanding of a good principle to log in core libs.
We are done if we implement a general logging interface which implements different logging libs (logrus).

## Dependencies

Remove dependencies in your core libs for logging. We implement a interface for the core libs. So its by the caller which logger he will use.
The logger must be implemented and have the function of the interfaces, like logrus and zap. When you implement your cmd programm you can choose which caller you need.

## No Panic function

Because we handle all our Log functionality in one Interface we can allow or disallow some functions of the underlying lib.
For example the logrus lib has a log.Panic function wich where catched by the lib and then exit with a os.Exit and not respect our defer calls.
So we will not Implement this function in the interface to have a clean exit in our defer main function. When you will exit with a log call and have a clean exit, use the log.Fatal function call in your lib.

## The disaster with logging

- package initilized logger, must be build on compile time, one global logger, you dont see where he exactly come from, runtime errors if not init (Use noop logger)
- package struct, must be init from the caller, canot use global in libs
- logs in tests are bad, use a in memory or noop logger

## Discussion

- dont log in libs, caller stuff
- global Package logger? (noop or arg), no (no global vars, no magic)
- three log levels? no Warning
- With fields? use json output!

- defined log entries (const strings)? not here, use a error package to define errors (dependency?)
- in memory logging for testing?
- default format and output type/json?
- errors.Wrap, yes but only when you use third parity libs and you dont nothing about the error.

## Detecting Problems

- global logrus package logger
- info logging in libs but it is not nessacary
- logging errors in libs
- logging WithFields only for Error, why?
- Only one log.Info in lib and dependency for logrus!
- go test (archive_test.go) used one line logrus, use t.Log
- eks.go using log not logrus
- splunk.go printf error not using log.Errorf

## Ressources

[https://dave.cheney.net/tag/logging](Dave Cheney logging)

[https://dave.cheney.net/2017/01/26/context-is-for-cancelation](dont use context to log)

[https://medium.com/@jfeng45/go-microservice-with-clean-architecture-application-logging-b43dc5839bce](clean app logging)

[https://medium.com/@jfeng45/go-microservice-with-clean-architecture-a08fa916a5db](clean log arch)

[https://github.com/jfeng45/servicetmpl](mircoservices with a good log arch)

[https://www.mountedthoughts.com/golang-logger-interface/](another golang-logger)

[github.com/amitrai48/logger](amitrai48 logger)

[https://www.datadoghq.com/blog/go-logging/](datadog talk about go-logging)

[https://github.com/go-log/log](go-log is a reference impl)

[https://github.com/go-logr/logr](next step impl of go-log)

[https://peter.bourgon.org/blog/2017/06/09/theory-of-modern-go.html](no magic in libs)
