#!/usr/bin/python3
import json
import pickle
import os

filename = "OpenSanction.pkl"

def Print_Selection(dict, choice):
    for ID in dict:
        for type in dict[ID]:
            if type == "properties":
                for property in dict[ID]["properties"]:
                    if property == "topics":
                        if choice in dict[ID]["properties"]["topics"]:
                            for kind in dict[ID]:
                                if kind == "properties":
                                    for property in dict[ID][kind]:
                                        for list_item in dict[ID][kind][property]:
                                            print(property + ": " + list_item)
                                else:
                                    if kind != "target":
                                        if kind == "referents" or kind == "datasets":
                                            for list_item in dict[ID][kind]:
                                                print(kind + ": " + list_item)
                                        else:
                                            print(kind +":"+ dict[ID][kind])
                                    else:
                                        if dict[ID][kind]:
                                            print("Target: True")
                                        else:
                                            print("Target: False")


def pickle_dict(dictionary):
    try:
        with open(filename, 'wb') as f:
            pickle.dump(dictionary, f)
        print("Dictionary pickled successfully!")
    except Exception as e:
        print(f"Error pickling dictionary: {e}")

def unpickle_dict(file_path):
    try:
        with open(file_path, 'rb') as f:
            dictionary = pickle.load(f)
        print("Dictionary unpickled successfully!")
        return dictionary
    except Exception as e:
        print(f"Error unpickling dictionary: {e}")

def main():
    
    dict_entities = {}
  
    if os.path.exists(filename):
        print("Loading from dictionary...Please wait!")
        dict_entities = unpickle_dict(filename)
    else:
        print("Building Entities...")
        with open("entities.ftm.json", "r") as file:
            for line in file:
                entity = json.loads(line)
                dict_entities[entity["id"]] = entity
        file.close()
        pickle_dict(dict_entities)

    dict_stats = {}
    #i = 0
    for ID in dict_entities:
        for type in dict_entities[ID]:
            if type == "properties":
                for property in dict_entities[ID][type]:
                    for list_item in dict_entities[ID][type][property]:
                        if property == "topics":
                            if not list_item in dict_stats:
                                dict_stats[list_item] = 1
                            else:
                                dict_stats[list_item] += 1
                        #print(property + ": " + list_item)
            else:
                if type == "target":
                    if dict_entities[ID][type]:
                        if not type in dict_stats:
                            dict_stats[type] = 0
                        dict_stats[type] += 1
                        #print("Target: True")
                    else:
                        if not "Not_Targetted" in dict_stats:
                            dict_stats["Not_Targetted"] = 0
                        dict_stats["Not_Targetted"] += 1
                        #print("Target: False")
        #i = i + 1
        #print(i)

    while(1):
        print("---=== Here are your choices ===---\n")
        for topic in dict_stats:
            print(topic+":"+ str(dict_stats[topic]))

        choice = input("Choose from the list above:")

        Print_Selection(dict_entities, choice)

                        
main()

