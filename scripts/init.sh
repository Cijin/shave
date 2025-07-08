#! /bin/sh
echo -n "Enter new go package name: "
read pkg_name

go mod edit -module $pkg_name

escaped_pkg_name=$(echo $pkg_name | sed 's/\//\\\//g')
find . \( -name "*.go" -o -name "*.templ" \) -print0 | xargs -0 sed -i -s "s/shave/${escaped_pkg_name}/g"

go get -u github.com/a-h/templ@latest
go get -u github.com/clerk/clerk-sdk-go/v2@latest
go mod tidy

cp env.example .env

rm -rf .git
git init
git add .
git commit -m "Batman"
