hidentd
=======
Yet another spoofing identd, because Go programmers really like reinventing the
wheel.

Mostly (ok, entirely) a go rewrite of https://github.com/cooper/fake-identd

The name is a portmanteau of hide and identd.

Installation
------------
hidentd is installed in the usual way
```bash
go install github.com/magisterquis/hidentd
```
Compiled binaries can be made available on request.  I'm usually on Freenode
with the nick `magisterquis`.

Running
-------
For the most part, hidentd won't need any options.  Probably a good idea to run
with a `-h` the first time to see the knobs.  Of note, `-u` changes the
username sent back to the requestor.

There's is no call to `daemon(3)` or anything fancy like that.
```bash
nohup ./hidentd >hidentd.log 2>&1 &
```

Networking
----------
By default, hidentd listens on port 61113.  This can be changed with the `-a`
option.  Please don't run it as root listening on port 113.  Instead, have it
listen not as root on a port > 1024 and forward 113 with your firewall.

Granted, it's a really simple program, but the less running as root the better.
