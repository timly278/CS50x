print("hello, world")
# a1 = int(input("nhap a1: "))
# a2 = int(input("nhap a2: "))
# s = input("Do you agree? ")

# if s.lower() in ["y", "yes"]:
#     print("Agreed.")
# elif s.lower() in ["n", "no"]:
#     print("Not agreed.")

# i = 0
# while i < 3:
#     print("meow")
#     i += 1

# for i in range(3):
#     print("meow", i)


#######################################
#               MARIO
#######################################

# def main():
#     height = get_height()
#     for i in range(height):
#         print("#")


# def get_height():
#     while True:
#         try:
#             n = int(input("Height: "))
#             print("n = ", n)
#             if n > 0:
#                 return n
#         except ValueError:
#             print("Not an integer")



# for i in range(4):
#     # print("?", end="")
#     print("?" * 4) 
# print()


#######################################
#               SCORE.py
#######################################
# def main():
#     scores = []
#     for i in range(3):
#         score = get_int("Score: ")
#         # scores += [score]
#         scores.append(score)

#     average = sum(scores) / len(scores)
#     print(f"Average: {average}")


#######################################
#               Uppercase.py
#######################################
# def main():
#     before = input("Before: ")
#     print("After: ", end="")
#     after = before.upper()
#     print(f"{after}")

#######################################
#               Greet.py
#######################################
# from sys import argv
# def main():
#     if len(argv) == 2:
#         print("hello,", argv[1])
#     else:
#         print("hello, world")

#     # for i in range(len(argv)):
#     #     print(f"argv[{i}] = {argv[i]}")
    
#     for arg in argv[1:]:
#         print(arg)

#######################################
#               Exit.py
#######################################
# import sys #import the whole library

# def main():
#     if len(sys.argv) != 2:
#         print("Missing command-line argument")
#         sys.exit(1)

#     print(f"hello, {sys.argv[1]}")
#     sys.exit(0) #programm returns no error


#######################################
#               Search.py
#######################################
# import sys #import the whole library

# names = ["Bill", "Charlie", "Fred", "Ginny"]

# def main():
#     name = input("Name: ")
#     if name in names:
#         print("Found")
#         sys.exit(0)
#     print("Not Found")
#     sys.exit(1)



#######################################
#               phonebook.py
#######################################

people = {
    "Carter": "+1-212-121-33323",
    "David": "+1-344-987-9999"
}

def main():
    name = input("Name: ")
    if name in people:
        print(f"Number: {people[name]}")
    else:
        print("not found name in the list")

#######################################
#               Swap.py
#######################################
x = 1
y = 2
x, y = y, x



def get_int(s):
    while True:
        try:
            n = int(input(s))
            if n > 0:
                return n
        except ValueError:
            print("Not an integer")

main()