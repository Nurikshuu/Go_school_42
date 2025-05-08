#!/bin/bash
interview=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f 2)
echo $interview
cat interviews/interview-$interview
echo $MAIN_SUSPECT
