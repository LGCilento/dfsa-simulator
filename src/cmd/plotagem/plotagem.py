'''import pandas as pd
import plotly.express as px

df = pd.read_csv('result_schoute.csv')

fig = px.line(df, x = 'TagNumber', y = 'EmptySlots', title='Slots Vazios')
fig.show()'''

import matplotlib.pyplot as plt
import csv

TagNumber=[]
SlotsSum_schoute = []
EmptySlots_schoute = []
SuccessfulSlots_schoute = []
CollisionSlots_schoute = []
SlotsSum_eomlee = []
EmptySlots_eomlee = []
SuccessfulSlots_eomlee = []
CollisionSlots_eomlee = []
SlotsSum_lowerbound = []
EmptySlots_lowerbound = []
SuccessfulSlots_lowerbound = []
CollisionSlots_lowerbound = []
with open('result_schoute.csv', 'r') as csvfile:
    plots= csv.reader(csvfile, delimiter=',')
    for row in plots:
        TagNumber.append(int(row[0]))
        SlotsSum_schoute.append(int(row[1]))
        EmptySlots_schoute.append(int(row[2]))
        SuccessfulSlots_schoute.append(int(row[3]))
        CollisionSlots_schoute.append(int(row[4]))
       # x.append(int(row[0]))
        #y.append(int(row[1]))
with open('result_eomlee.csv', 'r') as csvfile:
    plots= csv.reader(csvfile, delimiter=',')
    for row in plots:
        #TagNumber.append(int(row[0]))
        SlotsSum_eomlee.append(int(row[1]))
        EmptySlots_eomlee.append(int(row[2]))
        SuccessfulSlots_eomlee.append(int(row[3]))
        CollisionSlots_eomlee.append(int(row[4]))
with open('result_lowerbound.csv', 'r') as csvfile:
    plots= csv.reader(csvfile, delimiter=',')
    for row in plots:
        #TagNumber.append(int(row[0]))
        SlotsSum_lowerbound.append(int(row[1]))
        EmptySlots_lowerbound.append(int(row[2]))
        SuccessfulSlots_lowerbound.append(int(row[3]))
        CollisionSlots_lowerbound.append(int(row[4]))



fig, ax = plt.subplots(2, 2, figsize=(6, 4))
#print(fig, ax)
ax[0, 0].plot(TagNumber,SlotsSum_schoute, marker='o',label='schoute')
ax[0, 0].plot(TagNumber,SlotsSum_eomlee, marker='^',label='eomlee')
ax[0, 0].plot(TagNumber,SlotsSum_lowerbound, marker='x',label='lowerbound')
ax[0, 0].set(title='Número de Etiquetas X Número de slots')
ax[0, 0].set(xlabel='Número de Etiquetas')
ax[0, 0].set(ylabel='Número de slots')
ax[0, 0].legend(loc='upper left', shadow=True, fontsize='large')
ax[0, 0].grid(b=True, which='major', color='#666666', linestyle='--')

ax[0, 1].plot(TagNumber,EmptySlots_schoute, marker='o',label='schoute')
ax[0, 1].plot(TagNumber,EmptySlots_eomlee, marker='^',label='eomlee')
ax[0, 1].plot(TagNumber,EmptySlots_lowerbound, marker='x',label='lowerbound')
ax[0, 1].set(title='Número de Etiquetas X Número de slots Vazios')
ax[0, 1].set(xlabel = 'Número de Etiquetas')
ax[0, 1].set(ylabel = 'Número de slots Vazios')
ax[0, 1].legend(loc='upper left', shadow=True, fontsize='large')
ax[0, 1].grid(b=True, which='major', color='#666666', linestyle='--')

ax[1, 0].plot(TagNumber,CollisionSlots_schoute, marker='o',label='schoute')
ax[1, 0].plot(TagNumber,CollisionSlots_eomlee, marker='^',label='eomlee')
ax[1, 0].plot(TagNumber,CollisionSlots_lowerbound, marker='x',label='lowerbound')
ax[1, 0].set(title='Número de Etiquetas X Número de slots em Colisão')
ax[1, 0].set(xlabel = 'Número de Etiquetas')
ax[1, 0].set(ylabel = 'Número de slots Vazios')
ax[1, 0].legend(loc='upper left', shadow=True, fontsize='large')
ax[1, 0].grid(b=True, which='major', color='#666666', linestyle='--')

#plt.title('Número de Etiquetas X Número de slots')

#plt.xlabel('Número de Etiquetas')
#plt.ylabel('Número de slots')
#legend = plt.legend(loc='upper left', shadow=True, fontsize='large')
plt.show()