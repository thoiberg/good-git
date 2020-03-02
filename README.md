# good-git
An opinionated wrapper around git to provide a more friendly experience. Makes common commands super easy. Does not replace the real git cli for more complex commands.


## Running it
Build the executable and then run it:
```
$ go build gg.go
$ ./gg
```

## Command Line Interface

Good Git will (eventually from Homebrew) be available under the binary `gg`. These are the available commands:

## `gg grab`

Checkout a remote branch and switch to it locally, tracking the remote with the same branch name.

## `gg sync`

Pull any commits from your remote branch into your local, and push any local commits to the remote.

## `gg show`

Show branches. Local branches are shown first, then the remote branches under that (except the branches which are tracked by already-local branches). Each branch will have a number next to it which you can type to `grab` it

## `gg save "<commit message>"`

Commit all working files with a given commit message. Usually followed by `gg sync`

## `gg status` (or `gg huh`)

Show the status of the git repo right now.
