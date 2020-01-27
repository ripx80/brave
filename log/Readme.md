# Lets talk about logging

This is a discussion MR for logging principles and interfaces. Feel free to comment the principles and implementation and the questions in the Discussion section. I collect opportunities around the globe for this, you can find them in the Further Reading section.

## Principles

- dont log in libs, return errors or handle them not log them in libs
- But I have a hight level lib and I must log. Then use interfaces!
- Use a save method in main to exit and respect the defer calls (dont use log.Fatal)
- dont have dependencies at compile time in your core libs for logging (interface to the rescue)
- One place to set logger settings
- Use structured logs to query them
- No warning log level: Nobody reads warnings, because by definition nothing went wrong.

## Dependencies

Remove dependencies in your core libs for logging. We implement a interface for the core libs. So its by the caller which logger he will use.
The logger must be implemented and have the function of the interfaces, like logrus and zap. When you implement your cmd programm you can choose which caller you need.

## No Fatal function

Because we handle all our Log functionality in one Interface we can allow or disallow some functions of the underlying lib.
For example the logrus lib has a log.Fatal function wich where exit with a os.Exit and not respect our defer calls.
So we will not implement this function in the interface to have a clean exit in our defer main function (SafeExit). When you will exit with a log call and have a clean exit, use the log.Panic function call in your lib.

## The disaster with logging

- package inits logger (instances or global vars), must be build on compile time, one global logger, you dont see where he exactly come from
- if you use a package struct it must be init from the caller and cannot use global in libs, you must pass it by argument (clear solution with overhead)
- how we can test logs in our tests?
- No pre defined log msg, for example like errors "EOF" to check against

## Solution

- logger package which implements a logging interface
- package instance named "log" for same use like logrus (yes, its magic, not explicit)
- init a noop logger per default (silent mode, or when you will test without a big logrus/zap)
- pre defined log messages (only examples at the moment)
- memlog implements a in memory logger for testing!
- full control over the logger you choose.

## Discussion

- dont log in libs, caller stuff
- global Package logger? no (no global vars, no magic)?
- three log levels? no Warning, Panic, only Debug,Info,Error
- With fields? use json output!?
- defined log entries (const strings)? not here, use a error package to define errors (dependency?)
- default format and output type/json?
- errors.Wrap, yes but only when you use third parity libs.

## Further Reading

[no magic in libs](https://peter.bourgon.org/blog/2017/06/09/theory-of-modern-go.html)

[Dave Cheney logging](https://dave.cheney.net/tag/logging)

[dont use context to log](https://dave.cheney.net/2017/01/26/context-is-for-cancelation)

[clean app logging](https://medium.com/@jfeng45/go-microservice-with-clean-architecture-application-logging-b43dc5839bce)

[clean log arch](https://medium.com/@jfeng45/go-microservice-with-clean-architecture-a08fa916a5db)

[structured logging in golang](https://www.client9.com/structured-logging-in-golang/)

[another golang-logger](https://www.mountedthoughts.com/golang-logger-interface/)

[datadog talk about go-logging](https://www.datadoghq.com/blog/go-logging/)

[amitrai48 logger](https://github.com/amitrai48/logger)

[mircoservices with a good log arch](https://github.com/jfeng45/servicetmpl)

[go-log is a reference impl](https://github.com/go-log/log)

[next step impl of go-log](https://github.com/go-logr/logr)