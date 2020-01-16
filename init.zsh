unsetopt prompt_subst
precmd() {
PROMPT="$'%(?.%{\e[92m%}.%{\e[91m%})'❯'%f'"
}