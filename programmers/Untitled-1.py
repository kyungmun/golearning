
def solution(arrows):
    answer = 0
    nodes = [(0,0)]
    edges = []
    switchdict = {0: vertex_0, 1: vertex_1, 2: vertex_2, 3: vertex_3, 4: vertex_4, 5: vertex_5, 6: vertex_6, 7: vertex_7}
    set(edges)
    set(nodes)
    cur_node = [0,0]
    for arrow in arrows:
        node , edge = switchdict[arrow](cur_node)
        cross_edge = (0,0)
        if arrow == 1:
            cross_edge = ((cur_node[0] , cur_node[1] + 1 ) , ( node[0] , node[1] - 1 ),)
        elif arrow == 3:
            cross_edge = ((cur_node[0] , cur_node[1] - 1 ) , ( node[0] , node[1] + 1 ),)
        elif arrow == 5:
            cross_edge = ((node[0] , node[1] + 1 ) , ( cur_node[0] , cur_node[1] - 1 ),)
        elif arrow == 7:
            cross_edge = ((node[0] , node[1] - 1 ) , ( cur_node[0] , cur_node[1] + 1 ),)

        cur_node[0] = node[0]
        cur_node[1] = node[1]

        if edge not in edges:
            if cross_edge in edges:
                answer += 1
            if node in nodes :
                answer += 1
            else:
                nodes.append(node)
            edges.append(edge)

        print(node, edge, cross_edge, answer)
    return answer

def vertex_0(cur_node):
    return (cur_node[0], cur_node[1]+1), ((cur_node[0], cur_node[1]),(cur_node[0], cur_node[1]+1))

def vertex_1(cur_node):
    return (cur_node[0]+1, cur_node[1]+1), ((cur_node[0], cur_node[1]),(cur_node[0]+1, cur_node[1]+1))

def vertex_2(cur_node):
    return (cur_node[0]+1, cur_node[1]), ((cur_node[0], cur_node[1]),(cur_node[0]+1, cur_node[1]))

def vertex_3(cur_node):
    return (cur_node[0]+1, cur_node[1]-1), ((cur_node[0], cur_node[1]),(cur_node[0]+1, cur_node[1]-1))

def vertex_4(cur_node):
    return (cur_node[0], cur_node[1]-1), ((cur_node[0], cur_node[1]-1),(cur_node[0], cur_node[1]))

def vertex_5(cur_node):
    return (cur_node[0]-1, cur_node[1]-1), ((cur_node[0]-1, cur_node[1]-1),(cur_node[0], cur_node[1]))

def vertex_6(cur_node):
    return (cur_node[0]-1, cur_node[1]), ((cur_node[0]-1, cur_node[1]),(cur_node[0], cur_node[1]))

def vertex_7(cur_node):
    return (cur_node[0]-1, cur_node[1]+1), ((cur_node[0]-1, cur_node[1]+1),(cur_node[0], cur_node[1]))


solution([5, 2, 2, 7, 7, 2, 2, 5, 1, 5])    