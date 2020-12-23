package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/goccy/go-graphviz"
)

type Trace struct {
	originServiceName string
	toHost            string
	httMethod         string
	URI               string
}

type Edge struct {
	source      string
	destination string
	attribute   string
}
type Node struct {
	nodeName    string
	clusterName string
	endPoints   []string
}

type DependencyGraph struct {
	nodes []Node
	edges []Edge
}

func renderGraph(s string) error {
	graph, err := graphviz.ParseBytes([]byte(s))
	if err != nil {
		return err
	}

	graphviz := graphviz.New()
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		graphviz.Close()
	}()
	graphviz.RenderFilename(graph, "png", "rw.png")
	return nil
}

func createGraph(dependencyGraph DependencyGraph) string {
	graphTemplate := `
		digraph G { 
			node [color="#40b9e5" fontcolor="#40b9e5"] 
			edge [fontsize=11 color=gray50 fontcolor=gray50]  			
			compound=true
			%s
			%s
			%s
		}	   
	`
	subGraphTemplate := `
		subgraph %s {
			label="%s"
			%s
		}
	`
	edgeTemplate := "\"%s\" -> \"%s\" [label= %s %s]\n\t\t"

	var subGraphs strings.Builder
	var nodes strings.Builder
	var edges strings.Builder

	for _, node := range dependencyGraph.nodes {

		var endpoints strings.Builder
		for _, endpoint := range node.endPoints {
			endpoints.WriteString(fmt.Sprintf("\"%s\"\n\t\t", endpoint))
			nodes.WriteString(fmt.Sprintf("\"%s\"\n\t\t", endpoint))
		}
		subGraph := fmt.Sprintf(subGraphTemplate, node.clusterName, node.nodeName, endpoints.String())
		subGraphs.WriteString(subGraph)
	}

	for _, edge := range dependencyGraph.edges {
		clusterName := findClusterByNodeName(dependencyGraph.nodes, edge.source)
		edgeSourceEndpoint := getFirstEndpointsByNodeName(dependencyGraph.nodes, edge.source)
		edgeGraph := fmt.Sprintf(edgeTemplate, edgeSourceEndpoint, edge.destination, edge.attribute, fmt.Sprintf("ltail=%s", clusterName))

		edges.WriteString(edgeGraph)

	}

	return fmt.Sprintf(graphTemplate, subGraphs.String(), nodes.String(), edges.String())
}

func findClusterByNodeName(nodes []Node, nodeName string) string {
	for _, node := range nodes {
		if node.nodeName == nodeName {
			return node.clusterName
		}
	}
	return ""
}

func getFirstEndpointsByNodeName(nodes []Node, nodeName string) string {
	for _, node := range nodes {
		if node.nodeName == nodeName {
			if len(node.endPoints) > 0 {
				return node.endPoints[0]
			}
		}
	}
	return ""
}

func AppendIfMissing(slice []string, str string) []string {
	for _, ele := range slice {
		if ele == str {
			return slice
		}
	}
	return append(slice, str)
}

func getDependendyGraphStructure(traces []Trace) DependencyGraph {
	var nodeNames = make([]string, 0)
	var endpointsMap = make(map[string][]string)

	var nodes = make([]Node, 0)
	var edges = make([]Edge, 0)

	for _, trace := range traces {

		nodeNames = AppendIfMissing(nodeNames, trace.originServiceName)
		nodeNames = AppendIfMissing(nodeNames, trace.toHost)
		endPoints, ok := endpointsMap[trace.toHost]

		if ok {
			endpointsMap[trace.toHost] = AppendIfMissing(endPoints, trace.URI)
		} else {
			endPoints := make([]string, 0)
			endpointsMap[trace.toHost] = append(endPoints, trace.URI)
		}

		edges = append(edges, Edge{
			source:      trace.originServiceName,
			destination: trace.URI,
			attribute:   trace.httMethod,
		})
	}

	for index, nodeName := range nodeNames {
		var endpoints []string
		if nodeName == "client" {
			endpoints = []string{"client"}
		} else {
			endpoints = endpointsMap[nodeName]
		}
		nodes = append(nodes, Node{
			nodeName:    nodeName,
			clusterName: fmt.Sprintf("cluster%d", index),
			endPoints:   endpoints,
		})
	}

	return DependencyGraph{
		nodes: nodes,
		edges: edges,
	}
}

func main() {

	traces := []Trace{
		Trace{
			originServiceName: "client",
			toHost:            "cip-ms-sso-service",
			httMethod:         "GET",
			URI:               "/login",
		},
		Trace{
			originServiceName: "cip-ms-sso-service",
			toHost:            "ping-federate",
			httMethod:         "POST",
			URI:               "/ext/dropoff",
		},
	}

	graphText := createGraph(getDependendyGraphStructure(traces))
	fmt.Println(graphText)
	renderGraph(graphText)
}
