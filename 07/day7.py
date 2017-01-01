#!/usr/bin/env python3
import re
import string

def is_abba(string):
    if len(string) == 4:
        return string[0] != string[1] and string[0:2] == string[3:1:-1]
    else:
        for i in range(len(string) - 3):
            if is_abba(string[i:i+4]):
                return True
        return False

assert is_abba("abba")
assert not is_abba("babb")
assert is_abba("ioxxoj")
assert not is_abba("zxcvbn")
assert is_abba("xxyyx")

def support_tls(line):
    tab = re.split(r'\[|\]', line)
    return any([is_abba(outside) for outside in tab[0::2]]) and all(not is_abba(inside) for inside in tab[1::2])

assert support_tls("abba[mnop]qrst")
assert not support_tls("abcd[bddb]xyyx")
assert not support_tls("aaaa[qwer]tyui")
assert support_tls("ioxxoj[asdfgh]zxcvbn")
assert not support_tls("abba[xxyyx]azee")


def support_ssl(line):
    tab = re.split(r'\[|\]', line)
    return any(a+b+a in "-".join(tab[0::2]) and b+a+b in "-".join(tab[1::2]) for a in string.ascii_lowercase for b in string.ascii_lowercase if a != b)

assert support_ssl("aba[bab]xyz")
assert not support_ssl("xyx[xyx]xyx")
assert support_ssl("aaa[kek]eke")
assert support_ssl("zazbz[bzb]cdb")

with open("input.txt") as f:
    content = f.read().split('\n')
    print("Part One:", sum(support_tls(line) for line in content))
    print("Part Two:", sum(support_ssl(line) for line in content))
