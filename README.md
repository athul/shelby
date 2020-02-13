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

---
OR
For **MacOS** Users, install via _*Homebrew*_

```bash
brew install athul/tap/shelby
```

----
OR
- Download the binary from the [WorkFlow Artifacts](https://github.com/athul/shelby/actions?query=workflow%3A%22Go+Build%22)(Only for Linux)
- You might wanto to make the binary executable, run `chmod +x <binary_name>`
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
PROMPT=$'%(?.%{\e[92m%}.%{\e[91m%})❯%f'
```


### Extra Bits
- Displays the Current Git Branch
- `✔` shows if any staged files are present
- `[+]` shows if you've got Untracked Files
  - `[2+]` shows if you've got 2 untracked files
- `[!]` shows if you've got Unstaged Files
  - `[3!]` shows if you've got 3 unstaged files
- `↑` if your HEAD is ahead
- `↓` if your HEAD is behind
- `⇅` if your HEAD is diverged
- Dispalys any VirtualENVs you're working in
- Displays **Username** and **Hostname** of the machine while in SSH
- Small Size(~=2MB)


### Inspired From
- [StarShip by @matchai](https://starship.rs)
- [SpaceShip by @denysdovhan](https://github.com/denysdovhan/spaceship-prompt)

### **Code reused from** :heart:
- [Mímir by @talal](https://github.com/talal/mimir)
- [Powerline-Go by @justjanne](https://github.com/justjanne/powerline-go/)

## Support My work

<center><p ><img height='100' style='border:0px;height:36px;' src='https://imgix.bustle.com/uploads/image/2019/5/2/ffa82ad4-937e-412c-9bfd-33cb9252e88e-instagram-donate.jpg?w=1020&h=576&fit=crop&crop=faces&auto=format&q=70' border='0' alt='Donations' /></a>
</p>
<h3>Donation: UPI ID: <b>athul8720@oksbi</b></h3>
</center>


### 

**Logo Made by** [Sreeram aka @fillerink](https://github.com/fillerink)

The Installation Script can be found at https://github.com/athul/autom
