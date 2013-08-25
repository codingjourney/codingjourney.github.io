import sys
import time

def distance(left, right):
    result = max(len(left), len(right))
    rightMap = mapLetters(right)
    for x in range(len(left)):
      letter = left[x]
      if letter in rightMap:
        for y in rightMap[letter]:
          beforeMatch = max(x, y)
          if beforeMatch < result:
            xOut = x
            yOut = y
            while (xOut < len(left) and yOut < len(right) and
                   left[xOut] == right[yOut]):
              xOut += 1
              yOut += 1
            afterMatch = distance(left[xOut:], right[yOut:])
            result = min(result, beforeMatch + afterMatch)
    return result

def mapLetters(text):
    result = {}
    for offset in range(len(text)):
        letter = text[offset]
        if letter not in result:
            result[letter] = []
        result[letter].append(offset)
    return result

# command-line interface

if len(sys.argv) < 3:
    print("USAGE: python levenshtein.py /path/to/dict numberOfWords")
    sys.exit(1)

f = open("/usr/share/dict/american-english")
words = []
for word in f: words.append(word[:-1])

start = time.clock()
limit = int(sys.argv[2])
for i in range(0, limit - 1):
    for j in range(i + 1, limit):
        dist = distance(words[i], words[j])
        print("%s vs. %s: %d" % (words[i], words[j], distance(words[i], words[j])))

print("%d sec for %d pairs" % ((time.clock() - start), (limit * (limit - 1) / 2)))
