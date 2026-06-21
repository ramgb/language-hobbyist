import datetime
from node import Node

MIN_SIZE = 4

class Dictionary:
    def load_dictionary(self) -> Node:
        pass
    
    def save_dictionary(self, dict_file_path: str) -> None:
        pass
    
    def solve(self, special_char: str, chars: str) -> None:
        pass
    
    def solution_file_name(self, chars:str) -> str:
        ct = datetime.datetime.now()
        return "solutions/" + str(ct) + "_" + chars +".txt"