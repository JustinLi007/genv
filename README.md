# genv

`genv` automates the setup of my dev environment. In its current version, it tracks active project directories and creates tmux sessions based on a bash script.

# Install

```sh
go install github.com/JustinLi007/genv/cmd/genv@latest
```

# Usage

```sh-session
# display usage info
genv --help

# create tmux session at the current directory
genv tmux -d .

# create a tmux session using an absolute path
genv tmux -d /home/user/workspace/foo

# list marked directories
genv projects

# mark a directory
genv --new projects -d /home/user/workspace/foo

# create a tmux session with a marked directory
genv projects
0   ->  /home/user/workspace/foo
...
5   ->  /home/user/workspace/wow
genv projects -d 5 -e # tmux session working directory is /home/user/workspace/wow
```
