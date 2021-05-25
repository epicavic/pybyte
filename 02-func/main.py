def print_max(x, y):
    '''Prints the maximum of two numbers.

    The two values must be integers.'''
    x = int(x)
    y = int(y)
    if x > y:
        print(x, 'is maximum')
    else:
        print(y, 'is maximum')

print("--- calling print_max with 3 and 5 as arguments")
print_max(3, 5)
print("--- print print_max docstring")
print(print_max.__doc__)
