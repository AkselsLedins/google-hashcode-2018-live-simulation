result=$(go run main.go -o resources/output-files/a.out -i resources/input-files/a_example.in -noGui)
echo "A) result " $result
if [ $result == 10 ]; then
  echo "Good :)"
  exit 0
fi
echo "Bad :)"
exit 1
