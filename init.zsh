unsetopt prompt_subst
export SHELBY_SHELL="zsh"
precmd() {
PROMPT=$'%(?.%{\e[92m%}.%{\e[91m%})'❯'%f'
}