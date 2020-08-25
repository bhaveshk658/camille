#compdef camille

_camille() {
    local state

    _arguments \
        '1: :->vs'

    case $state in
        (vs) _arguments: '1:profiles:(-vs)';;
         (*) compadd "$@" Renekton Darius Sett
    esac
}

_camille "$@"
