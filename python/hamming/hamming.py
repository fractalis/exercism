def distance(strand_a, strand_b):

    diff = 0

    if len(strand_a) != len(strand_b):
        raise ValueError(r".+")

    for i, x in enumerate(strand_a):
        if strand_a[i] != strand_b[i]:
            diff += 1
    return diff
