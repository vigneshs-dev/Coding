from collections import defaultdict, deque

# Define dependencies between steps (containers)
# Format: "A": ["B"] means A depends on B (B must finish before A runs)
dependencies = {
    "app": ["db", "cache"],
    "db": [],
    "cache": ["redis"],
    "redis": []
}

# Topological Sort using Kahn's Algorithm
def topological_sort(dependency_graph):
    in_degree = defaultdict(int)
    graph = defaultdict(list)

    # Build graph and count in-degrees
    for node, deps in dependency_graph.items():
        for dep in deps:
            graph[dep].append(node)
            in_degree[node] += 1

    # Start with nodes having no dependencies
    queue = deque([node for node in dependency_graph if in_degree[node] == 0])
    execution_order = []

    while queue:
        current = queue.popleft()
        execution_order.append(current)

        for neighbor in graph[current]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)

    if len(execution_order) != len(dependency_graph):
        return "Cycle detected (invalid DAG)"
    
    return execution_order

# Run it
print("Container Start Order:")
print(topological_sort(dependencies))
