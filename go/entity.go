package main;

import (
    "strings"
    "encoding/json"
    "github.com/qedus/osmpbf"
);

type Entity struct {
    ID int64 `json:"id"`
    Type string `json:"type"`
    Name string `json:"name"`
}

type Node struct {
    *Entity
    Lat float64 `json:"lat"`
    Lon float64 `json:"lon"`
}

type Way struct {
    *Entity
    Centroid map[string]string `json:"centroid"`
    Bounds map[string]string `json:"bounds"`
    Nodes []map[string]string `json:"nodes,omitempty"`
}

func createNode(node *osmpbf.Node) string {
    entity := Entity{node.ID, "node", node.Tags["name"]};
    marshal := Node{&entity, node.Lat, node.Lon};

    json, _ := json.Marshal(marshal);

    return string(json);
}

func isValidEntity(Tags map[string]string) bool {
    name, ok := Tags["name"];

    return ok && strings.Contains(name, searchTerm);
}