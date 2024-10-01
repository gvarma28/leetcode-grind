echo "Enter Problem Number"
read p_num
cd ..
mkdir -p "p_$p_num"
cp ./template.go "./p_$p_num/p_$p_num.go"