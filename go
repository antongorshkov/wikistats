#!/bin/bash

i="1"

while [ $i -lt 5 ]
do
url="https://dumps.wikimedia.org/other/pagecounts-raw/2016/2016-0$i/md5sums.txt"
echo $url
wget -q $url
cat md5sums.txt | grep pagecounts | cut -d" " -f3 > md5names.txt
while read p; do
    data_url="https://dumps.wikimedia.org/other/pagecounts-raw/2016/2016-0$i/$p"
    unzipped=${p/.gz/}
    if [ ! -f $unzipped.top ]; then
        echo $unzipped
        wget -q $data_url
        gunzip $p
        cat $unzipped | gawk -F ' ' '$1=="en" && $3>100' | grep -v 'Special:' | grep -v 'File:' > $unzipped.top
        rm $unzipped
    fi
done < md5names.txt
rm md5sums.txt
rm md5names.txt
i=$[$i+1]
done
