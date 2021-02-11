
RED="\033[0;31m"
GREEN="\033[0;32m"
NC='\033[0m' # No Color

test_code() {
    runCommand=$1
    test_number=$2
    if ! $runCommand example/Code_$test_number.data | diff -wB example/Instruction_$test_number.data -; then
        echo -e "${RED}$test_number - Failed!${NC}"
    else
        echo -e "${GREEN}$test_number - Passed!${NC}"
    fi
}

echo "Testing Go Implementation"
test_code "go run assembler.go" 1
test_code "go run assembler.go" 2
echo ""

echo "Testing Python Implementation"
test_code "python3 assembler.py" 1
test_code "python3 assembler.py" 2
