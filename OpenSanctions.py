#!/usr/bin/python3
import json

def main():
    dict_entities_parsed = {}
    dict_targets_parsed = {}
    dict_topics_parsed = {}
    dict_entities = {}
    dict_targets = {}
    dict_topics = {}

    #print("Building Entities...")
    with open("entities.ftm.json", "r") as file:
        for line in file:
            entity = json.loads(line)
            dict_entities[entity["id"]] = entity
            
            #if dict_entities["target"]:
            #    if dict_entities["schema"] == "Person" or dict_entities["schema"] == "Family" or dict_entities["schema"] == "Company" or dict_entities["schema"] == "Organization":
            #        if "passport" in line:
            #            properties = {}
            #            properties = dict_entities["properties"]
            #            if "passportNumber" in line:
            #                for passport in properties["passportNumber"]:
            #                    dict_entities_passport[passport] = entity

    file.close()

    #print("Building Targets...")
    with open("targets.nested.json", "r") as file:
        for line in file:
            target = json.loads(line)
            dict_targets[target["id"]] = target

    file.close()

    #print("Building Topics...")
    with open("targets.nested.json", "r") as file:
        for line in file:
            topic = json.loads(line)
            dict_topics[topic["id"]] = topic
    
    file.close()

    #i = 0
    #for passport in dict_entities_passport:
    #    print("---->Entry #",i)
    #    for item in  dict_entities_passport[passport]:
    #        if item == "properties":
    #            for property in dict_entities_passport[passport][item]:
    #                for type in dict_entities_passport[passport][item][property]:
    #                    print(property + ":" + type)
    #        else:
    #            if item != "target":
    #                if item == "referents" or item == "datasets":
    #                    for list_item in dict_entities_passport[passport][item]:
    #                        print(item + ":" + list_item)
    #                else:
    #                    print(item + ":" + dict_entities_passport[passport][item])
    #                if item == "id":

    #    i = i + 1

    i = 0
    for ID in dict_entities:
        if ID in dict_targets and ID in dict_topics:
            for type in dict_entities[ID]:
                if type == "properties":
                    for property in dict_entities[ID][type]:
                        for list_item in dict_entities[ID][type][property]:
                            print(property + ": " + list_item)
                else:
                    if type != "target":
                        if type == "refetrents" or type == "datasets":
                            for list_item in dict_entities[ID][type]:
                                print(type + ": " + list_item)
                    else:
                        if dict_entities[ID][type]:
                            print("Target: True")
                        else:
                            print("Target: False")
                for type in dict_targets[ID]:
                    if type == "properties":
                        for property in dict_entities[ID][type]:
                            for list_item in dict_entities[ID][type][property]:
                                print(property + ": " + list_item)
                    else:
                        if type != "target":
                            if type == "refetrents" or type == "datasets":
                                for list_item in dict_entities[ID][type]:
                                    print(type + ": " + list_item)
                        else:
                            if dict_entities[ID][type]:
                                print("Target: True")
                            else:
                                print("Target: False")
                for type in dict_topics[ID]:
                    if type == "properties":
                        for property in dict_entities[ID][type]:
                            for list_item in dict_entities[ID][type][property]:
                                print(property + ": " + list_item)
                    else:
                        if type != "target":
                            if type == "refetrents" or type == "datasets":
                                for list_item in dict_entities[ID][type]:
                                    print(type + ": " + list_item)
                        else:
                            if dict_entities[ID][type]:
                                print("Target: True")
                            else:
                                print("Target: False")

        else:
            print("------------------------------------------------------------>NO!")
        i = i + 1
        print(i)



                        
main()

