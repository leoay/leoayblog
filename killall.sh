allPids=`pidof chrome`
# echo $allPids
arr=`echo $allPids | tr ' ' ' '`
# echo $arr
for var in $arr;
do
    echo $var
    kill -9 $var
done