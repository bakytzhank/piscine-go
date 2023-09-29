#!bin/bash
INTERVIEW_NO=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo $INTERVIEW_NO
cat interviews/interview-"$INTERVIEW_NO"
echo $MAIN_SUSPECT
