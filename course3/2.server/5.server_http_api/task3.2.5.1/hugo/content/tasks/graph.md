
---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

Нужно написать воркер, который будет строить граф на текущей странице, каждые 5 секунд
От 5 до 30 элементов, случайным образом. Все ноды графа должны быть связаны.
```go
type Node struct {
    ID int
    Name string
	Form string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
    Links []*Node
}
```

## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Граф

{{< columns >}}
```tpl
{{</*/* mermaid [class="text-center"]*/*/>}}
graph LR
A[A] --> B[B]
A --> C[C]
B --> D[D]
B --> C
C --> E[E]
C --> D
D --> E
D --> G[G]
E --> F((F))
E --> G
F --> G
F --> H([H])
G --> H
G --> I((I))
H --> I
H --> J((J))
I --> J
J --> K([K])
J --> M([M])
K --> L[L]
L --> M

{{</*/* /mermaid */*/>}}
```

<--->

{{< mermaid >}}
graph LR
A[A] --> B[B]
A --> C[C]
B --> D[D]
B --> C
C --> E[E]
C --> D
D --> E
D --> G[G]
E --> F((F))
E --> G
F --> G
F --> H([H])
G --> H
G --> I((I))
H --> I
H --> J((J))
I --> J
J --> K([K])
J --> M([M])
K --> L[L]
L --> M

{{< /mermaid >}}

{{< /columns >}}
