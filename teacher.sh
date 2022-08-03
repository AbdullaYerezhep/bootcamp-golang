INTERVIEW_NUMBER=$(head -n 179 streets/Buckingham_street | tail -n 1 | cut -c 16-)
echo $INTERVIEW_NUMBER
echo $(find ./interviews/ -type f -name "*$INTERVIEW_NUMBER*" -exec cat {} \; -print)
export MAIN_SUSPECT
echo $MAIN_SUSPECT