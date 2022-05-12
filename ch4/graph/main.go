package main

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	if graph[from] == nil {
		graph[from] = make(map[string]bool)
	}
	graph[from][to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
