import argparse
from nqueens import solve

def print_queens(arr_queens):
    for row in range(len(arr_queens)):
        for col in range(len(arr_queens)):
            if arr_queens[row][col] > 0:
                print('Q ', end='')
            else:
                print('- ', end='')
        print('')
        
if __name__=="__main__":
    parser = argparse.ArgumentParser("solver")
    parser.add_argument("queens", help="number of queens a.k.a the dimensions of the board", type=int)

    num_queens = parser.parse_args().queens

    if not num_queens:
        raise Exception('Input value for queens should not be empty')
    
    for solution in solve(num_queens):
        print_queens(solution)
        print('\n\n')