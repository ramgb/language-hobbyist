import pickle
import time
from classes.db.dictionary import Dictionary
from classes.db.node import Node

PICKLE_FILE_NAME = 'alpha.pkl'

class PickleDictionary(Dictionary):
    
    def load_dictionary(self) -> Node:
        return pickle.load(open(PICKLE_FILE_NAME, 'rb'))
    
    def save_dictionary(self, dict_file_path: str) -> None:
        root = Node(None, False)
        t0 = time.perf_counter()
        with open(dict_file_path, "r") as reader:
            line = reader.readline()
            while line != '':
                newRoot = root
                word = line.rstrip()
                for char in word[:-1]:
                    curNode = newRoot.__add_child__(Node(char, False))
                    newRoot = curNode
                newRoot.__add_child__(Node(word[-1], True))
                line = reader.readline()
        pickle.dump(root, open(PICKLE_FILE_NAME, 'wb'), protocol=pickle.HIGHEST_PROTOCOL)
    
    def solve_helper(self, root, special_char, chars, cur_str, used_special_character, writer):
        # boundary condition hitting a leaf
        if not root.childset:
            if used_special_character and len(cur_str) >= Dictionary.MIN_SIZE:
                writer.write(cur_str + "\n")
            return
        # another condition
        if root.is_word_end and used_special_character and len(cur_str) >= Dictionary.MIN_SIZE:
            writer.write(cur_str + "\n")
        
        for child_value, child_node in root.childset.items():
            if child_value in chars:
                self.solve_helper(
                    child_node,
                    special_char,
                    chars,
                    cur_str + child_value,
                    used_special_character or child_value == special_char,
                    writer) 
            
    def solve(self, special_char: str, chars: str) -> None:
        root = self.load_dictionary()
        writer = open(Dictionary.solution_file_name(self, chars), "w")
        self.solve_helper(root, special_char, chars, '', False, writer)
        writer.close()
        