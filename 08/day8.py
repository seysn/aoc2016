#!/usr/bin/env python3

class Display():
    def __init__(self, width=50, height=6):
        self.width = width
        self.height = height
        self.tab = [['.'] * width for i in range(height)]

    def __str__(self):
        res = ""
        for i in self.tab:
            for j in i:
                res += j
            res += '\n'
        return res[:-1]

    def set(self, x, y):
        "Turn on a pixel"
        self.tab[x][y] = '#'
        
    def add(self, width=1, height=1):
        "Add a rectangle at the top-left corner"
        for i in range(height):
            for j in range(width):
                self.set(i, j)

    def shift_h(self, y, pixels=1):
        "Horizontally shift a row"
        for i in range(pixels):
            self.tab[y].insert(0, self.tab[y][-i - 1])
        self.tab[y] = self.tab[y][:-pixels]

    def shift_v(self, x, pixels=1):
        "Vertically shift a column"
        tmp = []
        for i in range(self.height - pixels, self.height):
            tmp.append(self.tab[i][x])
        for i in range(self.height - pixels):
            tmp.append(self.tab[i][x])
        for i in range(self.height):
            self.tab[i][x] = tmp[i]

    def count(self):
        "Return the number of pixels"
        nb = 0
        for i in self.tab:
            for j in i:
                if j == '#': nb += 1
        return nb

# Tests
test = Display(5, 1)
assert str(test) == '.' * 5
test.add(2, 1)
assert str(test) == "##" + '.' * 3
test.shift_h(0, 2)
assert str(test) == "..##."
assert test.count() == 2
test.shift_h(0, 2)
assert str(test) == "#...#"

test = Display(1, 5)
assert str(test) == ".\n" * 4 + "."
test.add(1, 2)
assert str(test) == "#\n" * 2 + ".\n" * 2 + '.'
test.shift_v(0, 2)
assert str(test) == ".\n.\n#\n#\n."
assert test.count() == 2
test.shift_v(0, 2)
assert str(test) == "#\n.\n.\n.\n#"

# Main
if __name__ == "__main__":
    display = Display()
    with open("input.txt") as f:
        for line in f:
            t = line.split()
            if t[0] == "rect":
                t = t[1].split('x')
                display.add(int(t[0]), int(t[1]))
            elif t[1] == "row":
                display.shift_h(int(t[2][2:]), int(t[-1]))
            elif t[1] == "column":
                display.shift_v(int(t[2][2:]), int(t[-1]))
    print("Part One:", display.count())
    print("Part Two:", display, sep="\n")
