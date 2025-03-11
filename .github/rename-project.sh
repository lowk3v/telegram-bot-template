#!/usr/bin/env bash
while getopts a:n:u:d: flag
do
    case "${flag}" in
        a) author=${OPTARG};;
        n) name=${OPTARG};;
        d) description=${OPTARG};;
    esac
done

echo "Author: $author";
echo "Project Name: $name";
echo "Description: $description";

echo "Renaming project..."

name_upper="$(echo $name|tr 'a-z' 'A-Z')"

for filename in $(grep -liroE '(author_name|project_description|project_name)' .)
do
    if [[ $filename == *".github"* ]]; then
        continue
    fi
    if [[ $(uname) == "Darwin" ]]; then
      sed -i '' -e "s/author_name/$author/g" $filename
      sed -i '' -e "s/project_description/$description/g" $filename
      sed -i '' -e "s/PROJECT_NAME/$name_upper/g" $filename
      sed -i '' -e "s/project_name/$name/g" $filename
    else
      sed -e "s/author_name/$author/g" $filename
      sed -e "s/project_description/$description/g" $filename
      sed -e "s/PROJECT_NAME/$name_upper/g" $filename
      sed -e "s/project_name/$name/g" $filename
    fi
    echo "Renamed $filename"
done

# This command runs only once on GHA!
rm -rf .github/template.yml
