INTERVIEW_NUMBER=$(head -n 179 streets/Buckingham_Place | tail -n 1 | tr "#" : | cut -d : -f2)
echo $INTERVIEW_NUMBER
# echo $(find . -type f -name "*$INTERVIEW_NUMBER*" -exec cat {} \;)
cat interviews/interview-$INTERVIEW_NUMBER
echo $MAIN_SUSPECT