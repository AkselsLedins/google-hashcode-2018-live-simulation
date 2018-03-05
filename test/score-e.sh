result=$(go run main.go -o resources/output-files/e.out -i resources/input-files/e_high_bonus.in -noGui)
echo "E) result " $result
if [ $result == 16381105 ]; then
  echo "Good :)"
  exit 0
fi
echo "Bad :)"
exit 1
