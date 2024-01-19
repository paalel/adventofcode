import numpy as np
import networkx as nx
import matplotlib.pyplot as plt


def get_graph(line: str) -> bool:
    assert len(line) == 64
    matrix = np.array([int(x) for x in line]).reshape((8, 8))

    # create adjacency matrix from matrix
    adjacency_matrix = np.zeros((64, 64))
    for i in range(8):
        for j in range(8):
            if i > 0 and matrix[i - 1][j] == 1:
                adjacency_matrix[i * 8 + j][(i - 1) * 8 + j] = 1
            if i < 7 and matrix[i + 1][j] == 1:
                adjacency_matrix[i * 8 + j][(i + 1) * 8 + j] = 1
            if j > 0 and matrix[i][j - 1] == 1:
                adjacency_matrix[i * 8 + j][i * 8 + j - 1] = 1
            if j < 7 and matrix[i][j + 1] == 1:
                adjacency_matrix[i * 8 + j][i * 8 + j + 1] = 1

    graph = nx.from_numpy_array(adjacency_matrix)

    print(matrix)
    nx.draw(graph, with_labels=True)
    plt.show()
    return nx.is_connected(graph)


def check_graph(line: str) -> bool:
    graph = get_graph(line)
    for edge in graph.edges:
        new_graph = graph.copy()
        new_graph.remove_edge(*edge)


def solve(filename: str) -> int:
    count = 0
    with open(filename, "r") as f:
        for line in f:
            count += check_graph(line)

    return count


t0 = "1110101011111110110010001011100011110000100110000001100000000000"
t1 = "0001110010000100100001001000010011111100100001001000010000000100"

assert check_graph(t0)
assert check_graph(t1)

print(solve("sjokkis.txt"))
