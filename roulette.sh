#!/bin/bash 
while getopts s:i: flag
do
    case "${flag}" in
        s) SOURCE=${OPTARG};;
        i) INDEX=${OPTARG};;
    esac
done

if [ -n "${SOURCE}" ] && [ -z "${INDEX}" ]
    then
    SUM=$(find $SOURCE -type f -print | wc -l)
    echo "Total number of files in provided folder: $SUM"
    echo "I'm going to choose a random file now because you didn't specify an index. (-i number)"
    echo
    INDEX=$(shuf -i 1-$SUM -n 1)
fi

if [ -n "${INDEX}" ] && [ -z "${SOURCE}" ]
    then
    echo "You need to provide the source folder where files will be looked through. (-s source_folder/)"
    exit 1
fi

if [ -n "${SOURCE}" ] && [ -n "${INDEX}" ]
    then
    echo "Your roulette file iiiiiiiiiiis:"
    find $SOURCE -type f -print | head -n $INDEX | tail -n 1
    exit 0
fi

echo "I need some input you know? (-s source_folder/ and -i number)"
