
Persons_file = "Persons_Names_Alias.txt"
Persons_found = "Persons_Found.txt"
OutFile = "Reconciled_Persons.txt"

with open(Persons_file, "r") as f:
    Persons_hunted = f.readlines()

Persons_Found = []
with open(Persons_found, "r", errors='ignore') as f:
    for line in f:
        if line != ":\n":
            if len(line) < 300:
                Persons_Found.append(line)

PersonNameCheck = []
for Persons in Persons_hunted:
    if "firstName: " in Persons:
        PersonNameCheck.append(Persons.split("firstName: ")[1].strip())
    if "lastName: " in Persons:
        PersonNameCheck.append(Persons.split("lastName: ")[1].strip())
    if "firstName:" in Persons:
        PersonNameCheck.append(Persons.split("firstName:")[1].strip())
    if "lastName:" in Persons:
        PersonNameCheck.append(Persons.split("lastName:")[1].strip())

print("Reconcilliation...")
for FBdetails in Persons_Found:
    if FBdetails != ":\n":
        details = FBdetails.split(":")
        if len(details) > 7:
            i = 0
            for detail in details:
                if "male" == detail or "female" == detail:
                    FirstName = details[i-2]
                    LastName = details[i-1]
                    break
                i = i + 1

            if FirstName != "None":
                #print(FirstName+","+LastName)
                if FirstName in PersonNameCheck and LastName in PersonNameCheck:
                    print("Probable match --> %s %s [%s]" % (FirstName, LastName, FBdetails))
                    with open(OutFile, "a") as f:
                        f.write(FirstName +" "+ LastName +" [" + FBdetails + "]\n")
                    f.close()

        else:
            if "female" in details[1]:
                details = details[1].split(',"female",')
                details = details[0].split(',female,')
            else:
                if "male" in details[1]:
                    details = details[1].split(',"male",')
                    details = details[0].split(',male,')
                else:
                    details = [details[1]]
            try:
                FirstName = details[0].split(",")[-2].strip('"')
                LastName = details[0].split(",")[-1].strip('"')
            except:
                continue
            if FirstName != "None":
                #print(FirstName+","+LastName)
                if FirstName in PersonNameCheck and LastName in PersonNameCheck:
                    print("Probable match --> %s %s [%s]" % (FirstName, LastName, FBdetails))
                    with open(OutFile, "a") as f:
                        f.write(FirstName +" "+ LastName +" [" + FBdetails + "]\n")
                    f.close()