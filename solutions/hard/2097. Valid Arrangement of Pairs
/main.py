from collections import defaultdict, deque

class Solution(object):
    def validArrangement(self, pairs):
        """
        :type pairs: List[List[int]]
        :rtype: List[List[int]]
        """
        graph = defaultdict(deque)
        ins = defaultdict(int)
        outs = defaultdict(int)
        
        # Build graph and track in-degrees & out-degrees
        for s, e in pairs:
            graph[s].append(e)
            outs[s] += 1
            ins[e] += 1
        
        # Find the starting node for Eulerian path
        s_node = None
        for node in graph:
            if outs[node] > ins[node]:  
                s_node = node
                break
        if s_node is None:
            s_node = pairs[0][0]  # If all nodes are balanced, start anywhere
        
        # Hierholzer's Algorithm to find the Eulerian path
        stack = [s_node]
        path = []
        
        while stack:
            node = stack[-1]
            if graph[node]:  # If there are outgoing edges
                next_node = graph[node].popleft()
                stack.append(next_node)
            else:  # If no outgoing edges, backtrack
                path.append(stack.pop())
        
        # Reverse path and construct result
        path.reverse()
        return [[path[i], path[i+1]] for i in range(len(path) - 1)]

