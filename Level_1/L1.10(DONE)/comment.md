## L 1.10

### Задача
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

### Комментарий
Используется карта (map) для хранения температурных значений, сгруппированных по диапазонам с шагом 10 градусов.
Ключом карты является целое число, представляющее диапазон, а значением — слайс ([]float32) температур, которые попадают в этот диапазон.