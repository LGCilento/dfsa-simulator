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
SimulationTime_schoute = []
SimulatorTime_schoute = []
ErrorNumber_schoute = []

SlotsSum_eomlee = []
EmptySlots_eomlee = []
SuccessfulSlots_eomlee = []
CollisionSlots_eomlee = []
SimulationTime_eomlee = []
SimulatorTime_eomlee = []
ErrorNumber_eomlee = []

SlotsSum_lowerbound = []
EmptySlots_lowerbound = []
SuccessfulSlots_lowerbound = []
CollisionSlots_lowerbound = []
SimulationTime_lowerbound = []
SimulatorTime_lowerbound = []
ErrorNumber_lowerbound = []

SlotsSum_vogt = []
EmptySlots_vogt = []
SuccessfulSlots_vogt = []
CollisionSlots_vogt = []
SimulationTime_vogt = []
SimulatorTime_vogt = []
ErrorNumber_vogt = []

with open('result_schoute.csv', 'r') as csvfile:
    plots= csv.reader(csvfile, delimiter=',')
    for row in plots:
        TagNumber.append(int(row[0]))
        SlotsSum_schoute.append(int(row[1]))
        EmptySlots_schoute.append(int(row[2]))
        SuccessfulSlots_schoute.append(int(row[3]))
        CollisionSlots_schoute.append(int(row[4]))
        SimulationTime_schoute.append(float(row[5]))
        SimulatorTime_schoute.append(float(row[6]))
        ErrorNumber_schoute.append(int(row[7]))

with open('result_eomlee.csv', 'r') as csvfile:
    plots= csv.reader(csvfile, delimiter=',')
    for row in plots:
        #TagNumber.append(int(row[0]))
        SlotsSum_eomlee.append(int(row[1]))
        EmptySlots_eomlee.append(int(row[2]))
        SuccessfulSlots_eomlee.append(int(row[3]))
        CollisionSlots_eomlee.append(int(row[4]))
        SimulationTime_eomlee.append(float(row[5]))
        SimulatorTime_eomlee.append(float(row[6]))
        ErrorNumber_eomlee.append(int(row[7]))

with open('result_lowerbound.csv', 'r') as csvfile:
    plots= csv.reader(csvfile, delimiter=',')
    for row in plots:
        #TagNumber.append(int(row[0]))
        SlotsSum_lowerbound.append(int(row[1]))
        EmptySlots_lowerbound.append(int(row[2]))
        SuccessfulSlots_lowerbound.append(int(row[3]))
        CollisionSlots_lowerbound.append(int(row[4]))
        SimulationTime_lowerbound.append(float(row[5]))
        SimulatorTime_lowerbound.append(float(row[6]))
        ErrorNumber_lowerbound.append(int(row[7]))

with open('result_vogt.csv', 'r') as csvfile:
    plots= csv.reader(csvfile, delimiter=',')
    for row in plots:
        #TagNumber.append(int(row[0]))
        SlotsSum_vogt.append(int(row[1]))
        EmptySlots_vogt.append(int(row[2]))
        SuccessfulSlots_vogt.append(int(row[3]))
        CollisionSlots_vogt.append(int(row[4]))
        SimulationTime_vogt.append(float(row[5]))
        SimulatorTime_vogt.append(float(row[6]))
        ErrorNumber_vogt.append(int(row[7]))



fig, ax = plt.subplots(2, 3, figsize=(4, 10))
#print(fig, ax)
ax[0, 0].plot(TagNumber,SlotsSum_schoute, marker='o',label='schoute')
ax[0, 0].plot(TagNumber,SlotsSum_eomlee, marker='^',label='eomlee')
ax[0, 0].plot(TagNumber,SlotsSum_lowerbound, marker='x',label='lowerbound')
ax[0, 0].plot(TagNumber,SlotsSum_vogt, marker='+',label='vogt')
ax[0, 0].set(title='Número de Etiquetas X Número de slots')
ax[0, 0].set(xlabel='Número de Etiquetas')
ax[0, 0].set(ylabel='Número de slots')
ax[0, 0].legend(loc='upper left', shadow=True, fontsize='large')
ax[0, 0].grid(b=True, which='major', color='#666666', linestyle='--')

