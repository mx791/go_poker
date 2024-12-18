import requests
import numpy as np
import matplotlib.pyplot as plt

x, y = [], []

for count in [100, 1000, 5000, 10_000, 20_000, 40_000, 50_000, 60_000, 80_000]:
    url = f"http://localhost:8080/evaluate?eval_count={int(count)}"
    results = [
        float(requests.post(url, json={
            "personal_cards": ["As de coeur", "As de piques"],
            "common_cards": [],
        }).text) for _ in range(100)
    ]
    print(count, np.std(results))
    x.append(count)
    y.append(np.std(results))


plt.figure(figsize=(12, 6))
plt.plot(x, y)
plt.title("Variance des résultats en fonction du nombre de tirages")
plt.savefig("./images/benchmark.png")