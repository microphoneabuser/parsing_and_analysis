#%%
import json
import os
import pandas as pd
import matplotlib.pyplot as plt

dir = os.path.dirname(__file__)
filename = os.path.join(dir, 'comics_parsing_goroutine/comics.json')

file = open(filename, 'r', encoding='utf-8')

data = json.loads(file.read())

my_dict = {}

# print(type(data['comics']))

for elem in data['comics']:
    if my_dict.get(elem['year']):
        my_dict[elem['year']] += 1
    else:
        my_dict[elem['year']] = 1

print(my_dict)

ser = pd.Series(my_dict)
fig = plt.figure(num=None, figsize=(14, 8))
plt.plot(ser)
plt.xlabel('Год выпуска')
plt.ylabel('Количество комиксов')
plt.show()
fig.savefig('saved_figure.png')
# %%