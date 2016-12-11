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
    positions = []
    current_d = 'N'
    current_x, current_y = 0, 0
    for r in rules:
        current_d = DIRECTIONS[(DIRECTIONS.index(current_d) + (1 if r[0] == 'R' else -1)) % 4]
        for i in range(int(r[1])):
            curr_tmp = (current_x + D_VALUES[current_d][0] * i, current_y + D_VALUES[current_d][1] * i)
            if curr_tmp in positions:
                print("You already visited this location :")
                print("{} - You are {} blocks away".format(curr_tmp, abs(curr_tmp[0]) + abs(curr_tmp[1])))
            else:
                positions.append(curr_tmp)
        current_x += D_VALUES[current_d][0] * r[1]
        current_y += D_VALUES[current_d][1] * r[1]
            
def main():
    init_rules()
    execute()

if __name__ == "__main__":
    main()
