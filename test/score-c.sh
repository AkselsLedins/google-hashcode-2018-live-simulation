result=$(go run main.go -o resources/output-files/c.out -i resources/input-files/c_no_hurry.in -noGui)
echo "C) result " $result
if [ $result == 7377192 ]; then
  echo "Good :)"
  exit 0
fi
echo "Bad :)"
exit 1

