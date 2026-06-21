class Node:
    
    def __init__(self, value, is_word_end):
        self.value = value
        self.childset = {}
        self.parent = None
        self.valueset = {value} # Can this be deleted ? 
        self.is_word_end = is_word_end
        
    def __add_child__(self, child):
        if child.value not in self.childset:
            self.childset[child.value] = child
            child.parent = self
            child.__update_parent()
        return self.childset[child.value]
    
    def __update_parent(self):
        for value in self.valueset:
            self.parent.valueset.add(value)
    