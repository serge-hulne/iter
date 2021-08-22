'''
a = [1, 2, 3]
squared = [x**2 for x in a]
print(squared)

print("\n- - -")

for item in a :
    b = 2 * a

words = []
with open("text.txt") as text:
    for line in text:
        for word in line.split(" "):
            word = word.strip() 
            if len(word) > 1 :
                print(word)
                words.append(word)
words.sort()
print(words)

print("\n- - -")
'''

lines = open("text.txt")
words_array = [line.split() for line in lines]
words = [word.split() for sublist in words_array for word in sublist]
words.sort()
print(words)
