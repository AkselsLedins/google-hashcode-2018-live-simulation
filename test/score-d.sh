result=$(go run main.go -o resources/output-files/d.out -i resources/input-files/d_metropolis.in -noGui)
echo "D) result " $result
if [ $result == 3360230 ]; then
  echo "Good :)"
  exit 0
fi
echo "Bad :)"
exit 1

