# Shelby [![Build Status](https://travis-ci.org/athul/shelby.svg?branch=master)](https://travis-ci.org/athul/shelby) ![](https://github.com/athul/shelby/workflows/StarBoy/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/athul/shelby)](https://goreportcard.com/report/github.com/athul/shelby)

Shelby is a fast 🚀 ,lightweight 🎈 ,minimal✨, shell prompt written in Pure Go. 

![](/shelby.gif)

-------
**Here are some Benchmark Results**
![](/b1.png)
![](/b2.png)

-------
## Installation
Initially the Packaging is so dull and I'm really new to this stuff. Follow the steps below,
- Run
```bash
$ sh -c "$(curl -sL https://git.io/getshelby)"	
```
----
OR
- Download the binary from the [Releases](https://github.com/athul/shelby/releases) page
- Move the Binary to `/usr/local/bin/`
- Enjoy :tada:

OR

- Clone the Repo
- Run `go get -v`
- Run `go build`
- Move the Binary to `/usr/local/bin/`

I hope this gets improved soon
### Extra Bits
- `[+]` shows if you've got Untracked Files
  - `[2+]` shows if you've got 2 untracked files
- `[!]` shows if you've got Unstaged Files
  - `[3!]` shows if you've got 3 unstaged files
- Displays **Username** and **Hostname** of the machine while in SSH

## Usage

The following usage examples are just one example of how Shelby can be
configured. The examples below will result in a setup similar to the one shown
in the demo above: the prompt symbol (`❯`) changes to red if the previous
command exited with an error.

### Bash

Add this to your `.bashrc` file:

```bash
prompt_shelby_load() {
  if [ $? != 0 ]; then
    local prompt_symbol="\[\e[1;32m\]❯\[\e[0m\]"
  else
    local prompt_symbol="\[\e[0;35m\]❯\[\e[0m\]"
  fi

  PS1="$(/path/to/shelby)\n${prompt_symbol} " #change it with the original path
}
PROMPT_COMMAND=prompt_shelby_load
```

### Zsh

Add this to your `.zshrc` file:

```zsh
autoload -Uz add-zsh-hook
prompt_shelby_load() { /path/to/shelby } #change it with the original path
add-zsh-hook precmd prompt_shelby_load

prompt_symbol='❯'
PROMPT="%(?.%F{green}.%F{red})${prompt_symbol}%f "
```

### Inspired From
- [StarShip by @matchai](https://starship.rs)
- [SpaceShip by @denysdovhan](https://github.com/denysdovhan/spaceship-prompt)

**Most of the logic is derived from** :heart:
- [Mímir by @talal](https://github.com/talal/mimir)
- [Powerline-Go by @justjanne](https://github.com/justjanne/powerline-go/)
