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

# for filename in $(find . -name "*.*")
for filename in $(git ls-files)
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
      sed -i '' "s/author_name/$author/g" $filename
      sed -i '' "s/project_description/$description/g" $filename
      sed -i '' "s/PROJECT_NAME/$name_upper/g" $filename
      sed -i '' "s/project_name/$name/g" $filename
    fi
    echo "Renamed $filename"
done

mv configs/.env-example configs/.env

# This command runs only once on GHA!
rm -rf .github/template.yml
