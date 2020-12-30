#!/bin/bash

exec 10<&0

    echo "Building Go files..."
    eval "go build -o bin/trail -v ." 
    if [ $? -eq 0 ] 
     #Notifies the user that a the openssl command has evaluated with the key $LINE$i
     then echo "Build Complete" 
    fi
    echo "Tidying...."
    eval "go mod tidy" 
    if [ $? -eq 0 ] 
     #Notifies the user that a the openssl command has evaluated with the key $LINE$i
     then echo "Tidied." 
    fi
    echo "Updating dependencies...."
    eval "go mod vendor" 
    if [ $? -eq 0 ] 
     #Notifies the user that a the openssl command has evaluated with the key $LINE$i
     then echo "Updated." 
    fi
    echo "Pushing changes to GitHub with message 'Fixes and Updates'"
    eval "git add ."
    eval "git commit -m 'Fixes and Updates'"
    eval "git push" 
    if [ $? -eq 0 ] 
     #Notifies the user that a the openssl command has evaluated with the key $LINE$i
     then echo "Done." 
    fi
    exit 1

# restore stdin from filedescriptor 10
# and close filedescriptor 10
exec 0<&10 10<&-
