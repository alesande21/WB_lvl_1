## L 1.20

### Задача
Разработать программу, которая переворачивает слова в строке.

Пример: «snow dog sun — sun dog snow».

### Комментарий
Разделяем строку на слова, используя пробел в качестве разделителя функция split возвращает срез строк.
Перебираем срез строк, чтобы получить перевернутую последовательность слов. Объединяем перевернутые слова обратно в строку с помоощью метода join