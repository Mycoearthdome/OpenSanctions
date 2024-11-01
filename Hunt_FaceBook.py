
Persons_file = "Persons_Names_Alias.txt"
FaceBook_file = "FaceBook_2019.txt"
OutFile = "Persons_Found.txt"

print("Reading POI file...Please wait")
with open(Persons_file, "r") as f:
    Persons_Names = f.readlines()

f.close()

Persons_list = []
for entry in Persons_Names:
    entry = entry.split(":")
    if len(entry) > 1:
        Person_Name = entry[1].strip()
        Persons_list.append(Person_Name)


print("Loading FB_2019 data...Please wait!")
with open(FaceBook_file, "r", errors='ignore') as g:
    FaceBook2019_entries = g.readlines()

g.close()

print("Hunting...Please wait!")

Cleared = True
for Person in Persons_list:
    if not Cleared:
        with open(OutFile, "a") as out:
            for FB_data in Found:
                for entry in FB_data:
                    out.write(entry+":")
                out.write("\n")
        out.close()
    Found = []
    ResultsCount = 0
    Cleared = False
    for entry in FaceBook2019_entries:
        if Person in entry:
            print("Person found! --> %s [%s]" % (Person, entry))
            Found.append([Person,entry])
            ResultsCount += 1
            if ResultsCount == 36:
                Found = Found[:len(Found)-36]
                Cleared = True
                break
