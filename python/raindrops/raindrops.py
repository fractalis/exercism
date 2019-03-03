def raindrops(number):
    numStr = ""
    if number%3 == 0:
        numStr += "Pling"
    if number%5 == 0:
        numStr += "Plang"
    if number%7 == 0:
        numStr += "Plong"
    if numStr == "":
        numStr = str(number)
    return numStr

