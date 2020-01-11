<p align="center"><img src="/assets/shelby_logo.png" width="350px"/></p>

<div align="center">

[![Build Status](https://travis-ci.org/athul/shelby.svg?branch=master)](https://travis-ci.org/athul/shelby) ![](https://github.com/athul/shelby/workflows/Starboy/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/athul/shelby)](https://goreportcard.com/report/github.com/athul/shelby)

</div>

Shelby is a fast ⚡️ ,lightweight ☁️ ,minimal✨, shell prompt written in Pure Go. 

![](assets/shelby.gif)

-------
## Installation
Follow the steps below, and **Post Installation Instructions**
- Run
```bash
$ sh -c "$(curl -sL https://git.io/ishelby)"	
```
----
OR
- Download the binary from the [WorkFlow Artifacts](https://github.com/athul/shelby/actions?query=workflow%3A%22Go+Build%22).
- You might wanto to make the binary executable, run `chmod 775 <binary_name>`
- Move the Binary to `/usr/local/bin/`

Build From Source

- Clone the Repo
- Run `go get -v`
- Run `go build`
- Move the Binary to `~/.local/bin/shelby`

## Post Installation Instructions

> You must have to add the below code to either the Bashrc or Zshrc file inorder to Shelby to fully work.

The following usage examples are just one example of how Shelby can be
configured. The examples below will result in a setup similar to the one shown
in the demo above: the prompt symbol (`❯`) changes to red if the previous
command exited with an error.

### Bash

Add this to your `.bashrc` file:

```bash
prompt_shelby_load() {
if [ $? != 0 ]; then
    local prompt_symbol="\[\e[0;91m\]❯\[\e[0m\]"
  else
    local prompt_symbol="\[\e[0;92m\]❯\[\e[0m\]"
  fi

  PS1="$(~/.local/bin/shelby)\n${prompt_symbol} " 
}
PROMPT_COMMAND=prompt_shelby_load
```

### Zsh

Add this to your `.zshrc` file:

```zsh
autoload -Uz add-zsh-hook
prompt_shelby_cmd() { ~/.local/bin/shelby }
add-zsh-hook precmd prompt_shelby_cmd
prompt_symbol='❯'

PROMPT=$'%(?.%{\e[92m%}.%{\e[91m%})${prompt_symbol}%f'
```


### Extra Bits
- `[+]` shows if you've got Untracked Files
  - `[2+]` shows if you've got 2 untracked files
- `[!]` shows if you've got Unstaged Files
  - `[3!]` shows if you've got 3 unstaged files
- Dispalys any VirtualENVs you're working in
- Displays the Current Git Branch
- Displays **Username** and **Hostname** of the machine while in SSH
- Small Size(~=2MB)


### Inspired From
- [StarShip by @matchai](https://starship.rs)
- [SpaceShip by @denysdovhan](https://github.com/denysdovhan/spaceship-prompt)

### **Code reused from** :heart:
- [Mímir by @talal](https://github.com/talal/mimir)
- [Powerline-Go by @justjanne](https://github.com/justjanne/powerline-go/)

## Support My work
<a href="https://www.buymeacoffee.com/JeVlc7T" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" style="height: 20px !important;width: 217px !important;" ></a>

**Logo Made by** [Sreeram aka @fillerink](https://github.com/fillerink)

The Installation Script can be found at https://github.com/athul/autom
