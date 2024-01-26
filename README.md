![License](https://img.shields.io/badge/license-sushiware-red)
![Issues open](https://img.shields.io/github/issues/crashbrz/CrackAnalytics)
![GitHub pull requests](https://img.shields.io/github/issues-pr-raw/crashbrz/CrackAnalytics)
![GitHub closed issues](https://img.shields.io/github/issues-closed-raw/crashbrz/CrackAnalytics)
![GitHub last commit](https://img.shields.io/github/last-commit/crashbrz/CrackAnalytics)

# CrackAnalytics

CrackAnalytics is a GoLang program designed to analyze loaded cracked password lists in-depth. It primarily focuses on analyzing cracked passwords to identify patterns and facilitate the creation of new hashcat rules. This versatile tool provides a comprehensive set of features, including word counts, percentages, character sequence analysis, frequent character identification, and application of custom rules.

## Features
### Total Words and Word Percentages:
- Displays the total number of words in the loaded text file.<br>
- Provides word percentages based on the total number of strings loaded.

### Word Occurrences and Substring Matching:
- Finds occurrences of words or their substrings in a provided keyword list.<br>
- Reports the count of occurrences and the corresponding percentage.

### Word Count by Number of Characters:
- Categorizes words based on their length and counts the occurrences for each category.

### Top Character Sequences:
- Identifies the top 10 most used character sequences in the loaded text.<br>
- Displays their counts and percentages.

### Top Alpha, Number, and Symbol Characters:
- Enumerates the top 10 most frequent alpha, number, and symbol characters.<br>
- Provides counts and percentages for each category.

### Additional Rules:
- Implements custom rules such as checking for four digits at the end, two digits at the end, upper + lower + num + symbol, three digits at the end, first capital + last symbol, first capital + last number, digits between 2000 and 2999, and digits between 1900 and 1999.


Usage example:
```
go run ca.go -f <file_path> -keywords <keyword1,keyword2,...>

 ```
-f <file_path>: Specifies the path to the text file for analysis.<br>
-keywords <keyword1,keyword2,...>: Specifies a comma-separated list of keywords for substring matching.

Example
 ```
go run ca.go -f ./sample.txt -keywords metallica,matrix
Total words in the file: 10006
Word percentages based on the total number of strings loaded:
  metallica: 0.02%
  matrix: 0.01%
Occurrences of words or their substrings in the keyword list:
  metallica: 2
  matrix: 1
Word count by number of characters:
  Words with 14 characters: 6
  Words with 9 characters: 890
  Words with 11 characters: 122
  Words with 3 characters: 6
  Words with 2 characters: 3
  Words with 15 characters: 4
  Words with 7 characters: 2426
  Words with 10 characters: 412
  Words with 16 characters: 2
  Words with 12 characters: 37
  Words with 8 characters: 1782
  Words with 4 characters: 29
  Words with 13 characters: 17
  Words with 17 characters: 2
  Words with 18 characters: 1
  Words with 6 characters: 3492
  Words with 5 characters: 774
  Words with 1 characters: 1


Top 10 most frequent alpha characters:
1. a (Count: 6913, Percentage: 69.09%)
2. e (Count: 6685, Percentage: 66.81%)
3. i (Count: 4660, Percentage: 46.57%)
4. o (Count: 4518, Percentage: 45.15%)
5. n (Count: 4194, Percentage: 41.91%)
6. r (Count: 4185, Percentage: 41.82%)
7. l (Count: 3936, Percentage: 39.34%)
8. s (Count: 3782, Percentage: 37.80%)
9. t (Count: 2945, Percentage: 29.43%)
10. m (Count: 2341, Percentage: 23.40%)

Top 10 most frequent number characters:
1. 1 (Count: 2269, Percentage: 22.68%)
2. 2 (Count: 1021, Percentage: 10.20%)
3. 3 (Count: 694, Percentage: 6.94%)
4. 5 (Count: 428, Percentage: 4.28%)
5. 4 (Count: 419, Percentage: 4.19%)
6. 0 (Count: 406, Percentage: 4.06%)
7. 9 (Count: 367, Percentage: 3.67%)
8. 6 (Count: 366, Percentage: 3.66%)
9. 8 (Count: 333, Percentage: 3.33%)
10. 7 (Count: 298, Percentage: 2.98%)

Top 10 most frequent symbol characters:
1. ! (Count: 21, Percentage: 0.21%)
2. . (Count: 13, Percentage: 0.13%)
3. * (Count: 6, Percentage: 0.06%)
4. # (Count: 3, Percentage: 0.03%)
5. - (Count: 3, Percentage: 0.03%)
6. @ (Count: 2, Percentage: 0.02%)
7. ; (Count: 2, Percentage: 0.02%)
8. $ (Count: 2, Percentage: 0.02%)
9. ? (Count: 1, Percentage: 0.01%)
10. % (Count: 1, Percentage: 0.01%)

Additional Rules:
Four digits at the end: Count: 542, Percentage: 5.42%
Two digits at the end: Count: 1253, Percentage: 12.52%
Upper + lower + num + symbol: Count: 1, Percentage: 0.01%
Three digits at the end: Count: 752, Percentage: 7.52%
First capital, last symbol: Count: 1, Percentage: 0.01%
First capital, last number: Count: 15, Percentage: 0.15%
Digits are between 2000 and 2999: Count: 5, Percentage: 0.05%
Digits are between 1900 and 1999: Count: 0, Percentage: 0.00%

 ```
 
### Installation ###
Clone the repository in the desired location.<br>

### License ###
CrackAnalytics is licensed under the SushiWare license. Check [docs/license.txt](docs/license.txt) for more information.

### Usage/Help ###
You can contact me (@crashbrz) on Twitter or  https://de.linkedin.com/in/crashbrz on LinkedIn<br>

