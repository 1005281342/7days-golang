package gee

import (
	"fmt"
	"log"
	"strings"
)

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		// 已经处理完pattern
		n.pattern = pattern
		return
	}

	part := parts[height]
	// 获取子节点
	child := n.matchChild(part)
	if child == nil {
		// 子节点不存在，进行创建
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		log.Printf("n.pattern: %s", n.pattern)
		if n.pattern == "" {
			// 未匹配到节点
			return nil
		}
		return n
	}

	part := parts[height]
	// 获取子节点列表
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			log.Printf("n.pattern: %s, result: %+v", n.pattern, result)
			// 匹配到可用节点
			return result
		}
	}
	// 未匹配到节点
	return nil
}

func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		// TODO
		if child.part == part {
			log.Printf("child.part: %s, part: %s, child.isWild: %v", child.part, part, child.isWild)
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		log.Printf("child.part: %s, part: %s", child.part, part)
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	log.Printf("matchChildren nodes %v", nodes)
	return nodes
}
