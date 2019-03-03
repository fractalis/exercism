class Matrix(object):
    def __init__(self, matrix_string):
        rows = matrix_string.split("\n")
        self.matrix = [r.split(" ") for r in rows]
        self.matrix = [list(map(int, x)) for x in self.matrix]


    def row(self, index):
        return self.matrix[index-1]

    def column(self, index):
        cols = [x[index-1] for x in self.matrix]
        return cols
