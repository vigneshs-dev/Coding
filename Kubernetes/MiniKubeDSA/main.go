package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pod struct {
	Name   string
	CPU    int
	Memory int
}

type Node struct {
	Name       string
	CPUFree    int
	MemoryFree int
	Pods       []string
}

func (n *Node) CanFit(p Pod) bool {
	return n.CPUFree >= p.CPU && n.MemoryFree >= p.Memory
}

func (n *Node) AddPod(p Pod) {
	n.CPUFree -= p.CPU
	n.MemoryFree -= p.Memory
	n.Pods = append(n.Pods, p.Name)
}

var nodes []*Node
var pods []Pod

func schedulePods() {
	sort.Slice(pods, func(i, j int) bool {
		return (pods[i].CPU + pods[i].Memory) > (pods[j].CPU + pods[j].Memory)
	})

	for i := 0; i < len(pods); {
		pod := pods[i]
		scheduled := false
		for _, node := range nodes {
			if node.CanFit(pod) {
				node.AddPod(pod)
				fmt.Printf("✅ Pod %s scheduled on Node %s\n", pod.Name, node.Name)
				// Remove pod from unscheduled list
				pods = append(pods[:i], pods[i+1:]...)
				scheduled = true
				break
			}
		}
		if !scheduled {
			i++ // Try next pod if this one can’t be scheduled yet
		}
	}
}

func printStatus() {
	fmt.Println("\n📦 Cluster Status:")
	for _, node := range nodes {
		fmt.Printf("Node %s [Free CPU: %d, Free MEM: %d] → Pods: %v\n",
			node.Name, node.CPUFree, node.MemoryFree, node.Pods)
	}
	if len(pods) > 0 {
		fmt.Println("\n🕗 Unscheduled Pods:")
		for _, p := range pods {
			fmt.Printf("- %s (CPU: %d, MEM: %d)\n", p.Name, p.CPU, p.Memory)
		}
	}
	fmt.Println()
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("🎮 MiniKubeDSA CLI — type 'help' for commands\n")

	for {
		fmt.Print("> ")
		if !reader.Scan() {
			break
		}
		line := strings.TrimSpace(reader.Text())
		args := strings.Split(line, " ")

		if len(args) == 0 || args[0] == "" {
			continue
		}

		switch args[0] {
		case "help":
			fmt.Println(`
Commands:
  add node <name> <cpu> <memory>    → Add a new node
  add pod <name> <cpu> <memory>     → Add a new pod
  schedule                          → Run scheduler
  status                            → Show current cluster state
  exit                              → Quit
`)
		case "add":
			if len(args) != 5 {
				fmt.Println("❌ Usage: add node|pod <name> <cpu> <memory>")
				continue
			}
			name := args[2]
			cpu, _ := strconv.Atoi(args[3])
			mem, _ := strconv.Atoi(args[4])

			if args[1] == "node" {
				nodes = append(nodes, &Node{name, cpu, mem, []string{}})
				fmt.Printf("✅ Added Node %s (CPU: %d, MEM: %d)\n", name, cpu, mem)
			} else if args[1] == "pod" {
				pods = append(pods, Pod{name, cpu, mem})
				fmt.Printf("✅ Added Pod %s (CPU: %d, MEM: %d)\n", name, cpu, mem)
			} else {
				fmt.Println("❌ Unknown type:", args[1])
			}

		case "schedule":
			schedulePods()

		case "status":
			printStatus()

		case "exit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Unknown command. Type 'help' for a list.")
		}
	}
}
