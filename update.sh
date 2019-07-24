#!/bin/bash

# This script updates the current repository to the latest version of
# yara.
git submodule init
git submodule update

# Apply patches to submodule tree
cd yara_src/
echo Resetting the yara source tree.
git reset --hard

echo Applying patches.
patch -p1 < ../yara_src.diff
cd -

echo Copying files to golag tree.
cp yara_src/libyara/*.c .
cp yara_src/libyara/*.h .
cp yara_src/libyara/include/yara.h .
for i in yara_src/libyara/include/yara/*.h; do
    cp $i yara_`basename $i`
done

for i in yara_src/libyara/modules/{test,pe,elf,math,time,module_list}*; do
    cp $i modules_`basename $i`
done

cp yara_src/libyara/proc/linux.c proc_linux.c
cp yara_src/libyara/proc/windows.c proc_windows.c
cp yara_src/libyara/proc/mach.c proc_darwin.c

sed -i 's/yara\//yara_/g' *.h *.c
sed -i 's/modules\//modules_/g' *.h *.c
