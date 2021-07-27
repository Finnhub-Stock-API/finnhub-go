if [ $# -eq 0 ]
  then
    echo "No arguments supplied"
    exit
fi

git add -A
git commit -m "update v$1"
git tag v$1
git push origin v$1
git push
