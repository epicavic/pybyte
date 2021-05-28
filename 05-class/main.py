class Person:
    population = 0

    def __init__(self, fname, lname):
        """ All vars assigned using 'self.' are object variables.
        The ones assigned using 'Person.' or 'self.__class__.' are class variables.
        Object variable with the same name as a class variable will hide the class variable.
        Any private variable that is to be used only within the class or object should begin with 
        an underscore and all other names are public and can be used by other classes/objects
        """
        self.fname = fname
        self.lname = lname

        # same as self.__class__.population
        Person.population += 1

    def say_hi(self):
        print(f'How are you {self.fname} {self.lname} ?')

    def die(self):
        Person.population -= 1

        if Person.population == 0:
            print(f'{self.fname} {self.lname} was the last one.')
        else:
            print(f'{self.fname} {self.lname} has died.')

    # decorator. same as how_many = classmethod(how_many)
    @classmethod
    def how_many(cls):
        print(f'We have {Person.population} people in line.')

p1 = Person("James", "Bond")
p2 = Person("Rosika", "Miklos")
p3 = Person("Miss", "Moneypenny")
p1.say_hi()
p2.say_hi()
p3.say_hi()
Person.how_many()
p1.die()
p2.die()
p3.die()
