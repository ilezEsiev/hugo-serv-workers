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
1([Node 1]) --> 15([Node 15])
1([Node 1]) --> 7{Node 7}
1([Node 1]) --> 4{Node 4}
1([Node 1]) --> 9([Node 9])
1([Node 1]) --> 11(Node 11)
1([Node 1]) --> 2{Node 2}
2{Node 2} --> 9([Node 9])
4{Node 4} --> 11(Node 11)
4{Node 4} --> 14[Node 14]
4{Node 4} --> 6(Node 6)
4{Node 4} --> 7{Node 7}
4{Node 4} --> 10([Node 10])
5{Node 5} --> 15([Node 15])
5{Node 5} --> 11(Node 11)
5{Node 5} --> 7{Node 7}
5{Node 5} --> 13[Node 13]
6(Node 6) --> 8([Node 8])
6(Node 6) --> 11(Node 11)
6(Node 6) --> 10([Node 10])
6(Node 6) --> 15([Node 15])
6(Node 6) --> 3[Node 3]
7{Node 7} --> 14[Node 14]
7{Node 7} --> 15([Node 15])
7{Node 7} --> 3[Node 3]
7{Node 7} --> 13[Node 13]
7{Node 7} --> 2{Node 2}
7{Node 7} --> 11(Node 11)
7{Node 7} --> 8([Node 8])
8([Node 8]) --> 1([Node 1])
9([Node 9]) --> 8([Node 8])
9([Node 9]) --> 7{Node 7}
9([Node 9]) --> 10([Node 10])
9([Node 9]) --> 3[Node 3]
10([Node 10]) --> 7{Node 7}
10([Node 10]) --> 8([Node 8])
10([Node 10]) --> 3[Node 3]
10([Node 10]) --> 2{Node 2}
10([Node 10]) --> 1([Node 1])
11(Node 11) --> 3[Node 3]
11(Node 11) --> 2{Node 2}
12{Node 12} --> 7{Node 7}
12{Node 12} --> 6(Node 6)
12{Node 12} --> 14[Node 14]
12{Node 12} --> 9([Node 9])
13[Node 13] --> 14[Node 14]
13[Node 13] --> 2{Node 2}
13[Node 13] --> 1([Node 1])
13[Node 13] --> 6(Node 6)
13[Node 13] --> 10([Node 10])
13[Node 13] --> 9([Node 9])
13[Node 13] --> 8([Node 8])
14[Node 14] --> 8([Node 8])
14[Node 14] --> 2{Node 2}
14[Node 14] --> 15([Node 15])
14[Node 14] --> 3[Node 3]
14[Node 14] --> 6(Node 6)
14[Node 14] --> 5{Node 5}
{{< /mermaid >}}