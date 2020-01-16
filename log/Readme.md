# Lets talk about Logging

Principles

- dont log in libs, return errors
- But I have a hight level lib and I must log. Then use interfaces! (The package level logger anti pattern)
- Use a save method to exit and respect the defer calls (log.Panic/log.Fatal)
- dont have dependencies at compile time in your core libs for logging (interface to the rescue)
- One place to set logger settings

## Task description

[Go] Improve fearlessly logging principle
We believe that improve the logging mechanism in fearlessly (https://appsgit.bethel.jw.org/o11n/fearlessly) decouple dependency to a special logging lib.
it will reduce dependencies in our core code and increase the maintainability and will improve our understanding of a good principle to log in core libs.
We are done if we implement a general logging interface which implements different logging libs (logrus).

## Dependencies

Remove dependencies in your core libs for logging. We implement a interface for the core libs. So its by the caller which logger he will use.
The logger must be implemented and have the function of the interfaces, like logrus and zap. When you implement your cmd programm you can choose which caller you need.

## No Panic function

Because we handle all our Log functionality in one Interface we can allow or disallow some functions of the underlying lib.
For example the logrus lib has a log.Panic function wich where catched by the lib and then exit with a os.Exit and not respect our defer calls.
So we will not Implement this function in the interface to have a clean exit in our defer main function. When you will exit with a log call and have a clean exit, use the log.Fatal function call in your lib.

## Ressources

[https://dave.cheney.net/tag/logging](Dave Cheney logging)
[https://dave.cheney.net/2017/01/26/context-is-for-cancelation](dont use context to log)
[https://medium.com/@jfeng45/go-microservice-with-clean-architecture-application-logging-b43dc5839bce](clean app logging)
[https://medium.com/@jfeng45/go-microservice-with-clean-architecture-a08fa916a5db](clean log arch)
[https://github.com/jfeng45/servicetmpl](mircoservices with a good log arch)
[https://www.mountedthoughts.com/golang-logger-interface/](another golang-logger)
[github.com/amitrai48/logger](amitrai48 logger)
https://www.datadoghq.com/blog/go-logging/