#compdef camille

local -a pages oses
pages=$(camille --list-all)

_arguments \
    '(-vs r)'{-vs Renekton}'[display Renekton info]' \
    "*:page:(${pages})" && return 0
