#!/usr/bin/python3
import json
import pickle
import os

filename = "OpenSanction.pkl"

def Print_Persons(dict):
    for ID in dict:
        for type in dict[ID]:
            if type == "schema":
                if dict[ID]["schema"] == "Person":
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
        dict_entities = unpickle_dict(filename)
    else:
        with open("entities.ftm.json", "r") as file:
            for line in file:
                entity = json.loads(line)
                dict_entities[entity["id"]] = entity
        file.close()
        pickle_dict(dict_entities)
                    
    Print_Persons(dict_entities)

                        
main()

