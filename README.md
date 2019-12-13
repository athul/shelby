# Shelby [![Build Status](https://travis-ci.org/athul/shelby.svg?branch=master)](https://travis-ci.org/athul/shelby) ![](https://github.com/athul/shelby/workflows/StarBoy/badge.svg)

Shelby is a fast :rocket: ,lightweight 💨 ,minimal🧸, shell prompt written in Pure Go. 

![](/shelby.gif)

-----
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
    local prompt_symbol="\[\e[0;31m\]❯\[\e[0m\]"
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
PROMPT="%(?.%F{magenta}.%F{red})${prompt_symbol}%f "
```

### Its Still in WIP :sweat_smile:

**Most of the code is derived from** :heart:
- [Mímir by @talal](https://github.com/talal/mimir)
- [Gofu by @majewsky](https://github.com/majewsky/gofu#prettyprompt)
