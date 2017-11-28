ipify-appengine
===============

[![Travis](https://img.shields.io/travis/nrh/ipify-appengine/master.svg?style=flat&label=build)]()
[![Travis](https://img.shields.io/travis/nrh/ipify-appengine/production.svg?style=flat&label=deploy)]()

A fork of [@rdegges'](https://twitter.com/rdegges) [ipify](https://www.ipify.org/) with IPv6 support, hosted
on [Google App Engine](https://cloud.google.com/appengine)

Accessible via `http://api.ipify.io` and `https://api.ipify.io`


Feel free to use and abuse.


Examples
--------
```console
$ curl 'http://api.ipify.io'
2001:0db8:85a3:0000:0000:8a2e:0370:7334

$ curl 'https://api.ipify.io'
2001:0db8:85a3:0000:0000:8a2e:0370:7334

$ curl 'https://v4.ipify.io'
172.16.0.99

$ curl 'https://v4.ipify.io/?format=json'
{"ip":"172.16.0.99"}

$ curl 'https://v4.ipify.io/?format=jsonp'
callback({"ip":"172.16.0.99"});

$ curl 'https://v4.ipify.io/?format=jsonp&callback=asdf'
asdf({"ip":"172.16.0.99"});

$ curl -v -H'Origin: www.example.com' 'https://v4.ipify.io/?format=jsonp&callback=asdf' 2>&1 |egrep -i 'origin'
> Origin: www.example.com
< vary: Origin
< access-control-allow-origin: www.example.com
```
