
Entities_file = "entities.ftm.json"
Persons_file = "Reconciled_Persons.txt"
OutFile = "Probable_FB_Sanctions.txt"

print("Reading FOUND POI file...Please wait")
with open(Persons_file, "r") as f:
    Persons_Names = f.readlines()

f.close()

Persons_list = []
for entry in Persons_Names:
    if entry != "]\n":
        entry = entry.split(" ")
        Name = entry[0] +" " + entry[1]
        Persons_list.append([Name, ",".join(entry[2:])])

print("Loading Entities data...Please wait!")
with open(Entities_file, "r", errors='ignore') as g:
    Entities_entries = g.readlines()
g.close()

print("Concluding Hunt...Please wait!")

Cleared = True
for Infos in Persons_list:
    for entry in Entities_entries:
        if Infos[0] in entry: #Name
            print("%s\n[%s]\nFacebook->%s" % (Infos[0], entry, Infos[1]))
            with open(OutFile, "a") as f:
                f.write(Infos[0]+"\n"+entry+"\nFacebook->"+Infos[1]+"\n") #Add FACEBOOK to the possible Sanctions.
            f.close()
            break
