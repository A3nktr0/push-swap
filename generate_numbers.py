import random

numbers = random.sample(range(999), 100)

with open('100numbers.txt', 'w') as f:
    for n in numbers:
        f.write('%d\n' % n)
f.close()