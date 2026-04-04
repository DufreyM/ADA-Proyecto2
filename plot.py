import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv("results/data.csv")

plt.plot(df["n"], df["recursive"], label="Divide and Conquer")
plt.plot(df["n"], df["dp"], label="DP")

plt.xlabel("Número de intervalos")
plt.ylabel("Tiempo (segundos)")
plt.title("Comparación de algoritmos")
plt.legend()
plt.show()