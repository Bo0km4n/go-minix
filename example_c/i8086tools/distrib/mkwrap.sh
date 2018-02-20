# sh mkwrap.sh [-x] /usr/local/bin/7run /usr/local/minix2 /usr/local/bin/m2 usr/bin/cc

opt=""
if [ "x$1" = "x-x" ]
then
    opt=$1
    shift
fi

# ex. /usr/local/bin/7run
run=$1

# ex. /usr/local/minix2
root=$2

# ex. /usr/local/bin/m2cc
cmd=$3`basename $4`

# ex. /usr/local/minix2/usr/bin/cc
wrap=$2/$4

shift; shift; shift; shift

cat << EOT > $cmd
#!/bin/sh
opt=""
if [ "x\$1" = "x--s" ]; then opt="-s"; shift; fi
if [ "x\$1" = "x--v" ]; then opt="-v"; shift; fi
if [ "x\$1" = "x--m" ]; then opt="-m"; shift; fi
EOT

if [ "x$opt" = "x-x" ]
then
    echo "$run \$opt $@ -r $root $wrap \$@" >> $cmd
    echo "exit 0" >> $cmd
else
    echo "exec $run \$opt -r $root $wrap \$@" >> $cmd
fi

chmod 755 $cmd
