from collections import defaultdict
import re

def word_count(phrase):
    word_dict = defaultdict(lambda: 0)
    words = re.findall("[0-9]|[A-Za-z]+'*[A-Za-z]", phrase.replace("_", " "))
    for word in words:
        word_dict[word.lower()] += 1
    return word_dict
