#!/bin/bash

declare -a data=($(<input))

((total=0))
((pos=0))

function get_num {
    num=${data[$pos]}
    ((pos++))
}

function read_node {
    local i

    # Header
    get_num
    local num_children=$num

    get_num
    local num_meta=$num

    # Get children
    for ((i=0; i<$num_children; i++)); do
        read_node
    done

    # Get metadata
    for ((i=0; i<$num_meta; i++)); do
        get_num
        local meta=$num

        ((total+=$meta))
    done
}

read_node

echo $total
