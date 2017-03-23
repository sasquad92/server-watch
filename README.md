# server-watch

server-watch is a small and simple watchdog for a Linux service.

This watchdog will check if >>service<< is running  every X seconds.
If service is down, the watchdog will attmept to start it every Y seconds up to Z times.
Watchdog will write events to a log file.
Additionally watchdog will report fallowing events to provided email address:
* Service is down
* Service has been started after X attemps
* Service can't be started after X attemps

## before You run

Please check config file and fill it with correct data.

## to list services on ubuntu
```
service --status-all
```

## daemon

If You want to run this program as a daemon

```
go run main.go & disown
```

or try [daemonize](http://software.clapper.org/daemonize/)

## these articles helped me
* [working with services](https://blog.terminal.com/working-with-services/)
* [sending an email from go program](https://gist.github.com/jpillora/cb46d183eca0710d909a)