Web FileSystem
==============

# Plan #

## General idea ##

What I want is this:
Access the web from the files as if the entire web is mounted on my
machine. If I were to access `$MOUNTPATH/example.com` I would see the
content of the page. If were to open
`$MOUNTPATH/$SOMEDOMAIN/somefile.pdf` I should be able to read the
file. 

Certain other things could be done like having gobuster be used to
enumerate a certain path or like writing to a file could be some sort
of post request, but those are beyond the scope of what I wanna do
now; for now I just wanna be able to read from 

## Language choice ##
Rust might be easier for me, but I see a lot of golang jobs so imma go
with golang instead lolz.

C would've been nice, but I want a memory-safe language since I'm
usually not careful when it comes to buffer overflows and such.

## Tooling ##
Nix and/or docker. I'm also not that familiar with go's ecosystem, so
that'll be fun 🙂.
