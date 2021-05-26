import sys
print('The command line arguments are:', end='')
for i in sys.argv:
    print('', i, end='')
print()
print('The PYTHONPATH is:', sys.path)
