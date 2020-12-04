#!/bin/bash

declare -a data=($(<input))

((pos=0))
((total=0))

function get_num {
    num=${data[$pos]}
    ((pos++))
}

function read_node {
    local i
    local nodes
    local current_node=0
    local my_total=0

    declare -a nodes

    # Header
    get_num
    local num_children=$num

    get_num
    local num_meta=$num

    # Get children
    for ((i=0; i<$num_children; i++)); do
        read_node
        nodes[$((i+1))]=$total
    done

    # Get metadata
    for ((i=0; i<$num_meta; i++)); do
        get_num
        local meta=$num

        if [ $num_children -eq 0 ]; then
            ((my_total+=$meta))
        else
            if [ $meta -ne 0 ]; then
                local child_total=${nodes[$meta]}

                if [ -n "$child_total" ]; then
                    ((my_total=$my_total+$child_total))
                fi
            fi
        fi
    done

    total=$my_total
}

read_node

echo $total
