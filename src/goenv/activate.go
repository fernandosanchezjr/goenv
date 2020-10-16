package main

const activateScript string = `#!/bin/bash

PARENT_PATH="$( cd "$( dirname $( dirname "${BASH_SOURCE[0]}" ) )" && pwd )"

if [[ "${PATH}" == *"go/bin"* ]]; then
    PATH=$PATH:/usr/local/go/bin
    export PATH
fi

if [ "${GOROOT}" == "" ]; then
    GOROOT=/usr/local/go
    export GOROOT
fi

if [ "${GO_ENV}" == "" ]; then
    GO_ENV=${PARENT_PATH}
    echo Activating goenv in ${GO_ENV}
    export GO_ENV

    GOPATH=${GO_ENV}
    export GOPATH

    BIN_PATH=${GOPATH}/bin
    OLD_PATH=${PATH}
    PATH=${PATH}:${BIN_PATH}
    export OLD_PATH
    export PATH
elif [ "${GO_ENV}" != ${PARENT_PATH} ]; then
    GO_ENV=${PARENT_PATH}
    echo Switching to goenv in ${GO_ENV}
    GOPATH=${GO_ENV}
    export GOPATH

    BIN_PATH=${GOPATH}/bin
    PATH=${OLD_PATH}:${BIN_PATH}
    export PATH
else
    echo Already activated goenv in ${GO_ENV}
fi

get_deps () {
    go get >/dev/null 2>&1
}

install_sh () {
    LS_RESULTS=` + "`" + `ls -1 *.sh 2>/dev/null | wc -l` + "`" + `
    if [ "${LS_RESULTS}" != "0" ]; then
        chmod 755 *.sh
        cp *.sh ${BIN_PATH}/
    fi
}

build () {
    get_deps
    BUILD_RESULT=` + "`" + `go install 2>&1` + "`" + `
    if [ "${BUILD_RESULT}" != "" ]; then
        echo "${BUILD_RESULT}"
    else
        install_sh
    fi
}

list_installables () {
    LS_RESULTS=` + "`" + `ls -1 ${BIN_PATH}/* 2>/dev/null | grep -v activate` + "`" + `
    for result in ${LS_RESULTS}; do
        echo ${result##*/}
    done
}

install_to_system () {
    INSTALLABLES=` + "`" + `list_installables` + "`" + `
    for installable in ${INSTALLABLES}; do
        sudo install -m 755 ${BIN_PATH}/${installable} /usr/local/bin/${installable}
    done
}

make_install () {
    BUILD_RESULT=` + "`" + `build_all` + "`" + `
    if [ "${BUILD_RESULT}" != "" ]; then
        echo "${BUILD_RESULT}"
    else
        install_to_system
    fi
}

uninstall_from_system () {
    INSTALLABLES=` + "`" + `list_installables` + "`" + `
    for installable in ${INSTALLABLES}; do
        if [ -e /usr/local/bin/${installable} ]; then
            sudo rm /usr/local/bin/${installable}
        fi
    done
}

make_uninstall () {
    BUILD_RESULT=` + "`" + `build_all` + "`" + `
    if [ "${BUILD_RESULT}" != "" ]; then
        echo "${BUILD_RESULT}"
    else
        uninstall_from_system
    fi
}

deactivate () {
    PATH=${OLD_PATH}
    unset BIN_PATH
    unset GO_ENV
    unset GOPATH
    export PATH
    unset OLD_PATH
    unset get_deps
    unset install_sh
    unset build
    unset list_installables
    unset install_to_system
    unset make_install
    unset uninstall_from_system
    unset make_uninstall
    unset list_buildable_folders
    unset build_all
    unset make_clean
    unset deactivate
}

list_buildable_folders () {
    find ${GO_ENV}/ -iname *.go | xargs -L 1 -r dirname $1 | sort -u
}

build_all () {
    BUILDABLE_FOLDERS=` + "`" + `list_buildable_folders` + "`" + `
    OLD_PWD=` + "`" + `pwd` + "`" + `
    for folder in ${BUILDABLE_FOLDERS}; do
        cd ${folder}
        BUILD_RESULT=` + "`" + `build` + "`" + `
        if [ "${BUILD_RESULT}" != "" ]; then
            echo "${BUILD_RESULT}"
        fi
    done
    cd ${OLD_PWD}
}

make_clean () {
    INSTALLABLES=` + "`" + `list_installables` + "`" + `
    for installable in ${INSTALLABLES}; do
        rm ${BIN_PATH}/${installable}
    done
}
`
