#!/usr/bin/env python3

FILENAME = "input.txt"
DIRECTIONS = ['N', 'E', 'S', 'W']
D_VALUES = {
    'N': ( 0,  1),
    'E': ( 1,  0),
    'S': ( 0, -1),
    'W': (-1,  0)
}

rules = []

def init_rules():
    with open(FILENAME) as f:
        for r in f.read().replace(' ', '').split(','):
            rules.append((r[0], int(r[1:])))

def execute():
    current_d = 'N'
    current_x, current_y = 0, 0
    for r in rules:
        print("Current direction : {} - Position (x={}, y={}) --> Next move : {}".format(current_d, current_x, current_y, r))
        current_d = DIRECTIONS[(DIRECTIONS.index(current_d) + (1 if r[0] == 'R' else -1)) % 4]
        current_x += D_VALUES[current_d][0] * r[1]
        current_y += D_VALUES[current_d][1] * r[1]
    print("Current direction : {} - Position (x={}, y={}) --> Next move : {}".format(current_d, current_x, current_y, r))
    print("You are {} blocks away".format(abs(current_x) + abs(current_y)))
            
def main():
    init_rules()
    execute()

if __name__ == "__main__":
    main()
