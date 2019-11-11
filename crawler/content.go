package crawler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var ignoretags = map[string]bool{
	"link":   true,
	"style":  true,
	"script": true,
}

type BySum []NodeInfo

func (a BySum) Len() int           { return len(a) }
func (a BySum) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySum) Less(i, j int) bool { return a[i].ChildSum > a[j].ChildSum }

type NodeInfo struct {
	Nodes            []NodeInfo `json:"nodes"`
	NodeName         string     `json:"name"`
	NodeID           string     `json:"id"`
	NodeClass        string     `json:"className"`
	InheritClassName string     `json:"inherit"`
	Depth            int        `json:"depth"`
	ChildLength      int        `json:"child_length"`
	ChildSum         int        `json:"child_sum"`
	Text             string     `json:"text"`
}

type Candidates struct {
	nodes []NodeInfo
}

func (article *Article) ParseContent() {
	candidates := Candidates{}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(article.RawHTML))
	bodyElm := doc.Find("body")
	nodInfo := NodeInfo{}
	nodInfo.travelNode(bodyElm, 0)
	jOut, _ := json.MarshalIndent(nodInfo, "", "  ")
	ioutil.WriteFile("output.json", jOut, 0644)

	if len(nodInfo.Nodes) > 0 {
		candidates.getCandidates(nodInfo.Nodes[0], "")
		sort.Sort(BySum(candidates.nodes))

		if len(candidates.nodes) > 0 {
			log.Println(
				len(candidates.nodes[0].Nodes),
				candidates.nodes[0].InheritClassName,
				candidates.nodes[0].ChildLength,
				candidates.nodes[0].ChildSum,
			)
			contentElm := bodyElm.Find(candidates.nodes[0].InheritClassName).First()

			for key := range ignoretags {
				contentElm.Find(key).Remove()
			}

			if len(candidates.nodes[0].Nodes) > 3 {
				contentElm.Find("div").Remove()
			} else {
				contentElm.Children().Find("div").Remove()
			}

			content := contentElm.Text()
			rawHTML, _ := contentElm.Html()
			content = strings.TrimSpace(content)
			var re = regexp.MustCompile(`\s{2,}`)
			content = re.ReplaceAllString(content, " ")
			article.CleanContent = content

			var removeScripts = regexp.MustCompile(`(?m)(<script(.*)>([^<]+)<\/script>|<script(.*)><\/script>)`)
			rawHTML = removeScripts.ReplaceAllString(rawHTML, "")

			var removeComments = regexp.MustCompile(`(?m)<\!--([^-]+)--\>`)
			rawHTML = removeComments.ReplaceAllString(rawHTML, "")

			article.RawContent = rawHTML
		}
	}
}

func (candidates *Candidates) getCandidates(nodeInfo NodeInfo, rootClassName string) {
	inheritClass := rootClassName

	if nodeInfo.NodeID != "" {
		inheritClass = rootClassName + " #" + nodeInfo.NodeID
	} else if nodeInfo.NodeClass != "" {
		separatedClassName := strings.Split(nodeInfo.NodeClass, " ")
		inheritClass = rootClassName + " ." + separatedClassName[0]
	}

	candidates.nodes = append(candidates.nodes, NodeInfo{
		NodeName:         nodeInfo.NodeName,
		NodeID:           nodeInfo.NodeID,
		NodeClass:        nodeInfo.NodeClass,
		ChildSum:         nodeInfo.ChildSum,
		InheritClassName: inheritClass,
		Nodes:            nodeInfo.Nodes,
	})
	for _, n := range nodeInfo.Nodes {
		candidates.getCandidates(n, inheritClass)
	}
}

func debugNode(nodeInfo NodeInfo) {

	spaces := ""
	for i := 0; i < nodeInfo.Depth; i++ {
		spaces += "  "
	}

	log.Println(spaces, ">", nodeInfo.NodeName)
	if nodeInfo.NodeClass != "" {
		log.Println(spaces, "  |-", "class", nodeInfo.NodeClass)
	}
	if nodeInfo.NodeID != "" {
		log.Println(spaces, "  |-", "ID", nodeInfo.NodeID)
	}

	log.Println(spaces, "  |-", "Length", nodeInfo.ChildLength)
	log.Println(spaces, "  |-", "Text", nodeInfo.Text)
	log.Println(spaces, "  |-", "Sum", nodeInfo.ChildSum)

	for _, n := range nodeInfo.Nodes {
		debugNode(n)
	}

}

func (node *NodeInfo) travelNode(sel *goquery.Selection, depth int) {
	for ii := range sel.Nodes {
		chNode := NodeInfo{}
		chNode.Depth = depth
		single := sel.Eq(ii)
		nodeName := goquery.NodeName(single)

		if _, ok := ignoretags[nodeName]; !ok {
			className, hasClass := single.Attr("class")
			nodeID, hasID := single.Attr("id")

			chNode.NodeName = nodeName

			if hasClass {
				chNode.NodeClass = className
			}

			if hasID {
				chNode.NodeID = nodeID
			}

			ccNode := single.Clone()

			text := ccNode.Children().Remove().End().Text()
			text = strings.TrimSpace(text)

			chNode.ChildLength = len(text)
			chNode.Text = text

			// log.Println(depth, ii, "->", chNode.nodeName, className, text)

			if single.Children().Length() > 0 {
				chNode.travelNode(single.Children(), depth+1)
				for _, n := range chNode.Nodes {
					chNode.ChildSum += n.ChildLength
				}
			}
			node.Nodes = append(node.Nodes, chNode)
		}
	}
}
