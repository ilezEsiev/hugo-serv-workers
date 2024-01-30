---
menu:
    before:
        name: graph
        weight: 1
title: Построение графа
---
# Построение графа

{{< mermaid >}}
graph LR
1(Node 1) --> 2((Node 2))
2((Node 2)) --> 5(Node 5)
2((Node 2)) --> 4{Node 4}
4{Node 4} --> 3(Node 3)
{{< /mermaid >}}