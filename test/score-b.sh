result=$(go run main.go -o resources/output-files/b.out -i resources/input-files/b_should_be_easy.in -noGui)
echo "B) result " $result
if [ $result == 171231 ]; then
  echo "Good :)"
  exit 0
fi
echo "Bad :)"
exit 1
