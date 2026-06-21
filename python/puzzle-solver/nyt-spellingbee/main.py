import argparse
from classes.spellingbee import solver

if __name__=="__main__":
    parser = argparse.ArgumentParser("solver")
    parser.add_argument("chars", help="All the characters in the current spelling bee", type=str)

    chars = parser.parse_args().chars

    if not chars:
        raise Exception('Input string chars should be non-empty')

    root = solver.solve(chars)