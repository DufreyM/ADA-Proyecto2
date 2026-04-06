# Weighted Interval Scheduling

## Descripción

Este proyecto implementa y analiza el problema de **Weighted Interval Scheduling**, comparando dos enfoques:

* Divide and Conquer (recursivo)
* Programación Dinámica (bottom-up)

El objetivo es seleccionar un subconjunto de intervalos no traslapados que maximicen la ganancia total.

---

## Fundamento teórico

El problema cumple:

* **Optimal Substructure**
* **Overlapping Subproblems**

Por lo tanto, es adecuado para programación dinámica.

---

## Por qué no es greedy

Se demuestra mediante contraejemplo que una estrategia greedy no garantiza la solución óptima.

---

## Algoritmos implementados

### Divide and Conquer

Complejidad:

* Tiempo: O(2^n)

### Programación Dinámica

Complejidad:

* Tiempo: O(n²)
* Optimizado: O(n log n)

---

## Estructura del proyecto

```
models/        → estructura Job  
algorithms/    → lógica de algoritmos  
experiments/   → generación y medición  
results/       → datos experimentales  
```

---

## Resultados

Se realizaron pruebas con tamaños de entrada de 1 a 30.

Observaciones:

* El algoritmo recursivo presenta crecimiento exponencial
* El algoritmo DP es significativamente más eficiente
* Los resultados empíricos coinciden con el análisis teórico

---

## Ejecución

```bash
go run main.go
```

---

## 📈 Visualización

Los resultados se exportan a CSV para graficarse posteriormente.

---

## Video
https://youtu.be/ursuoeX0bB8

---
## Documento del proyecto
https://uvggt-my.sharepoint.com/:w:/g/personal/mej23648_uvg_edu_gt/IQAY8rDJEz_uS7Ut86BvqA-4AdQ35e3WUr2o6kqKxT5LbHA?e=gqyxRr

---
## 👨‍💻 Autores

* Leonardo Dufrey Mejía Mejía, 23648
* María José Girón Isidro, 23559
