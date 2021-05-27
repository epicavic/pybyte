class Person:
    def __init__(self, fname, lname):
        self.fname = fname
        self.lname = lname
    def say_hi(self):
        print(f'How are you {self.fname} {self.lname} ?')
p1 = Person("James", "Bond")
p2 = Person("Miss", "Moneypenny")
p1.say_hi()
p2.say_hi()
