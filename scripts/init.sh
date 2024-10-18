#! /bin/sh
echo -n "Enter new go package name: "
read pkg_name

# update go mod
go mod edit -module $pkg_name

escaped_pkg_name=$(echo $pkg_name | sed 's/\//\\\//g')

find . -name "*.go" -print0 | xargs -0 sed -i -s "s/shave/${escaped_pkg_name}/g"
