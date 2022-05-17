# Aide - твой помощник при работе с файлами
Запустите приложение через командную строку, с помощью флагов:

 + **-dir**: сгруппировать файлы по папкам внутри заданной директории
 + **-clear**: Очистить целевой каталог
 + **-file**: Создает нужное кол-во файлов в указанной директории
 + **-list**: Создает документ со списком всех файлов в директории

### Пример использования
###### Сгруппировать файлы в нужной директории:
        aide -dir .
сгруппирует файлы в корневой папке.

###### Очистить каталог:
        aide -clear ./dirname/
   очистит каталог, оставив только основную папку.     
###### Создать n кол-во файлов :
        aide -file ./dirname/
   после ввода данной команды, приложение запросит какое 
   количество файлов создать, вводить только цифры, после чего создастся
   3 типа файлов, формата ".txt", ".doc", ".zip"      
 ###### Список файлов:
         aide -list ./dirname/
 после воода данной команды, в выбранной директории создастся файл
 access.log со списком всех файлов и папок в данной директории.                
