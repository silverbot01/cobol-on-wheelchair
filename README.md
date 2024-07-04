COBOL on Wheelchair
===================

Micro web-framework for COBOL (in statu nascendi)

OK, [COBOL-ON-COGS](http://www.coboloncogs.org/HOME.HTM) was funny, but why not try that for real? ;)

Things you'll need to run _COBOL on Wheelchair_:

* [GNU Cobol](http://sourceforge.net/projects/open-cobol/) (```sudo apt-get install open-cobol```)
* Ability to run cgi-bin on Apache server


Installation
-------------

```
git clone https://github.com/azac/cobol-on-wheelchair/
cd cobol-on-wheelchair
./downhill.sh
```

Tutorial
--------

https://github.com/azac/cobol-on-wheelchair/blob/master/tutorial/index.md

Personal Additions
------------------

*Added simple Go HTTP server utilizing "net/http/cgi" to call the "the.cow" executable as a handler. Setup for port 8080 serving files over /static from the static directory.

*Moved images directory to /static/. Updated apache.conf's DocumentRoot to this. (the Go server doesn't rely on Apache, but I did this to not break things)

*main.go is the Go source code, go.mod is the dependency organization file (no third-party dependencies needed), either use 'go run main.go' or run a built executable from the same directory as 'the.cow'.
