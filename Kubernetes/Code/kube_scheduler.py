class Node:
    def __init__(self, name, cpu_free, memory_free):
        self.name = name
        self.cpu_free = cpu_free  # in millicores
        self.memory_free = memory_free  # in MiB

def can_schedule(node, pod_cpu, pod_memory):
    return node.cpu_free >= pod_cpu and node.memory_free >= pod_memory

def score_node(node):
    # Greedy heuristic: prefer nodes with more free resources
    return node.cpu_free + node.memory_free  # simple additive score

def schedule_pod(nodes, pod_cpu, pod_memory):
    # Step 1: Filter nodes
    eligible_nodes = [node for node in nodes if can_schedule(node, pod_cpu, pod_memory)]
    
    if not eligible_nodes:
        return "No eligible node found"

    # Step 2: Score nodes
    scored_nodes = [(node, score_node(node)) for node in eligible_nodes]

    # Step 3: Select best node (Greedy pick)
    best_node = max(scored_nodes, key=lambda x: x[1])[0]
    
    return f"Pod scheduled on: {best_node.name}"

# Example usage
nodes = [
    Node("node-a", 1000, 2000),
    Node("node-b", 500, 1000),
    Node("node-c", 800, 3000)
]

pod_cpu = 400
pod_memory = 500

print(schedule_pod(nodes, pod_cpu, pod_memory))
