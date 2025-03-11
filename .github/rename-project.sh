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

echo "Renaming project... on the platform $(uname)"

name_upper="$(echo $name|tr 'a-z' 'A-Z')"

for filename in $(grep -liroE '(author_name|project_description|project_name)' .)
do
    if [[ $filename == *".github"* ]]; then
        continue
    fi
    if [[ $(uname) == "Darwin" ]]; then
      sed -i '' -e "s/author_name/$author/g" $filename >> /dev/null
      sed -i '' -e "s/project_description/$description/g" $filename >> /dev/null
      sed -i '' -e "s/PROJECT_NAME/$name_upper/g" $filename >> /dev/null
      sed -i '' -e "s/project_name/$name/g" $filename >> /dev/null
    else
      sed -ie "s/author_name/$author/g" $filename >> /dev/null
      sed -ie "s/project_description/$description/g" $filename >> /dev/null
      sed -ie "s/PROJECT_NAME/$name_upper/g" $filename >> /dev/null
      sed -ie "s/project_name/$name/g" $filename >> /dev/null
    fi

    # checking
    [[ $(grep -liroE '(author_name|project_description|project_name)' $filename | wc -l) == 1 ]] \
      && echo "Rename failed for $filename" \
      || echo "Renamed$filename"
done

# This command runs only once on GHA!
rm -rf .github/template.yml
