# genv

`genv` automates the setup of my dev environment. In its current version, it tracks active project directories and creates tmux sessions based on a bash script.

# Install

```sh
go install github.com/JustinLi007/genv/cmd/genv@latest
```

# Usage

```sh
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

# Scripts

`genv` uses scripts that live in `~/.genv/action/`. The only script that `genv` will look for in its current version is `tmux.sh`. Add your own or start with the ones here [scripts](./scripts/action/).

# Autocompletion

Use the provided autocompletion script [here](./genv_completion)

```sh
# download the script and write it to /etc/bash_completion.d/
curl -LO https://raw.githubusercontent.com/JustinLi007/genv/main/genv_completion
cat genv_completion | sudo tee /etc/bash_completion.d/genv >/dev/null
```
