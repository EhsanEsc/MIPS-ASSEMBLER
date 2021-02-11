import re
import sys
import os

insts = open("MIPSInstructions.txt", "r")
commands = {}

for inst in insts.readlines():
    if(inst == ""):
        continue
    tmp = inst.split(" ")
    if(len(tmp) != 2):
        err(inst)
        continue
    command = tmp[0]
    ml_command = tmp[1]
    ml_command = re.sub(r"\n", '', ml_command)
    command = command.replace('(',",")
    command = command.replace(")","")
    tmp = command.split(',')
    tmp2 = ml_command.split(',')
    tmp.append(tmp2)
    commands[tmp[0]] = tmp

if(len(sys.argv) == 1):
    path = input("Enter Path: ")
    inp = open(path, "r")
else:
    inp = open(sys.argv[1], "r")

if not os.path.isfile("mips_inst.data"):
    oo = open("mips_inst.data", "x")
else:
    oo = open("mips_inst.data", "w")

for inst in inp.readlines():
    inst = inst.replace("\n","")
    tinst = inst
    inst.replace('(',",")
    inst = re.sub(r"\s+", ',', inst)
    inst = inst.replace('(',",")
    inst = inst.replace(')',"")
    if(inst[-1] == ")"):
        inst.pop()
    tmp = inst.split(',')
    ans = ""
    try:
        for x in commands[tmp[0].lower()][-1]:
            if(x[0] == "R"):
                counter = 0
                for y in commands[tmp[0].lower()][:-1]:
                    if(x == y):
                        t = tmp[counter][1:]
                        t = int(t)
                        res = "{0:b}".format(t)
                        if(len(res) < 5):
                            res = (5 - len(res)) * ("0") + res
                        ans += res
                        break
                    counter = counter + 1
            elif(x[0] == "N"):
                counter = 0
                for y in commands[tmp[0].lower()][:-1]:
                    if(x[0:2] == y):
                        res = "{0:b}".format(int(tmp[counter]))
                        if(len(res) < int(x[3:])):
                            res = (int(x[3:]) - len(res)) * ("0") + res
                        ans += res
                        break
                    counter = counter + 1
            else:
                ans += x
    except:
        print("Error in assembeling this command : ", end="")
        ans = tinst

    print(tinst, "->", ans)
    oo.write(ans)
    oo.write("\n")
oo.close()

print("Completed Successfully!")