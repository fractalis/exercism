def abbreviate(words):
    words = words.replace("_", "").replace(" - ", " ").replace("-", " ")
    abbrev = "".join(list(map(lambda x: x[0].upper(), words.split(" "))))
    return abbrev
