import pickle
import time
from classes.db.node import Node

DICT_1 = '/usr/share/dict/words'
DICT_2 = 'dicts/words_alpha.txt'
PICKLE_DICT_1='dicts/words.pkl'
PICKLE_DICT_2='dicts/alpha.pkl'

# TODO(ramgb): change this to an interface for db lookups

def save_dictionary(dictionary_file, pickle_file_name):
    root = Node(None, False)
    t0 = time.perf_counter()
    with open(dictionary_file, "r") as reader:
        line = reader.readline()
        while line != '':
            newRoot = root
            word = line.rstrip()
            for char in word[:-1]:
                curNode = newRoot.__add_child__(Node(char, False))
                newRoot = curNode
            newRoot.__add_child__(Node(word[-1], True))
            line = reader.readline()
    pickle.dump(root, open(pickle_file_name, 'wb'), protocol=pickle.HIGHEST_PROTOCOL)
    
def load_dictionary_from_file(pickle_file_name):
    return pickle.load(open(pickle_file_name, 'rb'))

def load_dictionary():
    return load_dictionary_from_file(PICKLE_DICT_2)

def write_dictionaries():
    save_dictionary(DICT_1, PICKLE_DICT_1)
    save_dictionary(DICT_2, PICKLE_DICT_2)