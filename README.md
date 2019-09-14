httppconntime
=============

httppconntime is detect HTTP persistent connection time(keep-alive time).

Installation
------------

```shell script
go get github.com/kjmkznr/httppconntime/cmd/httppconntime/
```

Usage
-----

Detect keep-alive time between 60s and 300s.

```
$ httppconntime -url http://localhost:8080/ -init 60s -max 300s
2019/09/14 07:57:55 Start probe HTTP Persistent Connection Time, between 1s and 12s
2019/09/14 07:57:55 Probe 0s - 6s
2019/09/14 07:58:01 Probe 6s - 3s
2019/09/14 07:58:04 Probe 3s - 4s
2019/09/14 07:58:08 Probe 4s - 5s
2019/09/14 07:58:13 HTTP Persistent Connection Time(KeepAlive Time) = 5s
```

CLI Options
-----------

```
  -init duration
        Initial wait time (default 1s)
  -max duration
        Max wait time (default 5m0s)
  -url string
        Target URL. like http://www.example.jp/
```

