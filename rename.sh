# any suggestion how to avoid this rename and use just replaces in go.mod is welcome
FROM="mvdan.cc/sh"
TO="github.com/openserverless-mvdansh-fork"
if test -n "$1"
then 
	TMP="$FROM"
       	FROM="$TO" 
	TO="$TMP"
fi
find . -name \*.go -o -name go.mod | while read file 
do
	if sed -i "s!$FROM!$TO!" $file 
        then echo "fixed $file"
	fi
done
