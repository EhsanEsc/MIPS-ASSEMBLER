# MIPS-ASSEMBLER

This Assembler converts MIPS instructions to binary format. It can be used in computer architucture projects.
Script used instruction signatures from https://opencores.org/projects/plasma/opcodes. 
You can add new instruction signature to MIPSInstructions.txt according to format.

Script takes one argument for address of instructions file. Then prints binary code of each instruction

It is implemented by python and go. Code is very very dirty becuase of tight deadline.
It was originally my arbitrary project for computer workshop lesson.

# Running
Go: go run assembler.go PATH_TO_INSTRUCTION_FILE
Python: python3 assembler.py PATH_TO_INSTRUCTION_FILE

# Testing
Run tester.sh script to test all codes with examples.

# TODO
1. Clean Python Code
2. Add implementation of Haskal.
3. Add more error handling
