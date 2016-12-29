from collections import Counter

test = '''
eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar
'''.strip()

def combine_chars(text, most_common = True):
    res = ""
    for line in zip(*(text.split())):
        c = Counter(line)
        res += c.most_common()[0 if most_common else -1][0]
    return res

assert combine_chars(test) == "easter"
assert combine_chars(test, False) == "advent"

with open("input.txt") as f:
    text = f.read()
    print("Part One:", combine_chars(text))
    print("Part Two:", combine_chars(text, False))
