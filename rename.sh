# any suggestion how to avoid this rename and use just replaces in go.mod is welcome
FROM="mvdan.cc/sh"
TO="github.com/openserverless-mvdansh-fork"
if test -n "$1"
then 
	TMP="$FROM"
       	FROM="$TO" 
	TO="$TMP"
fi
find . \( -name \*.go -o -name go.modi \) | while read file 
do echo $file
   sed -i "s!$FROM!$TO!" $file 
done
go mod tidy
