#!/bin/bash

in_files=$(ls lang/parser/tests/*.zc)
for in_file in $in_files; do
    out_file=$(echo $in_file | sed s/\.zc/\.json/)
    echo "$in_file ==> $out_file"
    go run cmd/zc/main.go parse -no-filename $in_file > $out_file
done

echo $files