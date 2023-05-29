Web FileSystem (WIP)
==============

Mount the web as a filesystem.

# Plan #
(Move this to an `.org` file plzzz)

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

## Cool things to add ##
What about protocols other than http(s)?
ipfs and ftp would be good ones to add, or maybe even gemini?
Another thing question I raise here is: where does the user specify
the protocol? Is it `.../ftp/someurl/somefile`? What about *guessing*
what the protocol is? Like `.../some-ipfs-address` and we'd get the
content.

Another nice-to-have would be writing. Uploading to ftp servers as
anonymous or registering to ipfs when Kubo is already set up would be
pretty cool.

## Language choice ##
Rust might be easier for me, but I see a lot of golang jobs so imma go
with golang instead lolz.

C would've been nice, but I want a memory-safe language since I'm
usually not careful when it comes to buffer overflows and such.

## Tooling ##
Nix and/or docker. I'm also not that familiar with go's ecosystem, so
that'll be fun 🙂.
