class SchoolMember:
    '''Base class (superclass)'''
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def tell(self):
        print(f'Name:"{self.name}" Age:"{self.age}"', end=" ")

class Teacher(SchoolMember):
    '''Subclass (derived class)'''
    def __init__(self, name, age, salary):
        SchoolMember.__init__(self, name, age)
        self.salary = salary

    def tell(self):
        SchoolMember.tell(self)
        print(f'Salary: "{self.salary}"')

class Student(SchoolMember):
    '''Subclass (derived class)'''
    def __init__(self, name, age, marks):          
        SchoolMember.__init__(self, name, age)
        self.marks = marks

    def tell(self):
        SchoolMember.tell(self)
        print(f'Marks: "{self.marks}"')

m = SchoolMember('Miklos', 20)
t = Teacher('Bond', 40, 30000)
s = Student('Moneypenny', 25, 75)

members = [m, t, s]
for member in members:
   member.tell()
   if member.name == 'Miklos':
       print()