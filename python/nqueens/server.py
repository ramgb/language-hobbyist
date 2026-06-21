from flask import Flask
from markupsafe import escape
from nqueens import solve

app = Flask(__name__)

@app.route('/nqueens/<queens>')
def solve_n_queens(queens):
    solutions = solve(int(queens))
    return '{ solutions : '  + str(solutions) + '}'