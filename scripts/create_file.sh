echo "Enter Problem Number"
read p_num
cd ..
mkdir -p "$p_num"
cp ./template.go "./$p_num/$p_num.go"