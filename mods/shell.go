package mods

var zsh = `autoload -Uz add-zsh-hook
prompt_shelby_cmd() { dir=$(/usr/local/bin/shelby info) }
add-zsh-hook precmd prompt_shelby_cmd
PROMPT=$'$dir\n%(?.%{\e[92m%}.%{\e[91m%})❯%f'`

var bash = `prompt_shelby_load() {
if [ $? != 0 ]; then
    local prompt_symbol="\[\e[0;91m\]❯\[\e[0m\]"
else
    local prompt_symbol="\[\e[0;92m\]❯\[\e[0m\]"
fi

PS1="$(/usr/local/bin/shelby info)\n${prompt_symbol} " 
}
PROMPT_COMMAND=prompt_shelby_load
`

//UseShell checks for the Shell Nature and returns the Structure
func UseShell(shell string) string {
	shells := map[string]string{
		"zsh":  zsh,
		"bash": bash,
	}
	return shells[shell]
}