ax[0, 1].plot(TagNumber,EmptySlots_schoute, marker='o',label='schoute')
ax[0, 1].plot(TagNumber,EmptySlots_eomlee, marker='^',label='eomlee')
ax[0, 1].plot(TagNumber,EmptySlots_lowerbound, marker='x',label='lowerbound')
ax[0, 1].plot(TagNumber,EmptySlots_vogt, marker='+',label='vogt')
ax[0, 1].set(title='Número de Etiquetas X Número de slots Vazios')
ax[0, 1].set(xlabel = 'Número de Etiquetas')
ax[0, 1].set(ylabel = 'Número de slots Vazios')
ax[0, 1].legend(loc='upper left', shadow=True, fontsize='large')
ax[0, 1].grid(b=True, which='major', color='#666666', linestyle='--')

ax[0, 2].plot(TagNumber,CollisionSlots_schoute, marker='o',label='schoute')
ax[0, 2].plot(TagNumber,CollisionSlots_eomlee, marker='^',label='eomlee')
ax[0, 2].plot(TagNumber,CollisionSlots_lowerbound, marker='x',label='lowerbound')
ax[0, 2].plot(TagNumber,CollisionSlots_vogt, marker='+',label='vogt')
ax[0, 2].set(title='Número de Etiquetas X Número de slots em Colisão')
ax[0, 2].set(xlabel = 'Número de Etiquetas')
ax[0, 2].set(ylabel = 'Número de slots Vazios')
ax[0, 2].legend(loc='upper left', shadow=True, fontsize='large')
ax[0, 2].grid(b=True, which='major', color='#666666', linestyle='--')

ax[1, 0].plot(TagNumber,SimulationTime_schoute, marker='o',label='schoute')
ax[1, 0].plot(TagNumber,SimulationTime_eomlee, marker='^',label='eomlee')
ax[1, 0].plot(TagNumber,SimulationTime_lowerbound, marker='x',label='lowerbound')
ax[1, 0].plot(TagNumber,SimulationTime_vogt, marker='+',label='vogt')
ax[1, 0].set(title='Número de Etiquetas X Tempo de Simulação')
ax[1, 0].set(xlabel='Número de Etiquetas')
ax[1, 0].set(ylabel='Número de slots')
ax[1, 0].legend(loc='upper left', shadow=True, fontsize='large')
ax[1, 0].grid(b=True, which='major', color='#666666', linestyle='--')

ax[1, 1].plot(TagNumber,SimulatorTime_schoute, marker='o',label='schoute')
ax[1, 1].plot(TagNumber,SimulatorTime_eomlee, marker='^',label='eomlee')
ax[1, 1].plot(TagNumber,SimulatorTime_lowerbound, marker='x',label='lowerbound')
ax[1, 1].plot(TagNumber,SimulatorTime_vogt, marker='+',label='vogt')
ax[1, 1].set(title='Número de Etiquetas X Tempo do Simulador')
ax[1, 1].set(xlabel = 'Número de Etiquetas')
ax[1, 1].set(ylabel = 'Número de slots Vazios')
ax[1, 1].legend(loc='upper left', shadow=True, fontsize='large')
ax[1, 1].grid(b=True, which='major', color='#666666', linestyle='--')

ax[1, 2].plot(TagNumber,ErrorNumber_schoute, marker='o',label='schoute')
ax[1, 2].plot(TagNumber,ErrorNumber_eomlee, marker='^',label='eomlee')
ax[1, 2].plot(TagNumber,ErrorNumber_lowerbound, marker='x',label='lowerbound')
ax[1, 2].plot(TagNumber,ErrorNumber_vogt, marker='+',label='vogt')
ax[1, 2].set(title='Número de Etiquetas X Erro Médio')
ax[1, 2].set(xlabel = 'Número de Etiquetas')
ax[1, 2].set(ylabel = 'Número de slots Vazios')
ax[1, 2].legend(loc='upper left', shadow=True, fontsize='large')
ax[1, 2].grid(b=True, which='minor', color='#666666', linestyle='--')

#ax[1, 2].xaxis.grid(True, which='minor')
#plt.title('Número de Etiquetas X Número de slots')

#plt.xlabel('Número de Etiquetas')
#plt.ylabel('Número de slots')
#legend = plt.legend(loc='upper left', shadow=True, fontsize='large')
plt.show()