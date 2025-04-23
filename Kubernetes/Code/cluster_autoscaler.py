class Pod:
    def __init__(self, name, cpu, memory):
        self.name = name
        self.cpu = cpu  # millicores
        self.memory = memory  # MiB

class Node:
    def __init__(self, cpu_capacity, memory_capacity):
        self.cpu_free = cpu_capacity
        self.memory_free = memory_capacity
        self.pods = []

    def can_fit(self, pod):
        return self.cpu_free >= pod.cpu and self.memory_free >= pod.memory

    def add_pod(self, pod):
        self.cpu_free -= pod.cpu
        self.memory_free -= pod.memory
        self.pods.append(pod.name)

# Greedy Fit Algorithm: First Fit Decreasing
def autoscale_cluster(pods, node_cpu, node_memory):
    pods = sorted(pods, key=lambda p: (p.cpu + p.memory), reverse=True)
    nodes = []

    for pod in pods:
        placed = False
        for node in nodes:
            if node.can_fit(pod):
                node.add_pod(pod)
                placed = True
                break
        if not placed:
            new_node = Node(node_cpu, node_memory)
            new_node.add_pod(pod)
            nodes.append(new_node)

    return nodes

# Example Pods and node type
pending_pods = [
    Pod("pod-a", 400, 512),
    Pod("pod-b", 600, 1024),
    Pod("pod-c", 300, 512),
    Pod("pod-d", 800, 2048),
    Pod("pod-e", 500, 1000),
]

node_template_cpu = 2000  # 2 cores
node_template_memory = 4096  # 4Gi

nodes_created = autoscale_cluster(pending_pods, node_template_cpu, node_template_memory)

# Output result
print(f"Number of new nodes needed: {len(nodes_created)}")
for i, node in enumerate(nodes_created):
    print(f"Node-{i+1} runs pods: {node.pods}")
